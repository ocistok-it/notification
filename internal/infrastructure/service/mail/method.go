package mail

import (
	"context"
	"github.com/ocistok-it/notification/internal/infrastructure/custerr"
	"github.com/ocistok-it/notification/internal/infrastructure/enums"
	"gopkg.in/gomail.v2"
)

func (s *Service) GetService() enums.Channel {
	return enums.MailService
}

func (s *Service) Send(ctx context.Context, message interface{}) error {
	payload, ok := message.(*Message)
	if !ok {
		return custerr.New("invalid_request", "failed to cast message")
	}

	msg := s.messageBuilder(*payload)

	if err := s.client.DialAndSend(msg); err != nil {
		return custerr.New("error_send_mail", err.Error())
	}

	return nil
}

func (s *Service) messageBuilder(msg Message) *gomail.Message {
	message := gomail.NewMessage()
	message.SetAddressHeader("From", s.from, s.fromName)
	message.SetHeader("To", msg.To...)
	message.SetHeader("Cc", msg.Cc...)
	message.SetHeader("Subject", msg.Subject)
	message.SetBody(DefaultMime, msg.Message)

	return message
}
