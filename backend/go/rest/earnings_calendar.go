package main

// EarningsCalendarQuery represents the query parameters for the earnings calendar API
type EarningsCalendarQuery struct {
	Symbol string `url:"symbol"`
	Offset int    `url:"offset"`
	Size   int    `url:"size"`
}

// EarningsCalendarResponse is not well-defined as it's scraped HTML.
// A struct might not be the best representation.
// Below is a placeholder for what the data might look like after parsing.
type EarningsCalendarEntry struct {
	Company       string
	Symbol        string
	EarningsDate  string
	EPSEstimate   float64
	ReportedEPS   float64
	Surprise      float64
}
