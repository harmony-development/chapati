package v1

import "net/http"
import "google.golang.org/protobuf/proto"
import "io/ioutil"
import "fmt"
import "github.com/gorilla/websocket"
import "net/url"
import "bytes"

import "github.com/golang/protobuf/ptypes/empty"

type PostboxServiceClient struct {
	client    *http.Client
	serverURL string

	Header    http.Header
	HTTPProto string
	WSProto   string
}

func NewPostboxServiceClient(url string) *PostboxServiceClient {
	return &PostboxServiceClient{
		client:    &http.Client{},
		serverURL: url,
		Header:    http.Header{},
		HTTPProto: "https",
		WSProto:   "wss",
	}
}

func (client *PostboxServiceClient) Sync(r *SyncRequest) (chan *PostBoxEvent, error) {
	u := url.URL{Scheme: client.WSProto, Host: client.serverURL, Path: "/protocol.sync.v1.PostboxService/Sync"}

	c, _, err := websocket.DefaultDialer.Dial(u.String(), client.Header)
	if err != nil {
		return nil, err
	}

	data, err := proto.Marshal(r)
	if err != nil {
		return nil, err
	}

	err = c.WriteMessage(websocket.BinaryMessage, data)
	if err != nil {
		return nil, err
	}

	outC := make(chan *PostBoxEvent)

	go func() {
		defer c.Close()

		msgs := make(chan []byte)

		go func() {
			for {
				_, message, err := c.ReadMessage()
				if err != nil {
					close(msgs)
					break
				}
				msgs <- message
			}
		}()

		defer close(outC)
		for {
			select {
			case msg, ok := <-msgs:
				if !ok {
					return
				}

				thing := new(PostBoxEvent)
				err = proto.Unmarshal(msg, thing)
				if err != nil {
					return
				}

				outC <- thing
			}
		}
	}()

	return outC, nil
}

func (client *PostboxServiceClient) PostEvent(r *PostEventRequest) (*empty.Empty, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s://%s/protocol.sync.v1.PostboxService/PostEvent", client.HTTPProto, client.serverURL), bytes.NewReader(input))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	for k, v := range client.Header {
		req.Header[k] = v
	}
	req.Header.Add("content-type", "application/hrpc")
	resp, err := client.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error posting request: %w", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	output := &empty.Empty{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}
