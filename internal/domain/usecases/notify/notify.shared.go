package notify

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ocistok-it/notification/internal/infrastructure/custerr"
	"github.com/ocistok-it/notification/internal/infrastructure/service/dingtalk"
	"github.com/ocistok-it/notification/internal/infrastructure/service/mail"
)

func (m *module) initDingtalk(ctx context.Context, metadata string) (interface{}, error) {
	var request dingtalk.Message

	if err := json.Unmarshal([]byte(metadata), &request); err != nil {
		return nil, custerr.New("json_encode", err.Error())
	}

	dingGroup, err := m.dingRepo.GetByRobotID(ctx, request.RobotID)
	if err != nil {
		return nil, err
	}

	request.AccessToken = dingGroup.AccessToken
	request.Content = fmt.Sprintf("%s - %s", dingGroup.Secret, request.Content)

	return &request, nil
}

func (m *module) initMail(ctx context.Context, metadata string) (interface{}, error) {
	var request mail.Message

	if err := json.Unmarshal([]byte(metadata), &request); err != nil {
		return nil, custerr.New("json_encode", err.Error())
	}

	if request.DefaultUser != "" {
		defaultUser, err := m.defMailRepo.GetByUser(ctx, request.DefaultUser)
		if err != nil {
			return nil, err
		}

		request.To = append(request.To, defaultUser.To...)
		request.Cc = append(request.Cc, defaultUser.Cc...)
	}

	return &request, nil
}
