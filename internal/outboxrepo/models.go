// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package outboxrepo

import (
	"encoding/json"
	"time"
)

type FoundationOutboxEvent struct {
	ID        int64
	Topic     string
	Key       string
	Payload   []byte
	Headers   json.RawMessage
	CreatedAt time.Time
}
