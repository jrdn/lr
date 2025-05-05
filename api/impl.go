package api

import "github.com/gorilla/websocket"

type lightroom struct {
	clientGuid string
	conn       *websocket.Conn
}

func (lr *lightroom) Register() error {
	// TODO
	return nil
}
