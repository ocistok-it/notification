package usecases

import (
	"context"
	"github.com/ocistok-it/notification/internal/infrastructure/dto"
)

type (
	NotifyUsecase interface {
		Send(ctx context.Context, request *dto.PushNotification) error
	}
)
