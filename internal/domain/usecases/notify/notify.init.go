package notify

import (
	"context"
	"github.com/ocistok-it/notification/internal/domain/repositories"
	"github.com/ocistok-it/notification/internal/domain/usecases"
	"github.com/ocistok-it/notification/internal/infrastructure/enums"
	"github.com/ocistok-it/notification/internal/infrastructure/service"
)

type (
	module struct {
		notifier  notifierMap
		initiator initiatorMap
		dingRepo  repositories.DingRobotRepository
	}

	notifierMap  map[enums.ServiceName]service.Service
	initiatorMap map[enums.ServiceName]initFn

	initFn func(ctx context.Context, metadata []byte) (interface{}, error)

	Opts struct {
		DingRepo repositories.DingRobotRepository
	}
)

func RegisterService(opts Opts, services ...service.Service) usecases.NotifyUsecase {
	m := module{
		notifier:  make(notifierMap),
		dingRepo:  opts.DingRepo,
		initiator: make(initiatorMap),
	}

	for _, v := range services {
		m.notifier[v.GetService()] = v
	}

	m.registerInitiator()

	return &m
}

func (m *module) registerInitiator() {
	m.initiator = initiatorMap{
		enums.DingTalkService: m.initDingtalk,
		enums.MailService:     m.initMail,
	}
}
