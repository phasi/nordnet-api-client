package nordnet_api_client

import "encoding/json"

// Batch : Representation of batch object
// for using with the Nordnet batch-api
type Batch struct {
	Batch string `json:"batch"`
}

// JSON : returns a JSON string representation of Batch{}
func (b Batch) JSON() string {
	jsonString, _ := json.Marshal(b)
	return string(jsonString)
}

// createBatch : Creates a correctly formatted Batch{}
// for using with the Nordnet batch-api
func createBatch(batchList BatchList) (batch Batch, err error) {
	jsonString, err := json.Marshal(batchList)
	if err != nil {
		return
	}
	batch.Batch = string(jsonString)
	return
}

// getBatchListFromBatch : deserializes batchlist and returns it
func getBatchListFromBatch(batch Batch) (batchList BatchList, err error) {
	err = json.Unmarshal([]byte(batch.Batch), &batchList)
	return
}

// BatchObject : Representation of single batch request object
type BatchObject struct {
	RelativeURL string `json:"relative_url"`
	Method      string `json:"method"`
}

// BatchList : List of batch requests
type BatchList []BatchObject
