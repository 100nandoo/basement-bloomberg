package main

// PriceHistoryQuery represents the query parameters for the price history API
type PriceHistoryQuery struct {
	Period1        int64  `url:"period1"`
	Period2        int64  `url:"period2"`
	Interval       string `url:"interval"`
	IncludePrePost bool   `url:"includePrePost"`
	Events         string `url:"events"`
}

// PriceHistoryResponse represents the structure of the response from the price history API
type PriceHistoryResponse struct {
	Chart struct {
		Result []struct {
			Meta struct {
				Currency           string  `json:"currency"`
				Symbol             string  `json:"symbol"`
				ExchangeName       string  `json:"exchangeName"`
				InstrumentType     string  `json:"instrumentType"`
				FirstTradeDate     int     `json:"firstTradeDate"`
				RegularMarketTime  int     `json:"regularMarketTime"`
				Gmtoffset          int     `json:"gmtoffset"`
				Timezone           string  `json:"timezone"`
				ExchangeTimezone   string  `json:"exchangeTimezoneName"`
				RegularMarketPrice float64 `json:"regularMarketPrice"`
				ChartPreviousClose float64 `json:"chartPreviousClose"`
				PreviousClose      float64 `json:"previousClose"`
				Scale              int     `json:"scale"`
				PriceHint          int     `json:"priceHint"`
			} `json:"meta"`
			Timestamp  []int64 `json:"timestamp"`
			Indicators struct {
				Quote []struct {
					Low    []float64 `json:"low"`
					Open   []float64 `json:"open"`
					Volume []int64   `json:"volume"`
					High   []float64 `json:"high"`
					Close  []float64 `json:"close"`
				} `json:"quote"`
				Adjclose []struct {
					Adjclose []float64 `json:"adjclose"`
				} `json:"adjclose"`
			} `json:"indicators"`
		} `json:"result"`
		Error interface{} `json:"error"`
	} `json:"chart"`
}
