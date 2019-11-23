package http

type IHttpBackendService interface {
	DoRequest(req Req) ([]byte, error)
}
