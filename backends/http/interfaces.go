package http

type IHttpBackendService interface {
	/*AirQuality() (*AQIRes, error)
	KPeopleGetData(KPeopleReq) (*KPeopleRes, error)
	DopaCheckLaser(DopaReq) (*DopaRes, error)*/
	DoRequest(req Req) ([]byte, error)
}
