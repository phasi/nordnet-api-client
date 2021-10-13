package nordnet_api_client

import (
	"net/http"
)

// ClientDriver : Interface for creating and sending requests.
// You may write your own driver to implement this interface if
// for example you're developing and you need to mock API responses.
type ClientDriver interface {
	SendRequest(req *http.Request) (bytes []byte, err error)
	PrepareGetRequest(path string) (req *http.Request, err error)
	PreparePostRequest(path string, jsonBytes []byte) (req *http.Request, err error)
}
