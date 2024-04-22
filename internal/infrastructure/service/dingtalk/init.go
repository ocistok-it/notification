package dingtalk

import (
	"github.com/ocistok-it/notification/internal/infrastructure/config"
	"github.com/ocistok-it/notification/internal/infrastructure/pkg/apicall"
)

type (
	Service struct {
		config  config.DingTalk
		apicall *apicall.Apicall
	}

	Opts struct {
		Config  config.DingTalk
		Apicall *apicall.Apicall
	}
)

func New(o *Opts) *Service {
	return &Service{
		apicall: o.Apicall,
		config:  o.Config,
	}
}
