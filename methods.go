package nordnet_api_client

import (
	"encoding/json"
	"fmt"
)

// * * Endpoints for interacting with Nordnet * * //

// GetWatchLists : Returns the user's stock watch lists
func (nc NordnetClient) GetWatchLists() (wlg []WatchListDetail, err error) {
	var wlgList []WatchListCatalogResponse
	batchList := BatchList{
		BatchObject{
			RelativeURL: "watchlists",
			Method:      "GET",
		},
	}
	batch, err := createBatch(batchList)
	if err != nil {
		return
	}
	jsonBytes, err := json.Marshal(batch)
	if err != nil {
		return
	}
	bytes, err := nc.post("/batch", jsonBytes)
	if err != nil {
		return
	}
	err = json.Unmarshal(bytes, &wlgList)
	if err != nil {
		return
	}
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
	jsonBytes, err := json.Marshal(batch)
	if err != nil {
		return
	}
	bytes, err := nc.post("/batch", jsonBytes)
	if err != nil {
		return
	}
	err = json.Unmarshal(bytes, &watchListResponse)
	if err != nil {
		return
	}
	wlg = watchListResponse[0].WatchList
	return
}

// SearchInstrument : searches instrument by ID
func (nc NordnetClient) SearchInstrument(id int) (response InstrumentDetail, err error) {
	var instrumentResponse InstrumentResponse
	baseAddr := "/instrument_search/query/stocklist?apply_filters=instrument_id%3D"
	bytes, err := nc.get(fmt.Sprintf("%s%v", baseAddr, id))
	if err != nil {
		return
	}
	err = json.Unmarshal(bytes, &instrumentResponse)
	if err != nil {
		return
	}
	response = instrumentResponse.Results[0]
	return
}

// GetHistoricalPrices : get price history from specified date
func (nc NordnetClient) GetHistoricalPrices(id int, fromDate string) (history PriceHistory, err error) {
	var priceHistoryResponse PriceHistoryResponse
	baseAddr := fmt.Sprintf("instruments/historical/prices/%v", id)
	fields := "?fields=open,high,low,last,volume"
	from := fmt.Sprintf("&from=%s", fromDate) // format 2007-10-15
	batchList := BatchList{
		BatchObject{
			RelativeURL: fmt.Sprintf("%s%s%s", baseAddr, fields, from),
			Method:      "GET",
		},
	}
	batch, err := createBatch(batchList)
	if err != nil {
		return
	}
	jsonBytes, err := json.Marshal(batch)
	if err != nil {
		return
	}
	bytes, err := nc.post("/batch", jsonBytes)
	if err != nil {
		return
	}
	err = json.Unmarshal(bytes, &priceHistoryResponse)
	if err != nil {
		return
	}
	history = priceHistoryResponse[0].Body[0]
	return
}

// GetDividends : Get dividends for a company/stock
func (nc NordnetClient) GetDividends(id int) (div []Dividend, err error) {
	var dividendResponse DividendsResponse
	batchList := BatchList{
		BatchObject{
			RelativeURL: fmt.Sprintf("company_data/dividends/%v", id),
			Method:      "GET",
		},
	}
	batch, err := createBatch(batchList)
	if err != nil {
		return
	}
	jsonBytes, err := json.Marshal(batch)
	if err != nil {
		return
	}
	bytes, err := nc.post("/batch", jsonBytes)
	if err != nil {
		return
	}
	err = json.Unmarshal(bytes, &dividendResponse)
	if err != nil {
		return
	}
	div = dividendResponse[0].Body
	return
}
