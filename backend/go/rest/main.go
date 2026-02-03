package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
)

const (
	query1URL = "https://query1.finance.yahoo.com"
	query2URL = "https://query2.finance.yahoo.com"
	rootURL   = "https://finance.yahoo.com"
)

// getPriceHistory retrieves historical price data for a given ticker.
func getPriceHistory(ticker string, params PriceHistoryQuery) (*PriceHistoryResponse, error) {
	v, _ := query.Values(params)
	apiURL := fmt.Sprintf("%s/v8/finance/chart/%s?%s", query2URL, ticker, v.Encode())

	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var history PriceHistoryResponse
	if err := json.Unmarshal(body, &history); err != nil {
		return nil, err
	}

	return &history, nil
}

// getQuoteSummary retrieves a summary of various data modules for a given stock symbol.
func getQuoteSummary(ticker string, params QuoteSummaryQuery) (*QuoteSummaryResponse, error) {
	v, _ := query.Values(params)
	apiURL := fmt.Sprintf("%s/v10/finance/quoteSummary/%s?%s", query2URL, ticker, v.Encode())

	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var summary QuoteSummaryResponse
	if err := json.Unmarshal(body, &summary); err != nil {
		return nil, err
	}

	return &summary, nil
}

// getFundamentals fetches time series data for financial statements.
func getFundamentals(ticker string, params FundamentalsQuery) (*FundamentalsResponse, error) {
	v, _ := query.Values(params)
	apiURL := fmt.Sprintf("%s/ws/fundamentals-timeseries/v1/finance/timeseries/%s?%s", query2URL, ticker, v.Encode())

	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var fundamentals FundamentalsResponse
	if err := json.Unmarshal(body, &fundamentals); err != nil {
		return nil, err
	}

	return &fundamentals, nil
}

// getEarningsCalendar retrieves earnings calendar data by scraping the HTML content.
func getEarningsCalendar(ticker string, params EarningsCalendarQuery) (string, error) {
	v, _ := query.Values(params)
	apiURL := fmt.Sprintf("%s/calendar/earnings?%s", rootURL, v.Encode())

	resp, err := http.Get(apiURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// Note: The response is HTML. You would need a library like goquery to parse it.
	// For demonstration, we're just returning the HTML as a string.
	return string(body), nil
}

func main() {
	// Example usage of the functions
	// You can uncomment and run these examples.

	// --- Get Price History ---
	// historyParams := PriceHistoryQuery{
	// 	Period1:        1609459200, // 2021-01-01
	// 	Period2:        1640995200, // 2022-01-01
	// 	Interval:       "1d",
	// 	IncludePrePost: true,
	// 	Events:         "div,splits",
	// }
	// history, err := getPriceHistory("AAPL", historyParams)
	// if err != nil {
	// 	log.Fatalf("Error getting price history: %v", err)
	// }
	// fmt.Printf("Price History for AAPL: %+v\n", history.Chart.Result[0].Meta)


	// --- Get Quote Summary ---
	// summaryParams := QuoteSummaryQuery{
	// 	Modules:    "assetProfile,summaryDetail",
	// 	CorsDomain: "finance.yahoo.com",
	// 	Formatted:  false,
	// 	Symbol:     "AAPL",
	// }
	// summary, err := getQuoteSummary("AAPL", summaryParams)
	// if err != nil {
	// 	log.Fatalf("Error getting quote summary: %v", err)
	// }
	// fmt.Printf("Quote Summary for AAPL: %+v\n", summary.QuoteSummary.Result[0].AssetProfile)


	// --- Get Fundamentals ---
	// fundParams := FundamentalsQuery{
	// 	Symbol:  "AAPL",
	// 	Type:    "annualNormalizedEBITDA,annualTaxEffectOfUnusualItems",
	// 	Period1: 1609459200,
	// 	Period2: 1640995200,
	// }
	// fundamentals, err := getFundamentals("AAPL", fundParams)
	// if err != nil {
	// 	log.Fatalf("Error getting fundamentals: %v", err)
	// }
	// fmt.Printf("Fundamentals for AAPL: %+v\n", fundamentals.Timeseries.Result[0])


	// --- Get Earnings Calendar ---
	// earningsParams := EarningsCalendarQuery{
	// 	Symbol: "AAPL",
	// 	Size:   5,
	// }
	// earningsHTML, err := getEarningsCalendar("AAPL", earningsParams)
	// if err != nil {
	// 	log.Fatalf("Error getting earnings calendar: %v", err)
	// }
	// fmt.Println("Earnings Calendar HTML for AAPL (first 500 chars):", earningsHTML[:500])
}
