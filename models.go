package nordnet_api_client

// PriceHistoryResponse : HTTP response for price history
type PriceHistoryResponse []struct {
	Code    int `json:"code"`
	Headers []struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"headers"`
	Body []PriceHistory `json:"body"`
}

// PriceHistory : Price history day by day
type PriceHistory struct {
	InstrumentID int `json:"instrument_id"`
	Delayed      int `json:"delayed"`
	Prices       []struct {
		Time   int64   `json:"time"`
		High   float64 `json:"high"`
		Low    float64 `json:"low"`
		Open   float64 `json:"open"`
		Last   float64 `json:"last"`
		Volume int     `json:"volume"`
	} `json:"prices"`
}
