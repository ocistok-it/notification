package consumer

import (
	"github.com/ocistok-it/notification/bootstrap/deps"
	"github.com/ocistok-it/notification/internal/infrastructure/dto"
	"github.com/ocistok-it/notification/internal/infrastructure/pkg/event"
)

type (
	consumer struct {
		consumer event.Consumer
		uc       *deps.Usecase
	}

	Opts struct {
		Consumer event.Consumer
		UC       *deps.Usecase
	}
)

func New(o *Opts) *consumer {
	return &consumer{
		consumer: o.Consumer,
		uc:       o.UC,
	}
}

func (c *consumer) Register() {
	c.consumer.Consume(dto.Event.Notify, c.notify)
}
