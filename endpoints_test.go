package nordnet_api_client

import (
	"testing"
)

// mockDriver :
var mockDriver MockDriver = MockDriver{
	MockDBFolderPath: "./mocks",
}

// fakeClient :
var fakeClient NordnetClient = NewNordnetClient(mockDriver)

func TestGetWatchLists(t *testing.T) {
	watchListDetail, err := fakeClient.GetWatchLists()
	if err != nil {
		t.Error(err)
	}
	t.Log(watchListDetail)
}

func TestGetWatchList(t *testing.T) {
	watchListResponse, err := fakeClient.GetWatchList(2800292)
	if err != nil {
		t.Error(err)
	}
	t.Log(watchListResponse)
}

func TestSearchInstrument(t *testing.T) {
	instrument, err := fakeClient.SearchInstrument(12345)
	if err != nil {
		t.Error(err)
	}
	t.Log(instrument)
}

func TestGetHistoricalPrices(t *testing.T) {
	prices, err := fakeClient.GetHistoricalPrices(1, "2021-09-30")
	if err != nil {
		t.Error(err)
	}
	t.Log(prices)
}
