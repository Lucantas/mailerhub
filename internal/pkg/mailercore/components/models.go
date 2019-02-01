package components

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"unique_index"`
	Password string
}

// Client type with the basic information to send e-mails
type Client struct {
	gorm.Model
	SenderEmail string
	Password    string
	Mailers     []Mailer
}

// Mail type with the basic information to be send by clients
type Mail struct {
	Subject string
	Body    string
}

// Mailer is a interface representing every service that can send mails
type Mailer interface {
	SendMail()
}

// MailService is the email skeleton
type MailService struct {
	SenderID  string
	Password  string
	AddressID []string
	Subject   string
	Body      string
	SMTP      *SMTP
}

// SMTP is the skeleton with info needed to send emails
type SMTP struct {
	Host string
	Port string
}

func (s *SMTP) ServerName() string {
	return s.Host + ":" + s.Port
}

// NewMailServ return a MailServ object in order to send e-mails based on the client configuration
func NewMailServ(
	host string,
	port string,
	senderID string,
	password string,
	addressID []string,
	subject string,
	body string,
) MailService {
	return MailService{
		senderID,
		password,
		addressID,
		subject,
		body,
		&SMTP{host, port},
	}

}

// Contact is the default type to work with the service
// that sends mails based on website forms
type Contact struct {
	MailService
}

// Campaign is the default type to work with the service
// that sends mail with the purpose to be delivered at a
// large amount of destinataries
type Campaign struct {
	MailService
}

func (u User) MatchPassword(password string) bool {
	if u.Password == password {
		return true
	}
	return false
}
