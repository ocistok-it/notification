package mail

import (
	"errors"
	"fmt"
	"gopkg.in/gomail.v2"
	"net/mail"
)

func (s *Service) getFrom(msg *gomail.Message) (string, error) {
	from := msg.GetHeader("Sender")
	if len(from) == 0 {
		from = msg.GetHeader("From")
		if len(from) == 0 {
			return "", errors.New(`gomail: invalid message, "From" field is absent`)
		}
	}

	return s.parseAddress(from[0])
}

func (s *Service) getRecipients(msg *gomail.Message) ([]string, error) {
	n := 0
	for _, field := range []string{"To", "Cc", "Bcc"} {
		addresses := msg.GetHeader(field)
		n += len(addresses)
	}
	list := make([]string, 0, n)

	for _, field := range []string{"To", "Cc", "Bcc"} {
		addresses := msg.GetHeader(field)
		for _, a := range addresses {
			addr, err := s.parseAddress(a)
			if err != nil {
				return nil, err
			}
			list = s.addAddress(list, addr)
		}
	}

	return list, nil
}

func (s *Service) addAddress(list []string, addr string) []string {
	for _, a := range list {
		if addr == a {
			return list
		}
	}

	return append(list, addr)
}

func (s *Service) parseAddress(field string) (string, error) {
	addr, err := mail.ParseAddress(field)
	if err != nil {
		return "", fmt.Errorf("gomail: invalid address %q: %v", field, err)
	}
	return addr.Address, nil
}
