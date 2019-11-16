package service

import (
	"fmt"
	"github.com/pruknil/goapp/backends/http"
	"github.com/pruknil/goapp/backends/socket/hsm"
)

type ExampleService struct {
	baseService
	http.IHTTPService
	hsm.IHSMService
	backendResp *hsm.StatusResponse
}

func (s *ExampleService) Validate() error {
	return nil
}

func (s *ExampleService) OutputMapping() error {
	s.Response = ResMsg{
		Header: ResHeader{},
		Body:   s.backendResp,
	}
	return nil
}

func (s *ExampleService) InputMapping() error {
	return nil
}

func (s *ExampleService) Business() error {
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
