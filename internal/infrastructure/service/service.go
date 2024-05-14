package service

import (
	"context"
	"github.com/ocistok-it/notification/internal/infrastructure/enums"
)

type (
	// Service encapsulates the various client.
	Service interface {
		GetService() enums.Channel
		Send(ctx context.Context, message interface{}) error
	}
)
