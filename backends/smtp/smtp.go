package smtp

import (
	"net/smtp"
)

type MySmtp struct {
	ISmtp
	umail, upw, host string
}

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

func (s *MySmtp) BuildMail(target string, body string, subject string) error {
	auth := genLoginAuth(s.umail, s.upw)

	contentType := "Content-Type: text/plain" + "; charset=UTF-8"
	msg := []byte("To: " + target +
		"\r\nFrom: " + s.umail +
		"\r\nSubject: " + subject +
		"\r\n" + contentType + "\r\n\r\n" +
		body)
	err := s.ISmtp.sendMail(s.host, auth, s.umail, []string{target}, msg)
	if err != nil {
		return err
	}
	return nil
}

type RealSmtp struct {
}

func (s *RealSmtp) sendMail(addr string, a smtp.Auth, from string, to []string, msg []byte) error {
	return smtp.SendMail(addr, a, from, to, msg)
}

type MockSmtp struct {
}

func (s *MockSmtp) sendMail(addr string, a smtp.Auth, from string, to []string, msg []byte) error {
	return nil
}
