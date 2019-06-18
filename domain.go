package producer

import (
	"context"
)

// Producer is the interface used for sending all events to a destination.
type Producer interface {
	// Produce ships the event to the destination and returns the final
	// version of the data sent.
	Produce(ctx context.Context, event interface{}) (interface{}, error)
}
