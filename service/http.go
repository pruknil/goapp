package service

import (
	"encoding/json"
	"github.com/pruknil/goapp/backends/http"
	"github.com/pruknil/goapp/backends/http/service"
	"github.com/pruknil/goapp/backends/socket/hsm"
)

type AirQualityService struct {
	baseService
	service.IHttpBackend
	backendResp *http.AQIRes
}

func (s *AirQualityService) OutputMapping() error {
	s.Response = ResMsg{
		Header: ResHeader{},
		Body:   s.backendResp,
	}
	return nil
}

func (s *AirQualityService) InputMapping() error {
	return nil
}

func (s *AirQualityService) Business() error {
	x, err := s.IHttpBackend.AirQuality()
	s.backendResp = x
	if err != nil {
		return err
	}
	return nil
}

type HsmService struct {
	baseService
	hsm.IHSMService
	backendResp *hsm.StatusResponse
}

func (s *HsmService) OutputMapping() error {
	s.Response = ResMsg{
		Header: ResHeader{},
		Body:   s.backendResp,
	}
	return nil
}

func (s *HsmService) InputMapping() error {
	return nil
}

func (s *HsmService) Business() error {
	x, err := s.IHSMService.CheckStatus()
	if err != nil {
		return err
	}
	s.backendResp = x
	return nil
}

type KPeopleService struct {
	baseService
	service.IHttpBackend
	backendReq http.KPeopleReq
	backendRes *http.KPeopleRes
}

func (s *KPeopleService) OutputMapping() error {
	s.Response = ResMsg{
		Header: ResHeader{},
		Body:   s.backendRes,
	}
	return nil
}

func (s *KPeopleService) InputMapping() error {
	jsonString, _ := json.Marshal(s.Request.Body)
	json.Unmarshal(jsonString, &s.backendReq)
	return nil
}

func (s *KPeopleService) Business() error {
	res, err := s.IHttpBackend.KPeopleGetData(s.backendReq)
	s.backendRes = res
	if err != nil {
		return err
	}
	return nil
}
