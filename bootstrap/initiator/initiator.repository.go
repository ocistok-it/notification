package initiator

import (
	"github.com/ocistok-it/notification/bootstrap/deps"
	"github.com/ocistok-it/notification/internal/domain/repositories"
	dinggroup "github.com/ocistok-it/notification/internal/domain/repositories/ding_robot"
	"github.com/ocistok-it/notification/internal/domain/repositories/template"
)

func (i *Initiator) InitRepository() *Initiator {
	i.repositories = &deps.Repository{
		Template:  i.newTemplateRepo(),
		DingGroup: i.newDingGroupRepo(),
	}

	return i
}

func (i *Initiator) newTemplateRepo() repositories.TemplateRepository {
	return template.New(template.Opts{DB: i.basic.MgoClient.Database(i.config.Database.Name)})
}

func (i *Initiator) newDingGroupRepo() repositories.DingRobotRepository {
	return dinggroup.New(dinggroup.Opts{DB: i.basic.MgoClient.Database(i.config.Database.Name)})
}
