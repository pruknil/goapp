package smtp

import "net/smtp"

type ISmtp interface {
	sendMail(addr string, a smtp.Auth, from string, to []string, msg []byte) error
}

type IMailService interface {
	BuildMail(to string, body string, subject string) error
}
