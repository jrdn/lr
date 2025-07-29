package wss

import (
	"log"
	"net/url"

	"github.com/gorilla/websocket"
)

func Connect(serverURL string) (*websocket.Conn, error) {
	u, err := url.Parse(serverURL)
	if err != nil {
		return nil, err
	}

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Printf("dial error: %v", err)
		return nil, err
	}
	log.Printf("connected to %s", serverURL)
	return c, nil
}
