package nordnet_api_client

import (
	"encoding/json"
)

// NordnetClient : Use NewNordnetClient() to initialize new client.
type NordnetClient struct {
	Driver ClientDriver
}

// NewNordnetClient : Creates new client with default driver
func NewNordnetClient() NordnetClient {
	client := NordnetClient{
		Driver: DefaultDriver,
	}
	return client
}

func (nc NordnetClient) get(path string) (bytes []byte, err error) {
	req, err := nc.Driver.PrepareGetRequest(path)
	if err != nil {
		return
	}
	bytes, err = nc.Driver.SendRequest(req)
	return
}

func (nc NordnetClient) post(path string, data []byte) (bytes []byte, err error) {
	req, err := nc.Driver.PreparePostRequest(path, data)
	if err != nil {
		return
	}
	bytes, err = nc.Driver.SendRequest(req)
	return
}

// LoginResponse : API Response for login refresh requests
type LoginResponse struct {
	Remaining   int    `json:"remaining"`
	LoggedIn    bool   `json:"logged_in"`
	Lang        string `json:"lang"`
	Country     string `json:"country"`
	SessionType string `json:"session_type"`
}

// Login : Refresh user login
func (nc NordnetClient) Login() (r LoginResponse, err error) {
	bytes, err := nc.get("/login")
	err = json.Unmarshal(bytes, &r)
	return
}
