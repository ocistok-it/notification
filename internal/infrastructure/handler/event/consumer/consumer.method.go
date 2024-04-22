package consumer

import (
	"context"
	"encoding/json"
	"github.com/ocistok-it/notification/internal/infrastructure/dto"
	"github.com/rs/zerolog/log"
)

func (c *consumer) notify(ctx context.Context, body []byte) error {
	var request dto.PushNotification
	if err := json.Unmarshal(body, &request); err != nil {
		log.Error().Err(err).Msg("error notify")
		return err
	}

	err := c.uc.NotifyUC.Send(ctx, &request)

	return err
}
