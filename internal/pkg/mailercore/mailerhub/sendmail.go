package mailerhub

import (
	"crypto/tls"
	"log"
	"mailer-service/internal/pkg/mailercore/components"
	"net/smtp"
)

type MailServ struct {
	MailService components.MailService
}

type protocol struct {
	SMTP components.SMTP
}

// NewMailServ return a MailServ object in order to send e-mails based on the client configuration
func NewMailServ(
	host string,
	port string,
	senderID string,
	password string,
	addressID []components.Address,
	subject string,
	body string,
) MailServ {
	return MailServ{
		components.MailService{
			senderID,
			password,
			addressID,
			subject,
			body,
			newSMTP(host, port),
		},
	}
}

func newSMTP(host string, port string) *components.SMTP {
	return &components.SMTP{Host: host, Port: port}
}

// Send is the responsible to send the messages over the SMTP protocol
func (m MailServ) Send() {

	auth := smtp.PlainAuth("", m.MailService.SenderID, m.MailService.Password, m.MailService.SMTP.Host)

	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         m.MailService.SMTP.Host,
	}

	conn, err := tls.Dial("tcp", m.MailService.SMTP.ServerName(), tlsconfig)
	if err != nil {
		log.Panic(err)
	}

	client, err := smtp.NewClient(conn, m.MailService.SMTP.Host)
	if err != nil {
		log.Panic(err)
	}

	if err = client.Auth(auth); err != nil {
		log.Panic(err)
	}

	if err = client.Mail(m.MailService.SenderID); err != nil {
		log.Panic(err)
	}
	for _, k := range m.MailService.AddressIDs {
		if err = client.Rcpt(k.Email); err != nil {
			log.Panic(err)
		}
	}

	w, err := client.Data()
	if err != nil {
		log.Panic(err)
	}

	_, err = w.Write([]byte(m.MailService.Body))
	if err != nil {
		log.Panic(err)
	}

	err = w.Close()
	if err != nil {
		log.Panic(err)
	}

	client.Quit()

	log.Println("Success!")
}
