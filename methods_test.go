package nordnet_api_client

import (
	"testing"
)

// mockDriver :
var mockDriver MockDriver = MockDriver{
	MockDBFolderPath: "./mocks",
}

// fakeClient :
var fakeClient NordnetClient = NordnetClient{
	Driver: mockDriver,
}

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

func TestGetDividends(t *testing.T) {
	dividends, err := fakeClient.GetDividends(1234567)
	if err != nil {
		t.Error(err)
	}
	t.Log(dividends)
}

func TestGetKeyFigures(t *testing.T) {
	keyfigures, err := fakeClient.GetKeyFigures(1234567)
	if err != nil {
		t.Error(err)
	}
	t.Log(keyfigures)
}
