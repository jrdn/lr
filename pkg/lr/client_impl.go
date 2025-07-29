//go:generate go tool lazygen $GOFILE

package lr

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/jrdn/lr/pkg/wss"
)

// NewClient returns a new instance of Client.
func NewClient() (Client, error) {
	websocketConn, err := wss.Connect("ws://localhost:7682")
	if err != nil {
		return nil, err
	}

	params := []any{
		"github.com/jrdn/lr",
		"v0.0.1",
	}
	clientID := getClientID()
	if clientID != "" {
		params = append(params, clientID)
	}

	client := &clientImpl{c: websocketConn}
	resp, err := do[StringResponse](client, &Request{
		Message: "register",
		Params:  params,
	})
	if err != nil {
		return nil, err
	}

	clientID = resp.Response[0]
	err = saveClientID(clientID)
	if err != nil {
		return nil, err
	}

	websocketConn.SetCloseHandler(func(code int, text string) error {
		return client.reconnect()
	})

	return client, nil
}

var _ Client = &clientImpl{}

// clientImpl is a concrete implementation of the Client interface.
type clientImpl struct {
	c *websocket.Conn
}

func (c *clientImpl) reconnect() error {
	client, err := NewClient()
	if err != nil {
		return err
	}
	c.c = client.(*clientImpl).c
	return nil
}

func (c *clientImpl) Close(objectHandle string) error {
	return ErrNotImplemented
}

func (c *clientImpl) AppVersionString() (string, error) {
	resp, err := do[StringResponse](c, &Request{
		Message: "appVersionString",
	})
	if err != nil {
		return "", err
	}
	return resp.Response[0], nil
}

func (c *clientImpl) ShowView(view View) error {
	return doNoResponse(c, &Request{
		Message: "showView",
		Params:  params(view),
	})
}

func (c *clientImpl) ShowSource(source Source) error {
	return doNoResponse(c, &Request{
		Message: "showSource",
		Params:  params(source),
	})
}

func (c *clientImpl) SetFocusPoint(x, y float64, animate bool) error {
	fmt.Println("X: ", x, "Y: ", y, "Animate: ", animate)
	return doNoResponse(c, &Request{
		Message: "setFocusPoint",
		Params:  params(x, y, animate),
	})
}

func (c *clientImpl) ParameterNames() ([]string, error) {
	resp, err := do[StringArrResponse](c, &Request{
		Message: "getParameterNames",
	})
	if err != nil {
		return nil, err
	}
	return resp.Response[0], nil
}

func (c *clientImpl) ParameterType(name string) (ParameterType, error) {
	resp, err := do[StringResponse](c, &Request{
		Message: "getParameterType",
		Params:  params(name),
	})
	if err != nil {
		return "", err
	}
	return ParameterType(resp.Response[0]), err
}

func (c *clientImpl) ParameterDefault(name string) (*ParameterValue, error) {
	resp, err := do[AnyResponse](c, &Request{
		Message: "getDefault",
		Params:  params(name),
	})
	if err != nil {
		return nil, err
	}
	return &ParameterValue{resp.Response[0]}, err
}

func (c *clientImpl) Increment(name string, amount float64) error {
	return doNoResponse(c, &Request{
		Message: "increment",
		Params:  params(name, amount),
	})
}
func (c *clientImpl) Decrement(name string, amount float64) error {
	return doNoResponse(c, &Request{
		Message: "decrement",
		Params:  params(name, amount),
	})
}

func (c *clientImpl) ResetToDefault(name string) error {
	return doNoResponse(c, &Request{
		Message: "resetToDefault",
		Params:  params(name),
	})
}

func (c *clientImpl) GetValue(name string) (*ParameterValue, error) {
	//TODO implement me
	panic("implement me")
}

func (c *clientImpl) SetValue(name string, value *ParameterValue) error {
	//TODO implement me
	panic("implement me")
}

/* lazygen:template simpleOperation
func (c *clientImpl) {{.n}}() error {
	return doSimple(c, "{{.m}}")
}
*/
//lazygen:instance simpleOperation n=GoBack m=goBack
//lazygen:instance simpleOperation n=GoForward m=goForward

//lazygen:instance simpleOperation n=SelectAll m=selectAll
//lazygen:instance simpleOperation n=SelectNone m=selectNone

//lazygen:instance simpleOperation n=NextPhoto m=nextPhoto
//lazygen:instance simpleOperation n=PrevPhoto m=previousPhoto

//lazygen:instance simpleOperation n=ZoomIn m=zoomIn
//lazygen:instance simpleOperation n=ZoomOut m=zoomOut
//lazygen:instance simpleOperation n=ZoomInSome m=zoomInSome
//lazygen:instance simpleOperation n=ZoomOutSome m=zoomOutSome
//lazygen:instance simpleOperation n=Zoom100 m=zoomOneToOne
//lazygen:instance simpleOperation n=ZoomFit m=zoomToFit
//lazygen:instance simpleOperation n=ZoomFill m=zoomToFill
//lazygen:instance simpleOperation n=ZoomToggle m=toggleZoom

//lazygen:instance simpleOperation n=RotateLeft m=rotateLeft
//lazygen:instance simpleOperation n=RotateRight m=rotateRight
//lazygen:instance simpleOperation n=FlipHorizontal m=flipHorizontal
//lazygen:instance simpleOperation n=FlipVertical m=flipVertical
//lazygen:instance simpleOperation n=RotateAspectRatio m=rotateAspectRatio
//lazygen:instance simpleOperation n=ToggleConstrainAspectRatio m=toggleConstrainAspectRatio

//lazygen:instance simpleOperation n=Undo m=undo
//lazygen:instance simpleOperation n=Redo m=redo

//lazygen:instance simpleOperation n=FlagPick m=flagPick
//lazygen:instance simpleOperation n=FlagReject m=flagReject
//lazygen:instance simpleOperation n=FlagUnflag m=flagUnflag
//lazygen:instance simpleOperation n=FlagIncrease m=flagIncrease
//lazygen:instance simpleOperation n=FlagDecrease m=flagDecrease
//lazygen:instance simpleOperation n=RatingIncrease m=ratingIncrease
//lazygen:instance simpleOperation n=RatingDecrease m=ratingDecrease
//lazygen:instance simpleOperation n=RatingClear m=rating0
//lazygen:instance simpleOperation n=Rating0 m=rating0
//lazygen:instance simpleOperation n=Rating1 m=rating1
//lazygen:instance simpleOperation n=Rating2 m=rating2
//lazygen:instance simpleOperation n=Rating3 m=rating3
//lazygen:instance simpleOperation n=Rating4 m=rating4
//lazygen:instance simpleOperation n=Rating5 m=rating5

//lazygen:instance simpleOperation n=OpenEnhance m=openEnhance
//lazygen:instance simpleOperation n=OpenHDRMerge m=openHDRMerge
//lazygen:instance simpleOperation n=OpenPanoramaMerge m=openPanoramaMerge
//lazygen:instance simpleOperation n=OpenHDRPanoramaMerge m=openHDRPanoramaMerge
//lazygen:instance simpleOperation n=OpenExport m=openExport
//lazygen:instance simpleOperation n=EditInPhotoshop m=editInPhotoShop
//lazygen:instance simpleOperation n=OpenPreferences m=openPreferences

//lazygen:instance simpleOperation n=ClearFilters m=clearFilters
//lazygen:instance simpleOperation n=ShowFilters m=showFilters

//lazygen:instance simpleOperation n=SetAutoTone m=setAutoTone
