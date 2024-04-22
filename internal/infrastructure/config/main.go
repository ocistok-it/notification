package config

type Main struct {
	Database Database `envPrefix:"DATABASE_"`
	Event    Event    `envPrefix:"EVENT_"`
	Service  Service  `envPrefix:"SERVICE_"`
}
