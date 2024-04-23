package initiator

import (
	"github.com/ocistok-it/notification/bootstrap/deps"
	"github.com/ocistok-it/notification/internal/domain/repositories"
	"github.com/ocistok-it/notification/internal/domain/repositories/ding_robot"
	mailuser "github.com/ocistok-it/notification/internal/domain/repositories/mail_user"
	"github.com/ocistok-it/notification/internal/domain/repositories/template"
)

func (i *Initiator) InitRepository() *Initiator {
	i.repositories = &deps.Repository{
		Template:    i.newTemplateRepo(),
		DingRobot:   i.newDingRobotRepo(),
		DefMailUser: i.newDefMailUser(),
	}

	return i
}

func (i *Initiator) newTemplateRepo() repositories.TemplateRepository {
	return template.New(template.Opts{DB: i.basic.MgoClient.Database(i.config.Database.Name)})
}

func (i *Initiator) newDingRobotRepo() repositories.DingRobotRepository {
	return dingrobot.New(dingrobot.Opts{DB: i.basic.MgoClient.Database(i.config.Database.Name)})
}

func (i *Initiator) newDefMailUser() repositories.DefaultMailUserRepository {
	return mailuser.New(mailuser.Opts{DB: i.basic.MgoClient.Database(i.config.Database.Name)})
}
