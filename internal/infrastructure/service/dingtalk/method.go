package dingtalk

import (
	"context"
	"github.com/ocistok-it/notification/internal/infrastructure/custerr"
	"github.com/ocistok-it/notification/internal/infrastructure/enums"
)

func (s *Service) GetService() enums.ServiceName {
	return enums.DingTalkService
}

func (s *Service) Send(ctx context.Context, message interface{}) error {
	payload, ok := message.(*Message)
	if !ok {
		return custerr.New("invalid_request", "failed to cast message")
	}

	// TODO (rizal) : implement others MessageType
	switch payload.Type {
	case TypeText:
		return s.robotSend(payload.AccessToken, payload.TextMessage())
	}

	return custerr.New("invalid_type", "message type not supported")
}
