package lr

import (
	"errors"

	"github.com/beevik/guid"
)

var (
	ErrNotImplemented     = errors.New("not implemented")
	ErrRegistrationFailed = errors.New("registration failed")
)

type RequestID string

func NewRequest() *Request {
	return &Request{
		RequestID: NewID(),
	}
}

type Request struct {
	RequestID string  `json:"requestId"`
	Object    *string `json:"object"`
	Message   string  `json:"message"`
	Params    []any   `json:"params"`
}

func params(x ...any) []any {
	return x
}

type Response[T any] struct {
	RequestID string
	Success   bool
	Response  T
}

type View string

var (
	ViewLoupe      View = "loupe"
	ViewGrid       View = "dynamic_grid"
	ViewSquareGrid View = "square_grid"
	ViewCompare    View = "compare"
)

func ParseView(s string) (View, error) {
	switch s {
	case "loupe":
		return ViewLoupe, nil
	case "grid":
		return ViewGrid, nil
	case "square":
		return ViewSquareGrid, nil
	case "compare":
		return ViewCompare, nil
	}
	return "", errors.New("unknown view")
}

type Source string

var (
	SourceAll   Source = "all_photos"
	SourceTrash Source = "trash"
)

func ParseSource(s string) (Source, error) {
	switch s {
	case "all_photos", "all":
		return SourceAll, nil
	case "trash":
		return SourceTrash, nil
	}
	return "", errors.New("unknown source")
}

type ParameterName string

type ParameterType string

var (
	ParameterTypeBoolean ParameterType = "boolean"
	ParameterTypeNumber  ParameterType = "number"
	ParameterTypeString  ParameterType = "string"
)

func NewID() string {
	return guid.New().String()
}

func NewParameterValue(x any) *ParameterValue {
	return &ParameterValue{x: x}
}

type ParameterValue struct {
	x       any
	Number  *float64 `json:"number,omitempty"`
	Boolean *bool    `json:"boolean,omitempty"`
	String  *string  `json:"string,omitempty"`
}

func (pv ParameterValue) GetNumber() (float64, error) {
	if pv.Number != nil {
		return *pv.Number, nil
	}
	if n, ok := pv.x.(float64); ok {
		return n, nil
	}
	return 0, errors.New("not a number")
}

func (pv ParameterValue) GetBoolean() (bool, error) {
	if pv.Boolean != nil {
		return *pv.Boolean, nil
	}
	if b, ok := pv.x.(bool); ok {
		return b, nil
	}
	return false, errors.New("not a boolean")
}

func (pv ParameterValue) GetString() (string, error) {
	if pv.String != nil {
		return *pv.String, nil
	}
	if s, ok := pv.x.(string); ok {
		return s, nil
	}
	return "", errors.New("not a string")
}

func (pv ParameterValue) Value() (any, error) {
	if pv.Number != nil {
		return *pv.Number, nil
	}
	if pv.Boolean != nil {
		return *pv.Boolean, nil
	}
	if pv.String != nil {
		return *pv.String, nil
	}
	return pv.x, nil
}
