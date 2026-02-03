package main

import (
	"fmt"
	"log"
	"net/url"

	"github.com/gorilla/websocket"
)

const (
	wsURL = "wss://streamer.finance.yahoo.com/?version=2"
)

// streamLiveData connects to the WebSocket and streams live data.
func streamLiveData(request LiveDataRequest) {
	u, _ := url.Parse(wsURL)
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	if err := c.WriteJSON(request); err != nil {
		log.Println("write:", err)
		return
	}

	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			return
		}

		// The message is base64 encoded protobuf. For simplicity, we just print it.
		// For a real application, you would decode it using the .proto file.
		fmt.Printf("recv: %s\n", message)
	}
}

func main() {
	// --- Stream Live Data ---
	liveRequest := LiveDataRequest{
		Subscribe: []string{"AAPL", "GOOG"},
	}
	streamLiveData(liveRequest)
}
