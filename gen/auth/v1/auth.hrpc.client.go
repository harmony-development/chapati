package v1

import "net/http"
import "google.golang.org/protobuf/proto"
import "io/ioutil"
import "fmt"
import "github.com/gorilla/websocket"
import "net/url"
import "bytes"

import "github.com/golang/protobuf/ptypes/empty"

type AuthServiceClient struct {
	client    *http.Client
	serverURL string

	Header    http.Header
	HTTPProto string
	WSProto   string
}

func NewAuthServiceClient(url string) *AuthServiceClient {
	return &AuthServiceClient{
		client:    &http.Client{},
		serverURL: url,
		Header:    http.Header{},
		HTTPProto: "https",
		WSProto:   "wss",
	}
}

func (client *AuthServiceClient) Federate(r *FederateRequest) (*FederateReply, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s://%s/protocol.auth.v1.AuthService/Federate", client.HTTPProto, client.serverURL), bytes.NewReader(input))
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
	output := &FederateReply{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}

func (client *AuthServiceClient) LoginFederated(r *LoginFederatedRequest) (*Session, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s://%s/protocol.auth.v1.AuthService/LoginFederated", client.HTTPProto, client.serverURL), bytes.NewReader(input))
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
	output := &Session{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}

func (client *AuthServiceClient) Key(r *empty.Empty) (*KeyReply, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s://%s/protocol.auth.v1.AuthService/Key", client.HTTPProto, client.serverURL), bytes.NewReader(input))
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
	output := &KeyReply{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}

func (client *AuthServiceClient) BeginAuth(r *empty.Empty) (*BeginAuthResponse, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s://%s/protocol.auth.v1.AuthService/BeginAuth", client.HTTPProto, client.serverURL), bytes.NewReader(input))
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
	output := &BeginAuthResponse{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}

func (client *AuthServiceClient) NextStep(r *NextStepRequest) (*AuthStep, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s://%s/protocol.auth.v1.AuthService/NextStep", client.HTTPProto, client.serverURL), bytes.NewReader(input))
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
	output := &AuthStep{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}

func (client *AuthServiceClient) StepBack(r *StepBackRequest) (*AuthStep, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s://%s/protocol.auth.v1.AuthService/StepBack", client.HTTPProto, client.serverURL), bytes.NewReader(input))
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
	output := &AuthStep{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}

func (client *AuthServiceClient) StreamSteps(r *StreamStepsRequest) (chan *AuthStep, error) {
	u := url.URL{Scheme: client.WSProto, Host: client.serverURL, Path: "/protocol.auth.v1.AuthService/StreamSteps"}

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

	outC := make(chan *AuthStep)

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

				thing := new(AuthStep)
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

func (client *AuthServiceClient) CheckLoggedIn(r *empty.Empty) (*empty.Empty, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s://%s/protocol.auth.v1.AuthService/CheckLoggedIn", client.HTTPProto, client.serverURL), bytes.NewReader(input))
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
