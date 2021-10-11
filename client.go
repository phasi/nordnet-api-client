package nordnet_api_client

import "encoding/json"

// NordnetClient : Use NewNordnetClient() to initialize new client.
type NordnetClient struct {
	Driver ClientDriver
}

// NewNordnetClient : Creates new client.
func NewNordnetClient(driver ClientDriver) NordnetClient {
	client := NordnetClient{
		Driver: driver,
	}
	return client
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
	req, err := nc.Driver.PrepareGetRequest("/login")
	bytes, err := nc.Driver.SendRequest(req)
	if err != nil {
		return
	}
	err = json.Unmarshal(bytes, &r)
	return
}
