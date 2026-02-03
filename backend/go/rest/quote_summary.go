package main

// QuoteSummaryQuery represents the query parameters for the quote summary API
type QuoteSummaryQuery struct {
	Modules      string `url:"modules"`
	CorsDomain   string `url:"corsDomain"`
	Formatted    bool   `url:"formatted"`
	Symbol       string `url:"symbol"`
}

// QuoteSummaryResponse represents the structure of the response from the quote summary API
type QuoteSummaryResponse struct {
	QuoteSummary struct {
		Result []struct {
			AssetProfile struct {
				Address1             string `json:"address1"`
				City                 string `json:"city"`
				State                string `json:"state"`
				Zip                  string `json:"zip"`
				Country              string `json:"country"`
				Phone                string `json:"phone"`
				Website              string `json:"website"`
				Industry             string `json:"industry"`
				Sector               string `json:"sector"`
				LongBusinessSummary  string `json:"longBusinessSummary"`
				FullTimeEmployees    int    `json:"fullTimeEmployees"`
				CompanyOfficers      []struct {
					MaxAge      int    `json:"maxAge"`
					Name        string `json:"name"`
					Age         int    `json:"age"`
					Title       string `json:"title"`
					YearBorn    int    `json:"yearBorn"`
					FiscalYear  int    `json:"fiscalYear"`
					TotalPay    int    `json:"totalPay"`
					ExercisedValue int `json:"exercisedValue"`
					UnexercisedValue int `json:"unexercisedValue"`
				} `json:"companyOfficers"`
				MaxAge int `json:"maxAge"`
			} `json:"assetProfile"`
			SummaryDetail struct {
				PreviousClose          float64 `json:"previousClose"`
				Open                   float64 `json:"open"`
				DayLow                 float64 `json:"dayLow"`
				DayHigh                float64 `json:"dayHigh"`
				RegularMarketPreviousClose float64 `json:"regularMarketPreviousClose"`
				RegularMarketOpen      float64 `json:"regularMarketOpen"`
				RegularMarketDayLow    float64 `json:"regularMarketDayLow"`
				RegularMarketDayHigh   float64 `json:"regularMarketDayHigh"`
				DividendRate           float64 `json:"dividendRate"`
				DividendYield          float64 `json:"dividendYield"`
				ExDividendDate         int     `json:"exDividendDate"`
				PayoutRatio            float64 `json:"payoutRatio"`
				FiveYearAvgDividendYield float64 `json:"fiveYearAvgDividendYield"`
				Beta                   float64 `json:"beta"`
				TrailingPE             float64 `json:"trailingPE"`
				ForwardPE              float64 `json:"forwardPE"`
				Volume                 int     `json:"volume"`
				RegularMarketVolume    int     `json:"regularMarketVolume"`
				AverageVolume          int     `json:"averageVolume"`
				AverageVolume10days    int     `json:"averageVolume10days"`
				AverageDailyVolume10Day int    `json:"averageDailyVolume10Day"`
				Bid                    float64 `json:"bid"`
				Ask                    float64 `json:"ask"`
				BidSize                int     `json:"bidSize"`
				AskSize                int     `json:"askSize"`
				MarketCap              int64   `json:"marketCap"`
				FiftyTwoWeekLow        float64 `json:"fiftyTwoWeekLow"`
				FiftyTwoWeekHigh       float64 `json:"fiftyTwoWeekHigh"`
				PriceToSalesTrailing12Months float64 `json:"priceToSalesTrailing12Months"`
				FiftyDayAverage        float64 `json:"fiftyDayAverage"`
				TwoHundredDayAverage   float64 `json:"twoHundredDayAverage"`
				TrailingAnnualDividendRate float64 `json:"trailingAnnualDividendRate"`
				TrailingAnnualDividendYield float64 `json:"trailingAnnualDividendYield"`
				Currency               string  `json:"currency"`
			} `json:"summaryDetail"`
		} `json:"result"`
		Error interface{} `json:"error"`
	} `json:"quoteSummary"`
}
