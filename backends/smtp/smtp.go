package smtp

import (
	"net/smtp"
)

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

type Config struct {
	from, password, host string
}
type MailService struct {
	ISmtp
	Config
}

func (s *MailService) BuildMail(to string, body string, subject string) error {
	auth := genLoginAuth(s.from, s.password)

	contentType := "Content-Type: text/plain" + "; charset=UTF-8"
	msg := []byte("To: " + to +
		"\r\nFrom: " + s.from +
		"\r\nSubject: " + subject +
		"\r\n" + contentType + "\r\n\r\n" +
		body)
	err := s.sendMail(s.host, auth, s.from, []string{to}, msg)
	if err != nil {
		return err
	}
	return nil
}

type Smtp struct {
}

func (s *Smtp) sendMail(addr string, a smtp.Auth, from string, to []string, msg []byte) error {
	return smtp.SendMail(addr, a, from, to, msg)
}

type MockSmtp struct {
}

func (s *MockSmtp) sendMail(addr string, a smtp.Auth, from string, to []string, msg []byte) error {
	return nil
}
