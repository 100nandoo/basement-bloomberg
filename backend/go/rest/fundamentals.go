package main

// FundamentalsQuery represents the query parameters for the fundamentals API
type FundamentalsQuery struct {
	Symbol  string `url:"symbol"`
	Type    string `url:"type"`
	Period1 int64  `url:"period1"`
	Period2 int64  `url:"period2"`
}

// FundamentalsResponse represents the structure of the response from the fundamentals API
type FundamentalsResponse struct {
	Timeseries struct {
		Result []struct {
			Meta struct {
				Symbol string `json:"symbol"`
				Type   string `json:"type"`
			} `json:"meta"`
			Timestamp []int64 `json:"timestamp"`
			AnnualTaxEffectOfUnusualItems []struct {
				AsOfDate string `json:"asOfDate"`
				ReportedValue struct {
					Raw int64 `json:"raw"`
					Fmt string `json:"fmt"`
				} `json:"reportedValue"`
			} `json:"annualTaxEffectOfUnusualItems,omitempty"`
			AnnualNormalizedEBITDA []struct {
				AsOfDate string `json:"asOfDate"`
				ReportedValue struct {
					Raw int64 `json:"raw"`
					Fmt string `json:"fmt"`
				} `json:"reportedValue"`
			} `json:"annualNormalizedEBITDA,omitempty"`
			// Add other financial statement keys here as needed
		} `json:"result"`
		Error interface{} `json:"error"`
	} `json:"timeseries"`
}
