package mailerhub

import (
	"mailer-service/internal/pkg/mailercore/components"
)

// Client type with the basic information to send e-mails
type Client struct {
	components.Client
}

// Mail type with the basic information to be send by clients
type Mail struct {
	components.Mail
}
