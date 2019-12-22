package service

import (
	"encoding/base64"
	"encoding/json"
	"github.com/skip2/go-qrcode"
)

type QRService struct {
	baseService
	frontendReq QRServiceReq
	frontendRes QRServiceRes
	//backendReq http.DopaReq
	//backendRes *http.DopaRes
}

type QRServiceReq struct {
	Input string `json:"input"`
}
type QRServiceRes struct {
	Base64Str string `json:"base64Str"`
}

func (s *QRService) OutputMapping() error {
	s.Response = ResMsg{
		Header: ResHeader{},
		Body:   s.frontendRes,
	}
	return nil
}

func (s *QRService) InputMapping() error {
	jsonString, _ := json.Marshal(s.Request.Body)
	json.Unmarshal(jsonString, &s.frontendReq)
	return nil
}

func (s *QRService) Business() error {
	var png []byte
	q, err := qrcode.New(s.frontendReq.Input, qrcode.Low)
	if err != nil {
		panic(err)
	}
	q.DisableBorder = true
	png, err = q.PNG(128)
	if err != nil {
		panic(err)
	}
	imgBase64Str := base64.StdEncoding.EncodeToString(png)
	//fmt.Println(imgBase64Str)
	s.frontendRes.Base64Str = imgBase64Str
	return nil
}
