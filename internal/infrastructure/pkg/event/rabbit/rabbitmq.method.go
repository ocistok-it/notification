package rabbit

import (
	"github.com/ocistok-it/notification/internal/infrastructure/pkg/event"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/rs/zerolog/log"
)

func (m *module) Use(middlewares ...event.Middleware) {
	m.middlewares = append(m.middlewares, middlewares...)
}

func (m *module) Start() {
	m.openChannel()

	for q, h := range m.handlers {
		message := m.consume(q)
		go func(message <-chan amqp.Delivery, handler event.Handler) {
			m.processMessage(message, handler)
		}(message, h)
	}
}

func (m *module) Consume(event string, handler event.Handler) {
	log.Info().Str("event", event).Msg("starting consume")
	m.handlers[event] = handler
}

func (m *module) Error() <-chan error {
	return m.errChan
}

func (m *module) Close() error {
	return m.conn.Close()
}
