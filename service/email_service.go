package service

import (
	"context"
	"fmt"
	"ticketing-system/config"

	"github.com/wneessen/go-mail"
)

type SMTPService struct {
	client *mail.Client
	from   string
}

func NewSMTPService(cfg config.SMTPConfig) *SMTPService {
	client, err := mail.NewClient(
		cfg.Host,
		mail.WithPort(cfg.Port),
		mail.WithSMTPAuth(mail.SMTPAuthAutoDiscover),
		mail.WithUsername(cfg.Username),
		mail.WithPassword(cfg.Password),
	)

	if err != nil {
		panic("Failed to create SMTP client: " + err.Error())
	}

	return &SMTPService{
		client: client,
		from:   cfg.From,
	}
}

func (s *SMTPService) Send(
	ctx context.Context,
	to string,
	subject string,
	body string,
	contentType mail.ContentType,
) error {
	msg := mail.NewMsg()
	if err := msg.From(s.from); err != nil {
		return fmt.Errorf("set from address: %w", err)
	}

	if err := msg.To(to); err != nil {
		return fmt.Errorf("set to address: %w", err)
	}

	msg.Subject(subject)
	msg.SetBodyString(contentType, body)

	if err := s.client.DialAndSendWithContext(ctx, msg); err != nil {
		return fmt.Errorf("send email: %w", err)
	}
	return nil
}
