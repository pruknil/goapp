package service

import (
	"fmt"
	"github.com/pruknil/goapp/backends/http"
	"github.com/pruknil/goapp/backends/socket/hsm"
)

//type commonFn func() error

type IHSMService interface {
	HSMStatus(ReqMsg) ResMsg
}

type HSMService struct {
	baseService
	hsm.IHSMService
	http.IHTTPService
	backendResp *hsm.StatusResponse
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
	var r *hsm.StatusResponse
	r, err := s.IHSMService.CheckStatus()
	s.backendResp = r

	x, err := s.IHTTPService.CheckStatus()
	fmt.Println(x)
	if err != nil {
		return err
	}
	return nil
}

func (s *HSMService) HSMStatus(req ReqMsg) ResMsg {

	r, _ := DoService(req, s)
	//if err != nil {
	//	return "Doservice Error"
	//}
	return r
}
