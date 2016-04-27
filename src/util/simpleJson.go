package util

import (
	"encoding/json"
)

type Message struct {
	Name string
	Body string
	Time int64
}

func toJson(m Message) ([]byte, error) {
	b, err := json.Marshal(m)
	if err == nil {
		return b,err
	}
	return nil,err
}