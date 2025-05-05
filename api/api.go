package api

import "github.com/gorilla/websocket"

type View string

const (
	LoupeView       View = "loupe"
	SquareGridView  View = "square_grid"
	DynamicGridView View = "dynamic_grid"
	CompareView     View = "compare"
)

type Lightroom interface {
	Register() error
	//Version() string
	//ShowView(v View)
	//
	//GoBack()
	//GoForward()
	//
	//SelectAll()
	//SelectNone()
}

func Connect() (Lightroom, error) {
	return ConnectURL("ws://127.0.0.1:7682")
}

func ConnectURL(url string) (Lightroom, error) {
	c, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		return nil, err
	}

	lr := &lightroom{
		conn: c,
	}

	err = lr.Register()
	if err != nil {
		return nil, err
	}

	return lr, nil
}
