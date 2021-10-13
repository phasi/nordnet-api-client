package nordnet_api_client

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

func appendHeaders(req *http.Request) {
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.159 Safari/537.36")
	req.Header.Add("client-id", "NEXT")
	req.Header.Add("Referer", "https://www.nordnet.fi/")
	req.Header.Add("Origin", "https://www.nordnet.fi")
	req.Header.Add("Accept-Language", "en-US,en;q=0.9,fi-FI;q=0.8,fi;q=0.7,es-CO;q=0.6,es;q=0.5")
	req.Header.Add("Sec-Fetch-Site", "same-origin")
	req.Header.Add("Sec-Fetch-Mode", "cors")
	req.Header.Add("Sec-Fetch-Dest", "empty")
}

// HTTPDriver : Implementation of ClientDriver interface
// to interact with Nordnet API. Use DefaultDriver if in doubt.
type HTTPDriver struct {
	APIBasePath      string
	Client           http.Client
	MandatoryHeaders http.Header
}

// DefaultDriver : Default driver for connecting to Nordnet API
var DefaultDriver HTTPDriver = HTTPDriver{
	APIBasePath: "https://www.nordnet.fi/api/2",
	Client: http.Client{
		Timeout: 5 * time.Second,
	},
	MandatoryHeaders: make(http.Header),
}

// PrepareGetRequest : Prepares GET request and returns it.
func (httpDriver HTTPDriver) PrepareGetRequest(path string) (req *http.Request, err error) {
	req, err = http.NewRequest("GET", fmt.Sprintf("%s%s", httpDriver.APIBasePath, path), nil)
	appendHeaders(req)
	for k, v := range httpDriver.MandatoryHeaders {
		req.Header[k] = v
	}
	return
}

// PreparePostRequest : Prepares POST request and returns it
func (httpDriver HTTPDriver) PreparePostRequest(path string, jsonBytes []byte) (req *http.Request, err error) {
	req, err = http.NewRequest("POST", fmt.Sprintf("%s%s", httpDriver.APIBasePath, path), bytes.NewBuffer(jsonBytes))
	appendHeaders(req)
	for k, v := range httpDriver.MandatoryHeaders {
		req.Header[k] = v
	}
	return
}

// SendRequest : Sends HTTP Request
func (httpDriver HTTPDriver) SendRequest(req *http.Request) (bytes []byte, err error) {
	resp, err := httpDriver.Client.Do(req)
	if err != nil {
		return
	}
	if resp.StatusCode == 401 {
		err = errors.New("permission denied, check cookies and 'ntag' header")
		return
	}
	defer resp.Body.Close()
	bytes, err = io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	return
}
