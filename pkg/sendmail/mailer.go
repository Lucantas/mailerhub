package sendmail

import (
	"crypto/tls"
	"log"
	"net/smtp"
)

// MailServ is the email skeleton
type MailServ struct {
	SenderID  string
	password  string
	AddressID []string
	Subject   string
	Body      string
	SMTP      *SMTP
}

// SMTP is the skeleton with info needed to send emails
type SMTP struct {
	host string
	port string
}

// NewMailServ return a MailServ object in order to send e-mails based on the client configuration
func NewMailServ(host string, port string, senderID string, password string, addressID []string, subject string, body string) MailServ {
	return MailServ{senderID, password, addressID, subject, body, newSMTP(host, port)}
}

func newSMTP(host string, port string) *SMTP {
	return &SMTP{host: host, port: port}
}

func (s *SMTP) serverName() string {
	return s.host + ":" + s.port
}

// SendMail is the responsible to send the messages over the SMTP protocol
func (m MailServ) SendMail() {

	auth := smtp.PlainAuth("", m.SenderID, m.password, m.SMTP.host)

	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         m.SMTP.host,
	}

	conn, err := tls.Dial("tcp", m.SMTP.serverName(), tlsconfig)
	if err != nil {
		log.Panic(err)
	}

	client, err := smtp.NewClient(conn, m.SMTP.host)
	if err != nil {
		log.Panic(err)
	}

	if err = client.Auth(auth); err != nil {
		log.Panic(err)
	}

	if err = client.Mail(m.SenderID); err != nil {
		log.Panic(err)
	}
	for _, k := range m.AddressID {
		if err = client.Rcpt(k); err != nil {
			log.Panic(err)
		}
	}

	w, err := client.Data()
	if err != nil {
		log.Panic(err)
	}

	_, err = w.Write([]byte(m.Body))
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
