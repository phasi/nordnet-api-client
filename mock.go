package nordnet_api_client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"
)

var endpoints map[string]string = map[string]string{
	"/instrument_search/query/stocklist?apply_filters=instrument_id%3D": "/instrument.json",
}

var batchFiles map[string]string = map[string]string{
	"^watchlists$":                   "/watchlists.json",
	"^watchlists/[0-9]{1,}/alldata$": "/watchlistresponse.json",
	"^instruments/historical/prices/[0-9]{1,}[?]fields=([a-z]{1,}[,]{0,}){1,}[&]from=[0-9]{4}-[0-9]{2}-[0-9]{2}$": "/price_history.json",
	"company_data/dividends/[0-9]{1,}$":  "/dividends.json",
	"company_data/keyfigures/[0-9]{1,}$": "/keyfigures.json",
}

func handleBatch(data []byte) (s string, err error) {
	var batch Batch
	err = json.Unmarshal(data, &batch)
	if err != nil {
		return
	}
	batchList, err := getBatchListFromBatch(batch)
	if err != nil {
		return
	}
	s, err = checkRegex(batchList[0].RelativeURL)
	if err != nil {
		return
	}
	return
}

func checkRegex(path string) (string, error) {
	for k, v := range batchFiles {
		regx, _ := regexp.Compile(k)
		if regx.MatchString(path) {
			return v, nil
		}
	}
	return "", fmt.Errorf("mock response for path %s not found", path)
}

func checkPath(path string) (string, error) {
	for k, v := range endpoints {
		if strings.HasPrefix(path, k) {
			return v, nil
		}
	}
	return "", fmt.Errorf("mock response for path %s not found", path)
}

func byteReader(file string) (contents []byte, err error) {
	contents, err = os.ReadFile(file)
	return
}

// MockDriver :
type MockDriver struct {
	MockDBFolderPath string
}

func (mockDriver MockDriver) getFullPath(relativePath string) string {
	return fmt.Sprintf("%s%s", mockDriver.MockDBFolderPath, relativePath)
}

// PrepareGetRequest :
func (mockDriver MockDriver) PrepareGetRequest(path string) (req *http.Request, err error) {
	req, err = http.NewRequest("GET", path, nil)
	return
}

// PreparePostRequest :
func (mockDriver MockDriver) PreparePostRequest(path string, data []byte) (req *http.Request, err error) {
	req, err = http.NewRequest("POST", path, bytes.NewBuffer(data))
	return
}

// SendRequest : Sends Mock Request
func (mockDriver MockDriver) SendRequest(req *http.Request) (bytes []byte, err error) {
	bytes, err = mockDriver.handler(req)
	return
}

func (mockDriver MockDriver) handler(req *http.Request) ([]byte, error) {
	var bytes []byte
	if req.Method == "POST" {
		data, err := io.ReadAll(req.Body)
		if err != nil {
			return bytes, err
		}
		if req.URL.RequestURI() == "/batch" {
			filepath, err := handleBatch(data)
			if err != nil {
				return bytes, err
			}
			bytes, err = byteReader(mockDriver.getFullPath(filepath))
			return bytes, err
		}
	} else if req.Method == "GET" {
		filepath, err := checkPath(req.URL.RequestURI())
		if err != nil {
			return bytes, err
		}
		bytes, err := byteReader(mockDriver.getFullPath(filepath))
		if err != nil {
			return bytes, err
		}
		return bytes, nil
	}
	return bytes, errors.New("not supported method")
}
