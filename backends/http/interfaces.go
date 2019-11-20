package http

type IHTTPService interface {
	AirQuality() (*AQIRes, error)
	KPeopleGetData(KPeopleReq) (*KPeopleRes, error)
	DopaCheckLaser(DopaReq) (*DopaRes, error)
}
