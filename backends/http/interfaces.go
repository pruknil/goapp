package http

import "net/http"

type IHttpBackendService interface {
	/*AirQuality() (*AQIRes, error)
	KPeopleGetData(KPeopleReq) (*KPeopleRes, error)
	DopaCheckLaser(DopaReq) (*DopaRes, error)*/
	DoRequest(req *http.Request) ([]byte, error)
}
