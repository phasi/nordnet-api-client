package nordnet_api_client

type InstrumentResponse struct {
	Rows      int                `json:"rows"`
	TotalHits int                `json:"total_hits"`
	Results   []InstrumentDetail `json:"results"`
}

type InstrumentDetail struct {
	InstrumentInfo struct {
		InstrumentID             int    `json:"instrument_id"`
		Name                     string `json:"name"`
		LongName                 string `json:"long_name"`
		Symbol                   string `json:"symbol"`
		InstrumentGroupType      string `json:"instrument_group_type"`
		InstrumentTypeHierarchy  string `json:"instrument_type_hierarchy"`
		InstrumentType           string `json:"instrument_type"`
		Isin                     string `json:"isin"`
		Currency                 string `json:"currency"`
		ClearingPlace            string `json:"clearing_place"`
		IsTradable               bool   `json:"is_tradable"`
		InstrumentPawnPercentage int    `json:"instrument_pawn_percentage"`
		IsShortable              bool   `json:"is_shortable"`
		IssuerID                 int    `json:"issuer_id"`
		IssuerName               string `json:"issuer_name"`
	} `json:"instrument_info"`
	StatusInfo struct {
		TradingStatus           string `json:"trading_status"`
		TranslatedTradingStatus string `json:"translated_trading_status"`
		TickTimestamp           int64  `json:"tick_timestamp"`
	} `json:"status_info"`
	MarketInfo struct {
		MarketID    int    `json:"market_id"`
		MarketSubID int    `json:"market_sub_id"`
		Identifier  string `json:"identifier"`
		TickSizeID  int    `json:"tick_size_id"`
	} `json:"market_info"`
	PriceInfo struct {
		Last struct {
			Price    float64 `json:"price"`
			Decimals int     `json:"decimals"`
		} `json:"last"`
		Open struct {
			Price    float64 `json:"price"`
			Decimals int     `json:"decimals"`
		} `json:"open"`
		Close struct {
			Price    float64 `json:"price"`
			Decimals int     `json:"decimals"`
		} `json:"close"`
		Turnover           float64 `json:"turnover"`
		TurnoverNormalized float64 `json:"turnover_normalized"`
		TurnoverVolume     int     `json:"turnover_volume"`
		Bid                struct {
			Price    float64 `json:"price"`
			Decimals int     `json:"decimals"`
		} `json:"bid"`
		Ask struct {
			Price    float64 `json:"price"`
			Decimals int     `json:"decimals"`
		} `json:"ask"`
		BidVolume int `json:"bid_volume"`
		AskVolume int `json:"ask_volume"`
		High      struct {
			Price    float64 `json:"price"`
			Decimals int     `json:"decimals"`
		} `json:"high"`
		Low struct {
			Price    float64 `json:"price"`
			Decimals int     `json:"decimals"`
		} `json:"low"`
		Diff struct {
			Diff     float64 `json:"diff"`
			Decimals int     `json:"decimals"`
		} `json:"diff"`
		DiffPct float64 `json:"diff_pct"`
		Spread  struct {
			Price    float64 `json:"price"`
			Decimals int     `json:"decimals"`
		} `json:"spread"`
		SpreadPct     float64 `json:"spread_pct"`
		TickTimestamp int64   `json:"tick_timestamp"`
		Realtime      bool    `json:"realtime"`
	} `json:"price_info"`
	ExchangeInfo struct {
		ExchangeCountry string `json:"exchange_country"`
	} `json:"exchange_info"`
	KeyRatiosInfo struct {
		Pe               float64 `json:"pe"`
		Ps               float64 `json:"ps"`
		Eps              float64 `json:"eps"`
		Pb               float64 `json:"pb"`
		DividendPerShare float64 `json:"dividend_per_share"`
		DividendYield    float64 `json:"dividend_yield"`
	} `json:"key_ratios_info"`
	HistoricalReturnsInfo struct {
		Yield1D  float64 `json:"yield_1d"`
		Yield1W  float64 `json:"yield_1w"`
		Yield1M  float64 `json:"yield_1m"`
		Yield3M  float64 `json:"yield_3m"`
		YieldYtd float64 `json:"yield_ytd"`
		Yield1Y  float64 `json:"yield_1y"`
		Yield3Y  float64 `json:"yield_3y"`
		Yield5Y  float64 `json:"yield_5y"`
		Realtime bool    `json:"realtime"`
	} `json:"historical_returns_info"`
	CompanyInfo struct {
		GeneralMeetingDate int64  `json:"general_meeting_date"`
		ReportDate         int64  `json:"report_date"`
		ReportType         string `json:"report_type"`
		ReportDescription  string `json:"report_description"`
	} `json:"company_info"`
	StatisticalInfo struct {
		StatisticsTimestamp int64 `json:"statistics_timestamp"`
		NumberOfOwners      int   `json:"number_of_owners"`
	} `json:"statistical_info"`
}
