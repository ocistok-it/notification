package mail

import (
	"gopkg.in/gomail.v2"
)

const (
	DefaultMime string = "text/html"
)

type (
	Service struct {
		fromName string
		from     string
		client   gomail.SendCloser
	}

	Opts struct {
		FromName string
		From     string
		Client   gomail.SendCloser
	}
)

func New(o *Opts) *Service {
	return &Service{
		fromName: o.FromName,
		from:     o.From,
		client:   o.Client,
	}
}
