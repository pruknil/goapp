package service

import (
	"encoding/json"
	"github.com/pruknil/goapp/backends/http"
	"github.com/pruknil/goapp/backends/smtp"
)

type MailService struct {
	baseService
	smtp.IMailService
	backendReq http.DopaReq
	backendRes *http.DopaRes
}

func (s *MailService) OutputMapping() error {
	s.Response = ResMsg{
		Header: ResHeader{},
		Body:   s.backendRes,
	}
	return nil
}

func (s *MailService) InputMapping() error {
	jsonString, _ := json.Marshal(s.Request.Body)
	json.Unmarshal(jsonString, &s.backendReq)
	return nil
}

func (s *MailService) Business() error {
	err := s.IMailService.BuildMail("pruknil@gmail.com", "FFFF", "xxxxxx")
	if err != nil {
		return err
	}
	return nil
}
