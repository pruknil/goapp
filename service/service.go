package service

import (
	"errors"
	"fmt"
	"github.com/pruknil/goapp/backends/socket/hsm"
	"log"
	"time"
)

//type commonFn func() error

func DoService(service ServiceTemplate) (ResMsg, error) {
	defer func(s time.Time) {
		log.Printf("elpased time %0.2d ns", time.Since(s).Nanoseconds())
	}(time.Now())
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

type Service interface {
	Echo(ReqMsg) ResMsg
	//Speak() string
}

type ServiceTemplate interface {
	Validate() error
	OutputMapping() error
	InputMapping() error
	Business() error
	setRequest(ReqMsg) error
	getResponse() ResMsg
}

type DemoService struct {
	ServiceTemplate
	hsm.IHSMService
	Request  ReqMsg
	Response ResMsg
}

func (s *DemoService) getResponse() ResMsg {
	fmt.Println("getResponse")
	return s.Response
}
func (s *DemoService) setRequest(r ReqMsg) error {
	fmt.Println("setRequest")
	s.Request = r
	return nil
}

func (s *DemoService) Validate() error {
	fmt.Println("Validate")
	return nil
}

func (s *DemoService) OutputMapping() error {
	fmt.Println("OutputMapping")
	return nil
}

func (s *DemoService) InputMapping() error {
	fmt.Println("InputMapping")
	return nil
}

func (s *DemoService) Business() error {
	s.Response = ResMsg{
		Header: ResHeader{},
		Body:   s.IHSMService.CheckStatus(),
	}
	return nil
}

func (s *DemoService) Echo(ReqMsg) ResMsg {
	s.setRequest(ReqMsg{
		Header: ReqHeader{
			FuncNm:       "a",
			RqUID:        "b",
			RqDt:         "",
			RqAppID:      "123",
			UserLangPref: "h",
		},
		Body: "Helllooo",
	})

	r, _ := DoService(s)
	//if err != nil {
	//	return "Doservice Error"
	//}
	return r
}
