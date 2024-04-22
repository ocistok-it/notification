package template

import (
	"github.com/ocistok-it/notification/internal/domain/repositories"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	module struct {
		db *mongo.Database
	}

	Opts struct {
		DB *mongo.Database
	}
)

func New(o Opts) repositories.TemplateRepository {
	return &module{
		db: o.DB,
	}
}
