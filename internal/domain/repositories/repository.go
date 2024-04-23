package repositories

import (
	"context"
	"github.com/ocistok-it/notification/internal/domain/entities"
)

type (
	DefaultMailUserRepository interface {
		GetByUser(ctx context.Context, user string) (*entities.DefaultMailUser, error)
	}

	DingRobotRepository interface {
		GetByRobotID(ctx context.Context, id string) (*entities.DingRobot, error)
	}

	TemplateRepository interface {
		GetByCode(ctx context.Context, code string) (*entities.Template, error)
	}
)
