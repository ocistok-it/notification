package repositories

import (
	"context"
	"github.com/ocistok-it/notification/internal/domain/entities"
)

type (
	DingRobotRepository interface {
		GetByRobotID(ctx context.Context, id string) (*entities.DingRobot, error)
	}

	TemplateRepository interface {
		GetByCode(ctx context.Context, code string) (*entities.Template, error)
	}
)
