package global

import (
	"encoding/json"
	"time"
)

type EventDTO[T any] struct {
	Type     string    `json:"type"`
	IssuedAt time.Time `json:"isat"`
	Data     T         `json:"data"`
}

func (ev EventDTO[T]) MarshalBinary() ([]byte, error) {
	return json.Marshal(ev)
}
