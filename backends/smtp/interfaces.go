package smtp

import "net/smtp"

type ISmtp interface {
	SendMail(addr string, a smtp.Auth, from string, to []string, msg []byte) error
}
