package service

import (
	"errors"
	"github.com/pruknil/goapp/backends/socket/hsm"
	"log"
	"time"
)

//type commonFn func() error

func DoService(req ReqMsg, service IServiceTemplate) (ResMsg, error) {
	defer func(s time.Time) {
		log.Printf("elpased time %0.2d ns", time.Since(s).Nanoseconds())
	}(time.Now())
	service.setRequest(req)
	if service.Validate() != nil {
		return ResMsg{}, errors.New("Validate Error")
	}

	if service.InputMapping() != nil {
		return ResMsg{}, errors.New("InputMapping Error")
	}

	if service.Business() != nil {
		return ResMsg{}, errors.New("Business Error")
	}

	if service.OutputMapping() != nil {
		return ResMsg{}, errors.New("OutputMapping Error")
	}

	return service.getResponse(), nil
}

type IService interface {
	HSMStatus(ReqMsg) ResMsg
}

type IServiceTemplate interface {
	Validate() error
	OutputMapping() error
	InputMapping() error
	Business() error
	setRequest(ReqMsg) error
	getResponse() ResMsg
}

type HSMService struct {
	IServiceTemplate
	Request  ReqMsg
	Response ResMsg
	hsm.IHSMService
	backendResp *hsm.HSMStatusResponse
}

func (s *HSMService) getResponse() ResMsg {
	return s.Response
}

func (s *HSMService) setRequest(r ReqMsg) error {
	s.Request = r
	return nil
}

func (s *HSMService) Validate() error {
	return nil
}

func (s *HSMService) OutputMapping() error {
	s.Response = ResMsg{
		Header: ResHeader{},
		Body:   s.backendResp,
	}
	return nil
}

func (s *HSMService) InputMapping() error {
	return nil
}

func (s *HSMService) Business() error {
	s.backendResp = s.IHSMService.CheckStatus()
	return nil
}

func (s *HSMService) HSMStatus(req ReqMsg) ResMsg {

	r, _ := DoService(req, s)
	//if err != nil {
	//	return "Doservice Error"
	//}
	return r
}
