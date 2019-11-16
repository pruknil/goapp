package service

import (
	"encoding/hex"
)

type ISocketService interface {
	HSMFunc01(ReqMsg) ResMsg
}

type SocketService struct {
	baseService
	backendResp []byte
}

func (s *SocketService) Validate() error {
	return nil
}

func (s *SocketService) OutputMapping() error {
	s.Response = ResMsg{
		Header: ResHeader{},
		Body:   s.backendResp,
	}
	return nil
}

func (s *SocketService) InputMapping() error {
	return nil
}

func (s *SocketService) Business() error {
	by, _ := hex.DecodeString("01010000001d0100000000000000000000030000002a000001a6084d3039393939394500000000000000000000000000")
	s.backendResp = by
	return nil
}

func (s *SocketService) HSMFunc01(req ReqMsg) ResMsg {

	r, _ := s.DoService(req, s)
	//if err != nil {
	//	return "Doservice Error"
	//}
	return r
}
