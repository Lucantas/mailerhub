package mailerhub

import (
	"mailer-service/internal/pkg/mailercore/components"
)

// Contact is the default type to work with the service
// that sends mails based on website forms
type Contact struct {
	components.Contact
}

// Campaign is the default type to work with the service
// that sends mail with the purpose to be delivered at a
// large amount of destinataries
type Campaign struct {
	components.Campaign
}

func (c Contact) SendMail(
	host string,
	port string,
	senderID string,
	password string,
	addressIDs []string,
	subject string,
	body string,
) {
	cs := NewMailServ(host, port, senderID, password, addressIDs, subject, body)

	cs.Send()
}

func newContactService(
	host string,
	port string,
	senderID string,
	password string,
	addressIDs []string,
	subject string,
	body string,
) Contact {
	return Contact{
		components.Contact{
			components.NewMailServ(
				host, port, senderID, password, addressIDs, subject, body,
			),
		},
	}
}
