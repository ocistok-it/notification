package bootstrap

import "github.com/ocistok-it/notification/internal/infrastructure/handler/event"

type Consumer struct {
	*bootstrap
	handler event.Handler
}

func NewConsumer() *Consumer {
	return &Consumer{
		bootstrap: setup(),
	}
}

func (r *Consumer) RegisterHandler() *Consumer {
	r.handler = event.New(r.deps)
	return r
}

func (r *Consumer) Run() {
	go r.handler.Run()

	r.errorHandler(r.handler.Error())
}
