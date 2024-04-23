package dto

import "github.com/ocistok-it/notification/internal/infrastructure/enums"

type PushNotification struct {
	Service  enums.ServiceName `json:"service"`
	Metadata string            `json:"metadata"`
}
