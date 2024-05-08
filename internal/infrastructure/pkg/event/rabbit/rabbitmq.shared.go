package rabbit

import (
	"github.com/cenkalti/backoff/v4"
	"github.com/ocistok-it/notification/internal/infrastructure/custerr"
	"github.com/ocistok-it/notification/internal/infrastructure/pkg/event"
	"github.com/ocistok-it/notification/internal/infrastructure/util"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/rs/zerolog/log"
	"time"
)

func (m *module) processMessage(delivery <-chan amqp.Delivery, handler event.Handler) {
	for deliver := range delivery {
		go func(d amqp.Delivery) {
			if err := m.doProcessing(d.Body, handler); err != nil {
				if err := d.Nack(false, false); err != nil {
					m.logErrConsume(err, d, "failed to sent message to dlq")
				}
				m.logErrConsume(err, d, "failed to process message")
				return
			}
			if err := d.Ack(false); err != nil {
				m.logErrConsume(err, d, "failed to ack message")
			}
		}(deliver)
	}
}

func (m *module) doProcessing(body []byte, handler event.Handler) error {
	opts := backoff.NewExponentialBackOff(
		backoff.WithMaxElapsedTime(2*time.Second),
		backoff.WithMaxInterval(10*time.Second),
	)

	return backoff.Retry(func() error {
		return handler(util.ContextBackground(), body)
	}, opts)
}

func (m *module) openChannel() {
	channel, err := m.conn.Channel()
	if err != nil {
		m.errChan <- custerr.New("rabbitmq_open_channel", "failed open channel").WithStacktrace(err)
		return
	}
	m.channel = channel
}

func (m *module) consume(queue string) <-chan amqp.Delivery {
	msgs, err := m.channel.ConsumeWithContext(
		util.ContextBackground(),
		queue, // queue
		"",    // consumer-tag
		false, // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)
	if err != nil {
		m.errChan <- custerr.New("rabbitmq_consume", err.Error())
		return nil
	}

	return msgs
}

func (m *module) logErrConsume(err error, d amqp.Delivery, message string) {
	if v, ok := (err).(*custerr.CustErr); ok {
		log.Err(err).
			Str("code", v.Code).
			Str("body", string(d.Body)).
			Msg(message)

		return
	}
	log.Err(err).
		Str("body", string(d.Body)).
		Msg(message)
}
