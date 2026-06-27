package service

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
	"ticketing-system/config"

	"github.com/wneessen/go-mail"
)

type VerifyEmailData struct {
	Name            string
	VerificationURL string
}

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

// TODO : make send mail use goroutine

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

func (s *SMTPService) SendVerificationEmail(ctx context.Context, toEmail string, userName string, token string) error {
	verifyURL := config.GetEnvOrPanic("EMAIL_VERIFY_URL") + "?token=" + token

	data := VerifyEmailData{
		Name:            userName,
		VerificationURL: verifyURL,
	}

	// Parse the HTML file
	tmpl, err := template.ParseFiles("public/templates/verify-email.html")
	if err != nil {
		return fmt.Errorf("failed to parse template file: %w", err)
	}

	// Render the template with the dynamic data into a buffer
	var bodyBuffer bytes.Buffer
	if err := tmpl.Execute(&bodyBuffer, data); err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}

	err = s.Send(
		ctx,
		toEmail,
		"Verify Your Email Address",
		bodyBuffer.String(),
		mail.TypeTextHTML,
	)

	if err != nil {
		return fmt.Errorf("failed to send verification email: %w", err)
	}

	return nil
}
