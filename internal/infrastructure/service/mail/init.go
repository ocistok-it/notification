package mail

import (
	"gopkg.in/gomail.v2"
)

const (
	DefaultMime string = "text/html"
)

type (
	Service struct {
		senderName string
		client     gomail.SendCloser
	}

	Opts struct {
		SenderName string
		Client     gomail.SendCloser
	}
)

func New(o *Opts) *Service {
	return &Service{
		senderName: o.SenderName,
		client:     o.Client,
	}
}
