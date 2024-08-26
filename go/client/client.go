package main

import (
	"log"
	"net/url"
	"os"
	"os/signal"

	"github.com/gorilla/websocket"
)

func main() {
	// Define the Websocket server URL
	u := url.URL{
		Scheme: "ws",
		Host:   "localhost:8080",
		Path:   "/conn",
	}
	log.Printf("Connecting to %s", u.String())

	// Connect to the WebSocket server
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("Dial Error :", err)
	}
	defer c.Close()

	// Create a channel to listen for OS signals
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	// Send a message to the server
	err = c.WriteMessage(websocket.TextMessage, []byte("as"))
	if err != nil {
		log.Fatal("Write message Error : ", err)
	}

	// Start the go routine for continously listen for message from the server
	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Printf("Read Error :", err)
				return
			}
			log.Printf("Recieved : %s", message)

		}
	}()

	// Use a ticker to periodically send ping messages to the server

	for {
		select {
		case <-done:
			return
		case <-interrupt:
			log.Println("Interrupt received, shutting down...")

			// Close the WebSocket connection cleanly
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("Close error:", err)
				return
			}

			select {
			case <-done:
			}
			return
		}
	}

}
