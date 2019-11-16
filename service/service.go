package service

import (
	"github.com/pruknil/goapp/backends/http"
	"github.com/pruknil/goapp/backends/socket/hsm"
)

//type commonFn func() error

type IHttpService interface {
	HSMStatus(ReqMsg) ResMsg
}

type HttpService struct {
	baseService
	hsm.IHSMService
	http.IHTTPService
	//backendResp *hsm.StatusResponse
}

/*
func (s *HttpService) Validate() error {
	return nil
}

func (s *HttpService) OutputMapping() error {
	s.Response = ResMsg{
		Header: ResHeader{},
		Body:   s.backendResp,
	}
	return nil
}

func (s *HttpService) InputMapping() error {
	return nil
}

func (s *HttpService) Business() error {
	var r *hsm.StatusResponse
	r, err := s.IHSMService.CheckStatus()
	s.backendResp = r

	x, err := s.IHTTPService.CheckStatus()
	fmt.Println(x)
	if err != nil {
		return err
	}
	return nil
}*/

func (s *HttpService) HSMStatus(req ReqMsg) ResMsg {
	ee := &ExampleService{
		baseService:  s.baseService,
		IHTTPService: s.IHTTPService,
		IHSMService:  s.IHSMService,
	}
	r, _ := ee.DoService(req, ee)
	//r, _ := s.DoService(req, s)
	//if err != nil {
	//	return "Doservice Error"
	//}
	return r
}
