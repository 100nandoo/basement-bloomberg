package websocket

// LiveDataRequest represents the subscription message for the live data WebSocket
type LiveDataRequest struct {
	Subscribe []string `json:"subscribe"`
}

// LiveDataResponse represents the structure of a message from the live data WebSocket
type LiveDataResponse struct {
	ID            string  `json:"id"`
	Price         float64 `json:"price"`
	Time          int64   `json:"time"`
	Currency      string  `json:"currency"`
	Exchange      string  `json:"exchange"`
	QuoteType     string  `json:"quoteType"`
	MarketHours   string  `json:"marketHours"`
	ChangePercent float64 `json:"changePercent"`
	DayVolume     int64   `json:"dayVolume"`
	Change        float64 `json:"change"`
	PriceHint     int     `json:"priceHint"`
}
