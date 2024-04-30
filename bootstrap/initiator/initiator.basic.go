package initiator

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/ocistok-it/notification/bootstrap/deps"
	"github.com/ocistok-it/notification/internal/infrastructure/pkg/event"
	"github.com/ocistok-it/notification/internal/infrastructure/pkg/event/rabbit"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"gopkg.in/gomail.v2"
	"time"
)

func (i *Initiator) InitBasic() *Initiator {
	i.basic = &deps.Basic{
		MgoClient: i.newMgoClient(),
		Consumer:  i.newConsumer(),
		Mailer:    i.newMailer(),
	}

	return i
}

func (i *Initiator) newMailer() gomail.SendCloser {
	cfg := i.config.Service.Mail

	dialer := gomail.NewDialer(cfg.Host, cfg.Port, cfg.Identity, cfg.Password)

	dialer.TLSConfig = &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         "mail.ocistok.co.id",
	}

	sender, err := dialer.Dial()

	if err != nil {
		log.Fatal().Err(err).Msg("error connect to server mail")
	}

	return sender
}

func (i *Initiator) newConsumer() event.Consumer {
	cfg := i.config.Event.Listener

	conn, err := amqp.Dial(cfg.Broker)
	if err != nil {
		log.Fatal().Err(err).Msg("error connect to broker")
	}

	return rabbit.NewRabbitMQ(rabbit.Opts{
		Conn: conn,
	})
}

func (i *Initiator) newMgoClient() *mongo.Client {
	cfg := i.config.Database

	connectTimeout, err := time.ParseDuration(cfg.ConnectTimeout)
	if err != nil {
		log.Fatal().Err(err).Str("value", cfg.ConnectTimeout).Msg("error parsing value connection timeout")
	}

	connectCtx, cancelConnectCtx := context.WithTimeout(context.Background(), connectTimeout)
	defer cancelConnectCtx()

	opts := []*options.ClientOptions{
		options.Client().ApplyURI(cfg.URI),
		options.Client().SetConnectTimeout(connectTimeout),
		options.Client().SetAppName("notification"),
	}

	client, err := mongo.Connect(connectCtx, opts...)
	if err != nil {
		fmt.Println(cfg.URI)
		log.Fatal().Err(err).Msg("error connect to mongo")
	}

	if err = client.Ping(connectCtx, readpref.Primary()); err != nil {
		log.Fatal().Err(err).Msg("error verify mongo client")
	}

	return client
}
