package service

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/pruknil/goapp/backends/http"
	"github.com/skip2/go-qrcode"
)

type QRService struct {
	baseService
	backendReq http.DopaReq
	backendRes *http.DopaRes
}

func (s *QRService) OutputMapping() error {
	s.Response = ResMsg{
		Header: ResHeader{},
		Body:   s.backendRes,
	}
	return nil
}

func (s *QRService) InputMapping() error {
	jsonString, _ := json.Marshal(s.Request.Body)
	json.Unmarshal(jsonString, &s.backendReq)
	return nil
}

func (s *QRService) Business() error {
	var png []byte
	q, err := qrcode.New("https://www.google.co.th", qrcode.Low)
	if err != nil {
		panic(err)
	}
	q.DisableBorder = true
	png, err = q.PNG(128)
	if err != nil {
		panic(err)
	}
	imgBase64Str := base64.StdEncoding.EncodeToString(png)
	fmt.Println(imgBase64Str)
	return nil
}
