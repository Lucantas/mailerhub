package mailerdb

import "mailer-service/internal/pkg/mailercore/components"

// Client type with the basic information to send e-mails
type client struct {
	C components.Client
}

// Mail type with the basic information to be send by clients
type mail struct {
	M components.Mail
}

type mailer struct {
	M components.Mailer
}

type conn struct {
	Host, Port, User, DBName, Password string
	SSL                                bool
}
