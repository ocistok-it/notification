package deps

import (
	"github.com/ocistok-it/notification/internal/infrastructure/pkg/event"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/gomail.v2"
)

type Basic struct {
	RabitClient *amqp.Connection
	MgoClient   *mongo.Client
	Consumer    event.Consumer
	Mailer      *gomail.Dialer
}
