package service

import (
	"encoding/json"
	"github.com/pruknil/goapp/backends/http"
	"github.com/pruknil/goapp/backends/http/service"
)

type DopaService struct {
	baseService
	service.IHttpBackend
	backendReq http.DopaReq
	backendRes *http.DopaRes
}

func (s *DopaService) OutputMapping() error {
	s.Response = ResMsg{
		Header: ResHeader{},
		Body:   s.backendRes,
	}
	return nil
}

func (s *DopaService) InputMapping() error {
	jsonString, _ := json.Marshal(s.Request.Body)
	json.Unmarshal(jsonString, &s.backendReq)
	return nil
}

func (s *DopaService) Business() error {
	res, err := s.IHttpBackend.DopaCheckLaser(s.backendReq)
	s.backendRes = res
	if err != nil {
		return err
	}
	return nil
}
