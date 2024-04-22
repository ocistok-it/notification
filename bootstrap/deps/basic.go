package deps

import (
	"github.com/ocistok-it/notification/internal/infrastructure/pkg/event"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/gomail.v2"
)

type Basic struct {
	MgoClient *mongo.Client
	Consumer  event.Consumer
	Mailer    gomail.SendCloser
}
