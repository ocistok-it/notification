package config

type Database struct {
	URI                  string `env:"URI"`
	Name                 string `env:"NAME"`
	ConnectTimeout       string `env:"CONNECTION_TIMEOUT"`
	PingTimeout          string `env:"PING_TIMEOUT"`
	RetryConnectInterval string `env:"RETRY_CONNECT_INTERVAL"`
}
