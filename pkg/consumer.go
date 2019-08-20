package kafkaexample

import (
	"context"
)

// Consumer an instance that consumes messages
type Consumer interface {
	// Read read into the stream
	Read(ctx context.Context, chMsg chan Message, chErr chan error)
}
