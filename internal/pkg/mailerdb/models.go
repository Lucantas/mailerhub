package mailerdb

// Client type with the basic information to send e-mails
type client struct {
	SenderEmail string
	Password    string
}

// Mail type with the basic information to be send by clients
type mail struct {
	Subject string
	Body    string
}

type conn struct {
	Host, Port, User, DBName, Password string
	SSL                                bool
}
