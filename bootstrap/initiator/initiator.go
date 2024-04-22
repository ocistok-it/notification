package initiator

import (
	"github.com/caarlos0/env/v6"
	"github.com/ocistok-it/notification/bootstrap/deps"
	"github.com/ocistok-it/notification/internal/infrastructure/config"
	"github.com/rs/zerolog/log"
)

type (
	Initiator struct {
		config       *config.Main
		basic        *deps.Basic
		repositories *deps.Repository
		usecase      *deps.Usecase
	}
)

func New() *Initiator {
	return &Initiator{
		config: parseConfig(),
	}
}

func parseConfig() *config.Main {
	cfg := &config.Main{}
	if err := env.Parse(cfg); err != nil {
		log.Fatal().Err(err).Msg("failed to parse config")
	}

	return cfg
}

func (i *Initiator) SetupDeps() *deps.Deps {
	return &deps.Deps{
		Config:  i.config,
		Basic:   i.basic,
		Usecase: i.usecase,
	}
}
