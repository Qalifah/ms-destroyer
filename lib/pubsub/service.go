package pubsub

import "context"

type Service interface {
	Publish(ctx context.Context, data []byte) error
}
