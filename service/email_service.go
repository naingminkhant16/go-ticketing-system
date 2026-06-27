package service

import (
	"bytes"
	"context"
	"html/template"
	"log"
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
) {
	msg := mail.NewMsg()
	if err := msg.From(s.from); err != nil {
		log.Printf("set from address: %w", err)
		return
	}

	if err := msg.To(to); err != nil {
		log.Printf("set to address: %w", err)
		return
	}

	msg.Subject(subject)
	msg.SetBodyString(contentType, body)

	if err := s.client.DialAndSendWithContext(ctx, msg); err != nil {
		log.Printf("Failed to send email: %s", err)
		return
	}

}

func (s *SMTPService) SendVerificationEmail(
	ctx context.Context,
	toEmail string,
	userName string,
	token string,
) {
	verifyURL := config.GetEnvOrPanic("EMAIL_VERIFY_URL") + "?token=" + token

	data := VerifyEmailData{
		Name:            userName,
		VerificationURL: verifyURL,
	}

	// Parse the HTML file
	tmpl, err := template.ParseFiles("public/templates/verify-email.html")
	if err != nil {
		log.Println("failed to parse template file")
		return
	}

	// Render the template with the dynamic data into a buffer
	var bodyBuffer bytes.Buffer
	if err := tmpl.Execute(&bodyBuffer, data); err != nil {
		log.Println("failed to execute template")
		return
	}

	s.Send(
		ctx,
		toEmail,
		"Verify Your Email Address",
		bodyBuffer.String(),
		mail.TypeTextHTML,
	)
}
