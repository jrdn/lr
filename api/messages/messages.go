package messages

import "crypto/rand"

func NewRequestID() string {
	return rand.Text()
}

type Message struct {
	RequestID string   `json:"requestId"`
	Object    *string  `json:"object"`
	Message   string   `json:"message"`
	Params    []string `json:"params"`
}

func NewRegisterMessage(appName string, appVersion string, clientGuid string) *Message {
	params := []string{appName, appVersion}
	if clientGuid != "" {
		params = append(params, clientGuid)
	}

	return &Message{
		RequestID: NewRequestID(),
		Object:    nil,
		Message:   "register",
		Params:    params,
	}
}
