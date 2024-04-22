package mail

import (
	"context"
	"github.com/ocistok-it/notification/internal/infrastructure/custerr"
	"github.com/ocistok-it/notification/internal/infrastructure/enums"
	"gopkg.in/gomail.v2"
)

func (s *Service) GetService() enums.ServiceName {
	return enums.MailService
}

func (s *Service) Send(ctx context.Context, message interface{}) error {
	payload, ok := message.(*Message)
	if !ok {
		return custerr.New("invalid_request", "failed to cast message")
	}

	msg := s.messageBuilder(*payload)

	from, err := s.getFrom(msg)
	if err != nil {
		return err
	}

	to, err := s.getRecipients(msg)
	if err != nil {
		return custerr.New("err_set_recipients", err.Error())
	}

	return s.client.Send(from, to, msg)
}

func (s *Service) messageBuilder(msg Message) *gomail.Message {
	message := gomail.NewMessage()
	message.SetAddressHeader("From", "info.rizalfadlila@gmail.com", "Rizal Fadlila")
	message.SetHeader("To", msg.To...)
	message.SetHeader("Cc", msg.Cc...)
	message.SetHeader("Subject", msg.Subject)
	message.SetBody(DefaultMime, msg.Message)

	return message
}
