package dinggroup

import (
	"context"
	"github.com/ocistok-it/notification/internal/domain/entities"
	"github.com/ocistok-it/notification/internal/infrastructure/constants"
	"github.com/ocistok-it/notification/internal/infrastructure/custerr"
	"go.mongodb.org/mongo-driver/bson"
)

func (m *module) GetByRobotID(ctx context.Context, id string) (*entities.DingRobot, error) {
	var result entities.DingRobot

	filter := bson.M{"robot_id": id}

	err := m.db.Collection(constants.CollDingRobot).
		FindOne(ctx, filter).
		Decode(&result)

	if err != nil {
		return nil, custerr.New("query_ding", err.Error())
	}

	return &result, nil
}
