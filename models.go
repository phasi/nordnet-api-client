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

type WatchListCatalogResponse struct {
	Code    int `json:"code"`
	Headers []struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"headers"`
	WatchListDetail []WatchListDetail `json:"body"`
}

type WatchListDetail struct {
	ListID   int    `json:"list_id"`
	ListName string `json:"list_name"`
}

type WatchListResponse struct {
	Code    int `json:"code"`
	Headers []struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"headers"`
	WatchList WatchList `json:"body"`
}

type WatchList struct {
	IsFull           bool `json:"is_full"`
	TotalItemsInList int  `json:"total_items_in_list"`
	Results          []struct {
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
		} `json:"instrument_info"`
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
			TurnoverVolume     int     `json:"turnover_volume"`
			Turnover           float64 `json:"turnover"`
			TurnoverNormalized float64 `json:"turnover_normalized"`
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
		FundInfo struct {
		} `json:"fund_info"`
		AnnualGrowthInfo struct {
			AnnualGrowth1Y float64 `json:"annual_growth_1y"`
		} `json:"annual_growth_info"`
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
		IndicatorInfo struct {
		} `json:"indicator_info"`
		WatchlistMeta struct {
			DisplayOrder int `json:"display_order"`
			ItemID       int `json:"item_id"`
		} `json:"watchlist_meta"`
	} `json:"results"`
}

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

type DividendsResponse []struct {
	Code    int `json:"code"`
	Headers []struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"headers"`
	Body []Dividend `json:"body"`
}

type Dividend struct {
	ExcludingDate string  `json:"excluding_date"`
	PayDate       string  `json:"pay_date"`
	AmountPaid    float64 `json:"amount_paid"`
	Currency      string  `json:"currency"`
	Type          string  `json:"type"`
}

type KeyFiguresResponse []struct {
	Code    int `json:"code"`
	Headers []struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"headers"`
	Body KeyFigures `json:"body"`
}

type KeyFigures struct {
	ForwardPeratio            float64 `json:"forward_peratio"`
	PriceBookings             float64 `json:"price_bookings"`
	PriceSales                float64 `json:"price_sales"`
	DividendPerShare          float64 `json:"dividend_per_share"`
	EarningsPerShare          float64 `json:"earnings_per_share"`
	ForwardDividendYield      float64 `json:"forward_dividend_yield"`
	OneYearEstimatedEpsGrowth float64 `json:"one_year_estimated_eps_growth"`
}

type InstrumentPriceResponse []struct {
	Code    int `json:"code"`
	Headers []struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"headers"`
	Body []InstrumentPrice `json:"body"`
}
type InstrumentPrice struct {
	InstrumentID   int     `json:"instrument_id"`
	Delay          int     `json:"delay"`
	Ask            float64 `json:"ask"`
	Bid            float64 `json:"bid"`
	BidVolume      int     `json:"bid_volume"`
	AskVolume      int     `json:"ask_volume"`
	TickTimestamp  int64   `json:"tick_timestamp"`
	TradeTimestamp int64   `json:"trade_timestamp"`
	Last           float64 `json:"last"`
	LastVolume     int     `json:"last_volume"`
	High           float64 `json:"high"`
	Low            float64 `json:"low"`
	Open           float64 `json:"open"`
	Close          float64 `json:"close"`
	Vwap           float64 `json:"vwap"`
	Turnover       float64 `json:"turnover"`
	TurnoverVolume int     `json:"turnover_volume"`
	LotSize        int     `json:"lot_size"`
}
