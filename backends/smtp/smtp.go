package smtp

import (
	"net/smtp"
)

const umail string = "p_nilsuriyakon@hotmail.com"
const upw string = "Aoom1346"
const host string = "smtp.office365.com:587"

type loginAuth struct {
	username, password string
}

func genLoginAuth(username, password string) smtp.Auth {
	return &loginAuth{username, password}
}

func (a *loginAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	return "LOGIN", []byte(a.username), nil
}

func (a *loginAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		switch string(fromServer) {
		case "Username:":
			return []byte(a.username), nil
		case "Password:":
			return []byte(a.password), nil
		}
	}
	return nil, nil
}

//SendMail 发送邮件
func SendMail(target string, body string, subject string) error {
	auth := genLoginAuth(umail, upw)

	contentType := "Content-Type: text/plain" + "; charset=UTF-8"
	msg := []byte("To: " + target +
		"\r\nFrom: " + umail +
		"\r\nSubject: " + subject +
		"\r\n" + contentType + "\r\n\r\n" +
		body)
	err := smtp.SendMail(host, auth, umail, []string{target}, msg)
	if err != nil {
		return err
	}
	return nil
}
