package event

import (
	"context"
	"github.com/ocistok-it/notification/bootstrap/deps"
	"github.com/ocistok-it/notification/internal/infrastructure/handler/event/consumer"
	"github.com/ocistok-it/notification/internal/infrastructure/pkg/event"
	amqp "github.com/rabbitmq/amqp091-go"
)

type (
	module struct {
		consumer event.Consumer
		deps     *deps.Deps
		err      chan error
	}

	Handler interface {
		Run()
		Error() <-chan error
	}

	ConsumerFn func(ctx context.Context, event string, channel *amqp.Channel)
)

func New(deps *deps.Deps) Handler {
	handler := &module{
		consumer: deps.Basic.Consumer,
		deps:     deps,
		err:      nil,
	}

	opts := &consumer.Opts{
		Consumer: handler.consumer,
		UC:       handler.deps.Usecase,
	}

	consumer.New(opts).Register()

	return handler
}

func (m *module) Run() {
	m.deps.Basic.Consumer.Start()
}

func (m *module) Error() <-chan error {
	return m.consumer.Error()
}
