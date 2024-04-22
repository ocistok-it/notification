package config

type (
	Event struct {
		Listener Listener `envPrefix:"LISTENER_"`
	}
	Listener struct {
		Broker           string `env:"BROKER"`
		MaxRetryAttempts int    `env:"MAX_RETRY_ATTEMPTS"`
	}
)
