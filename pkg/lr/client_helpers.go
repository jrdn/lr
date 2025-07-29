package lr

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
)

func do[T any](c *clientImpl, r *Request) (*Response[T], error) {
	if r.RequestID == "" {
		r.RequestID = NewID()
	}

	msg, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}

	err = c.c.WriteMessage(websocket.TextMessage, msg)
	if err != nil {
		return nil, err
	}

	t, d, err := c.c.ReadMessage()
	if err != nil {
		return nil, err
	}

	if t != websocket.TextMessage {
		return nil, fmt.Errorf("unexpected message type: %d", t)
	}

	var data Response[T]
	err = json.Unmarshal(d, &data)
	if err != nil {
		return nil, err
	}

	if !data.Success {
		return nil, fmt.Errorf("Failed: %+v", data.Response)
	}

	return &data, nil
}

func doNoResponse(c *clientImpl, r *Request) error {
	_, err := do[NoResponse](c, r)
	return err
}

func doSimple(c *clientImpl, msg string) error {
	return doNoResponse(c, &Request{Message: msg})
}
