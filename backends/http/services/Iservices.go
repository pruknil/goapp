package services

import "github.com/pruknil/goapp/backends/http"

type IHttpBackend interface {
	AirQuality() (*http.AQIRes, error)
	KPeopleGetData(http.KPeopleReq) (*http.KPeopleRes, error)
	DopaCheckLaser(http.DopaReq) (*http.DopaRes, error)
}
