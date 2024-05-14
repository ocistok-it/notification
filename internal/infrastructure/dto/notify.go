package dto

import "github.com/ocistok-it/notification/internal/infrastructure/enums"

type PushNotification struct {
	Channel  enums.Channel `json:"channel"`
	Metadata string        `json:"metadata"`
}
