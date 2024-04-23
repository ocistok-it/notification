package mailuser

import (
	"context"
	"github.com/ocistok-it/notification/internal/domain/entities"
	"github.com/ocistok-it/notification/internal/infrastructure/constants"
	"github.com/ocistok-it/notification/internal/infrastructure/custerr"
	"go.mongodb.org/mongo-driver/bson"
)

func (m *module) GetByUser(ctx context.Context, user string) (*entities.DefaultMailUser, error) {
	var result entities.DefaultMailUser

	filter := bson.M{"user": user}

	err := m.db.Collection(constants.CollDefaultMailUser).
		FindOne(ctx, filter).
		Decode(&result)

	if err != nil {
		return nil, custerr.New("query_default_mail_user", err.Error())
	}

	return &result, nil
}
