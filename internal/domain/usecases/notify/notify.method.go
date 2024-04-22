package notify

import (
	"context"
	"github.com/ocistok-it/notification/internal/infrastructure/custerr"
	"github.com/ocistok-it/notification/internal/infrastructure/dto"
)

func (m *module) Send(ctx context.Context, request *dto.PushNotification) error {
	notifier, found := m.notifier[request.Service]
	if !found {
		return custerr.ErrServiceUnsupported()
	}

	message, err := m.initiator[request.Service](ctx, request.Metadata)
	if err != nil {
		return err
	}

	return notifier.Send(ctx, message)
}
