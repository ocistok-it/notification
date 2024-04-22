package event

import (
	"context"
)

type (
	Consumer interface {
		Use(middlewares ...Middleware)
		Consume(event string, handler Handler)
		Start()
		Error() <-chan error
		Close() error
	}

	Middleware func(next Handler) Handler
	Handler    func(ctx context.Context, body []byte) error
)
