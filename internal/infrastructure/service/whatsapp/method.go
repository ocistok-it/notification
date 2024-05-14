package whatsapp

import (
	"context"
	"github.com/ocistok-it/notification/internal/infrastructure/enums"
	amqp "github.com/rabbitmq/amqp091-go"
)

func (s *Service) GetService() enums.Channel {
	return enums.Whatsapp
}

func (s *Service) Send(ctx context.Context, message interface{}) error {
	ch, err := s.rabbitConn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	err = ch.PublishWithContext(ctx, "", s.queueName, false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        []byte(message.(string)),
	})

	return err
}
