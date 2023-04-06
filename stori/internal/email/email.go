package email

import (
	"stori/config"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type EmailModule struct{}

func (m *EmailModule) ProvideEmailModule() Email {
	return &email{
		Sengrid: sendgrid.NewSendClient(config.Config.SG_KEY),
	}
}

type email struct {
	Sengrid *sendgrid.Client
}

type Email interface {
	Send(to, from, subject, body string) error
}

func (e *email) Send(to, from, subject, body string) error {
	sgto := mail.NewEmail("", to)
	sgfrom := mail.NewEmail("Stori", from)

	message := mail.NewSingleEmail(sgfrom, subject, sgto, "", body)
	_, err := e.Sengrid.Send(message)
	if err != nil {
		return err
	}

	return nil
}
