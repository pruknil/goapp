package service

import (
	"errors"
	"log"
	"time"
)

type IServiceTemplate interface {
	Validate() error
	OutputMapping() error
	InputMapping() error
	Business() error
	setRequest(ReqMsg) error
	getResponse() ResMsg
}

type baseService struct {
	IServiceTemplate
	Request  ReqMsg
	Response ResMsg
}

func (s *baseService) getResponse() ResMsg {
	return s.Response
}

func (s *baseService) setRequest(r ReqMsg) error {
	s.Request = r
	return nil
}
func (s *baseService) Validate() error {
	return nil
}

func (s *baseService) DoService(req ReqMsg, service IServiceTemplate) (ResMsg, error) {
	defer func(s time.Time) {
		log.Printf("elpased time %0.2d ns", time.Since(s).Nanoseconds())
	}(time.Now())
	service.setRequest(req)
	if service.Validate() != nil {
		return ResMsg{}, errors.New("validate error")
	}

	if service.InputMapping() != nil {
		return ResMsg{}, errors.New("InputMapping Error")
	}

	if service.Business() != nil {
		return ResMsg{}, errors.New("business error")
	}

	if service.OutputMapping() != nil {
		return ResMsg{}, errors.New("OutputMapping Error")
	}

	return service.getResponse(), nil
}
