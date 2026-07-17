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

func (s *SMTPService) Send(
	ctx context.Context,
	to string,
	subject string,
	body string,
	contentType mail.ContentType,
) {
	msg := mail.NewMsg()
	if err := msg.From(s.from); err != nil {
		log.Printf("set from address: %v", err)
		return
	}

	if err := msg.To(to); err != nil {
		log.Printf("set to address: %v", err)
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
	var verifyMailTemplate = `<!DOCTYPE html><html lang="es"><head>    <meta charset="UTF-8">    <title>Verify Your Email</title>    <style>        .container { font-family: Arial, sans-serif; padding: 20px; max-width: 600px; margin: 0 auto; background-color: #f9f9f9; border: 1px solid #e0e0e0; border-radius: 8px; }        .button { display: inline-block; padding: 12px 24px; margin: 20px 0; font-size: 16px; color: #ffffff !important; background-color: #007bff; text-decoration: none; border-radius: 4px; font-weight: bold; }        .footer { margin-top: 30px; font-size: 12px; color: #777777; }    </style></head><body><div class="container">    <h2>Hello {{.Name}},</h2>    <p>Thank you for registering! Please click the button below to verify your email address and activate your account.</p>    <a href="{{.VerificationURL}}" class="button" target="_blank">Verify Email Address</a>    <p>If the button doesn't work, you can also copy and paste the following link into your browser:</p>    <p><a href="{{.VerificationURL}}">Verify Link</a></p>    <hr>    <p class="footer">Link will expire in 24hr.</p></div></body></html>`

	verifyURL := config.GetEnvOrPanic("EMAIL_VERIFY_URL") + "?token=" + token

	data := VerifyEmailData{
		Name:            userName,
		VerificationURL: verifyURL,
	}

	// Parse the HTML file
	tmpl, err := template.New("verify-email").Parse(verifyMailTemplate)
	if err != nil {
		log.Printf("Failed to parse template: %v", err)
		return
	}

	// Render the template with the dynamic data into a buffer
	var bodyBuffer bytes.Buffer
	if err := tmpl.Execute(&bodyBuffer, data); err != nil {
		log.Printf("Failed to execute template: %v", err)
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
