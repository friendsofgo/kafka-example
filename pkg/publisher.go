package kafkaexample

import (
	"context"
)

// Publisher an instance that publish messages
type Publisher interface {
	// Publish publish a message into a stream
	Publish(ctx context.Context, payload interface{}) error
}