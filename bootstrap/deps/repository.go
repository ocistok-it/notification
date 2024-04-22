package deps

import "github.com/ocistok-it/notification/internal/domain/repositories"

type (
	Repository struct {
		Template  repositories.TemplateRepository
		DingGroup repositories.DingRobotRepository
	}
)
