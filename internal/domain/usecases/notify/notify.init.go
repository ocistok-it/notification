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

		dingRepo    repositories.DingRobotRepository
		defMailRepo repositories.DefaultMailUserRepository
	}

	notifierMap  map[enums.Channel]service.Service
	initiatorMap map[enums.Channel]initFn

	initFn func(ctx context.Context, metadata string) (interface{}, error)

	Opts struct {
		DingRepo    repositories.DingRobotRepository
		DefMailRepo repositories.DefaultMailUserRepository
	}
)

func RegisterService(opts Opts, services ...service.Service) usecases.NotifyUsecase {
	m := module{
		dingRepo:    opts.DingRepo,
		defMailRepo: opts.DefMailRepo,
		notifier:    make(notifierMap),
		initiator:   make(initiatorMap),
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
		enums.Whatsapp:        m.initWhatsapp,
	}
}
