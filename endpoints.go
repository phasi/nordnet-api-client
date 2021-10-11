package nordnet_api_client

import (
	"encoding/json"
	"fmt"
)

// * * Endpoints for interacting with Nordnet * * //

// GetWatchLists : Returns the user's stock watch list
func (nc NordnetClient) GetWatchLists() (wlg []WatchListDetail, err error) {
	var wlgList []WatchListCatalogResponse
	batchList := BatchList{
		BatchObject{
			RelativeURL: "watchlists", //2800292/alldata
			Method:      "GET",
		},
	}
	batch, err := createBatch(batchList)
	if err != nil {
		return
	}
	req, err := nc.Driver.PreparePostRequest("/batch", batch)
	if err != nil {
		return
	}
	bytes, err := nc.Driver.SendRequest(req)
	if err != nil {
		return
	}
	err = json.Unmarshal(bytes, &wlgList)
	wlg = wlgList[0].WatchListDetail
	return
}

// GetWatchList : Get single stock watchlist
func (nc NordnetClient) GetWatchList(watchListID int) (wlg WatchList, err error) {
	var watchListResponse []WatchListResponse
	batchList := BatchList{
		BatchObject{
			RelativeURL: fmt.Sprintf("watchlists/%v/alldata", watchListID),
			Method:      "GET",
		},
	}
	batch, err := createBatch(batchList)
	if err != nil {
		return
	}
	req, err := nc.Driver.PreparePostRequest("/batch", batch)
	if err != nil {
		return
	}
	bytes, err := nc.Driver.SendRequest(req)
	if err != nil {
		return
	}
	err = json.Unmarshal(bytes, &watchListResponse)
	wlg = watchListResponse[0].WatchList
	return
}

func (nc NordnetClient) SearchInstrument(id int) (response InstrumentDetail, err error) {
	var instrumentResponse InstrumentResponse
	baseAddr := "/instrument_search/query/stocklist?apply_filters=instrument_id%3D"
	req, err := nc.Driver.PrepareGetRequest(fmt.Sprintf("%s%v", baseAddr, id))
	if err != nil {
		return
	}
	bytes, err := nc.Driver.SendRequest(req)
	if err != nil {
		return
	}
	err = json.Unmarshal(bytes, &instrumentResponse)
	response = instrumentResponse.Results[0]
	return
}
