package rest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

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
	resp, err := GetClient().Get(apiURL)
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

// GetQuoteSummary retrieves a summary of various data modules for a given stock symbol.
func GetQuoteSummary(ticker string, params QuoteSummaryQuery) (*QuoteSummaryResponse, error) {
	v, _ := query.Values(params)
	apiURL := fmt.Sprintf("%s/v10/finance/quoteSummary/%s?%s", query2URL, ticker, v.Encode())

	resp, err := GetClient().Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	// fmt.Printf("API URL: %s\nResponse Body: %s\n", apiURL, string(body))

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

	resp, err := GetClient().Get(apiURL)
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

	resp, err := GetClient().Get(apiURL)
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
