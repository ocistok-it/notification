package config

type (
	Service struct {
		DingTalk DingTalk `envPrefix:"DINGTALK_"`
		Mail     Mail     `envPrefix:"MAIL_"`
	}

	DingTalk struct {
		Endpoint string `env:"ENDPOINT"`
	}

	Mail struct {
		From     string `env:"EMAIL_FROM"`
		FromName string `env:"EMAIL_FROM_NAME"`
		Identity string `env:"IDENTITY"`
		Password string `env:"PASSWORD"`
		Port     int    `env:"PORT"`
		Host     string `env:"HOST"`
	}
)
