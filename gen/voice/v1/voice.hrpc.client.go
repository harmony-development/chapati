package v1

import "net/http"
import "google.golang.org/protobuf/proto"
import "io/ioutil"
import "fmt"
import "github.com/gorilla/websocket"
import "net/url"
import "bytes"

type VoiceServiceClient struct {
	client    *http.Client
	serverURL string

	Header    http.Header
	HTTPProto string
	WSProto   string
}

func NewVoiceServiceClient(url string) *VoiceServiceClient {
	return &VoiceServiceClient{
		client:    &http.Client{},
		serverURL: url,
		Header:    http.Header{},
		HTTPProto: "https",
		WSProto:   "wss",
	}
}

func (client *VoiceServiceClient) Connect(r *ConnectRequest) (*ConnectResponse, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s://%s/protocol.voice.v1.VoiceService/Connect", client.HTTPProto, client.serverURL), bytes.NewReader(input))
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
	output := &ConnectResponse{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}

func (client *VoiceServiceClient) StreamState(r *StreamStateRequest) (chan *Signal, error) {
	u := url.URL{Scheme: client.WSProto, Host: client.serverURL, Path: "/protocol.voice.v1.VoiceService/StreamState"}

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

	outC := make(chan *Signal)

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

				thing := new(Signal)
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
