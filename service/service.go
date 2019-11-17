package service

import (
	"fmt"
	"github.com/pruknil/goapp/backends/http"
	"github.com/pruknil/goapp/backends/socket/hsm"
	"reflect"
)

//type commonFn func() error

type IHttpService interface {
	HSMStatus(ReqMsg) ResMsg
}

type HttpService struct {
	baseService
	hsm.IHSMService
	http.IHTTPService
	Routes map[string]interface{}
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
	route, ok := s.Routes[req.Header.FuncNm]
	if !ok {
		fmt.Println("notfound")
	}
	//ee := &ExampleService{
	//	baseService:  s.baseService,
	//	IHTTPService: s.IHTTPService,
	//	IHSMService:  s.IHSMService,
	//}
	ee := route.(ExampleService)
	r, _ := ee.DoService(req, &ee)
	//r, _ := s.DoService(req, s)
	//if err != nil {
	//	return "Doservice Error"
	//}
	return r
}

func invoke(any interface{}, name string, args ...interface{}) (reflect.Value, error) {
	method := reflect.ValueOf(any).MethodByName(name)
	methodType := method.Type()
	numIn := methodType.NumIn()
	if numIn > len(args) {
		return reflect.ValueOf(nil), fmt.Errorf("Method %s must have minimum %d params. Have %d", name, numIn, len(args))
	}
	if numIn != len(args) && !methodType.IsVariadic() {
		return reflect.ValueOf(nil), fmt.Errorf("Method %s must have %d params. Have %d", name, numIn, len(args))
	}
	in := make([]reflect.Value, len(args))
	for i := 0; i < len(args); i++ {
		var inType reflect.Type
		if methodType.IsVariadic() && i >= numIn-1 {
			inType = methodType.In(numIn - 1).Elem()
		} else {
			inType = methodType.In(i)
		}
		argValue := reflect.ValueOf(args[i])
		if !argValue.IsValid() {
			return reflect.ValueOf(nil), fmt.Errorf("Method %s. Param[%d] must be %s. Have %s", name, i, inType, argValue.String())
		}
		argType := argValue.Type()
		if argType.ConvertibleTo(inType) {
			in[i] = argValue.Convert(inType)
		} else {
			return reflect.ValueOf(nil), fmt.Errorf("Method %s. Param[%d] must be %s. Have %s", name, i, inType, argType)
		}
	}
	return method.Call(in)[0], nil
}
