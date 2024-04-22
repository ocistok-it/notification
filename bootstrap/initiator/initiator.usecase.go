package initiator

import (
	"fmt"
	"github.com/ocistok-it/notification/bootstrap/deps"
	"github.com/ocistok-it/notification/internal/domain/usecases"
	"github.com/ocistok-it/notification/internal/domain/usecases/notify"
	"github.com/ocistok-it/notification/internal/infrastructure/pkg/apicall"
	"github.com/ocistok-it/notification/internal/infrastructure/service/dingtalk"
	"github.com/ocistok-it/notification/internal/infrastructure/service/mail"
)

func (i *Initiator) InitUsecase() *Initiator {
	i.usecase = &deps.Usecase{
		NotifyUC: i.newNotifyUC(),
	}

	return i
}

func (i *Initiator) newNotifyUC() usecases.NotifyUsecase {
	cfg := i.config.Service

	dt := dingtalk.New(&dingtalk.Opts{
		Config:  cfg.DingTalk,
		Apicall: apicall.New(),
	})

	email := mail.New(&mail.Opts{
		SenderName: fmt.Sprintf("%s <%s>", cfg.Mail.FromName, cfg.Mail.From),
		Client:     i.basic.Mailer,
	})

	opts := notify.Opts{
		DingRepo: i.repositories.DingGroup,
	}

	return notify.RegisterService(opts, dt, email)
}
