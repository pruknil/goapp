package service

import (
	"fmt"
	"github.com/pruknil/goapp/backends/http"
	"github.com/pruknil/goapp/backends/socket/hsm"
)

//type commonFn func() error

type IHttpService interface {
	DoService(ReqMsg) ResMsg
}

type HttpService struct {
	baseService
	hsm.IHSMService
	http.IHTTPService
	Routes map[string]IServiceTemplate
}

func (s *HttpService) DoService(req ReqMsg) ResMsg {
	route, ok := s.Routes[req.Header.FuncNm]
	if !ok {
		fmt.Println("notfound")
	}

	r, _ := route.DoService(req, route)
	//r, _ := s.DoService(req, s)
	//if err != nil {
	//	return "Doservice Error"
	//}
	return r
}
