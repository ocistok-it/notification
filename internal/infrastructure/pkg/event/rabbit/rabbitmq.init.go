package rabbit

import (
	"github.com/ocistok-it/notification/internal/infrastructure/custerr"
	"github.com/ocistok-it/notification/internal/infrastructure/pkg/event"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/rs/zerolog/log"
)

type (
	handlers map[string]event.Handler

	module struct {
		conn        *amqp.Connection
		channel     *amqp.Channel
		handlers    handlers
		middlewares []event.Middleware
		errChan     chan error
	}

	Opts struct {
		Conn             *amqp.Connection
		MaxRetryAttempts int
	}
)

func NewRabbitMQ(opts Opts) event.Consumer {
	return &module{
		conn:     parseConn(opts.Conn),
		handlers: make(handlers),
	}
}

func parseConn(conn interface{}) *amqp.Connection {
	rbConn, ok := conn.(*amqp.Connection)
	if !ok {
		log.Fatal().Err(custerr.ErrInvalidConnection()).Msg("failed parse connection")
	}
	return rbConn
}
