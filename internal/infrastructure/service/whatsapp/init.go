package whatsapp

import amqp "github.com/rabbitmq/amqp091-go"

type (
	Service struct {
		queueName  string
		rabbitConn *amqp.Connection
	}

	Opts struct {
		QueueName  string
		RabbitConn *amqp.Connection
	}
)

func New(o *Opts) *Service {
	return &Service{
		queueName:  o.QueueName,
		rabbitConn: o.RabbitConn,
	}
}
