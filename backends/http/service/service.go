package service

import (
	"bytes"
	"encoding/json"
	"github.com/pruknil/goapp/backends/http"
	"net/url"
	"strings"
)

type HttpBackendService struct {
	http.IHttpBackendService
}

func New(service http.IHttpBackendService) IHttpBackend {
	return &HttpBackendService{IHttpBackendService: service}
}

func (s *HttpBackendService) AirQuality() (*http.AQIRes, error) {
	q := url.Values{}
	q.Set("bbox", "100.1972940089172,13.47902387099105,100.78038479108272,13.979292336476917")
	q.Set("units.temperature", "celsius")
	q.Set("units.distance", "kilometer")
	q.Set("AQI", "US")
	q.Set("language", "th")
	header := make(map[string][]string)
	header["Authorization"] = []string{"Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsIng1dCI6IkJCOENlRlZxeWFHckdOdWVoSklpTDRkZmp6dyIsImtpZCI6IkJCOENlRlZxeWFHckdOdWVoSklpTDRkZmp6dyJ9.eyJhdWQiOiJodHRwczovL2thc2lrb3JuYmFua2dyb3VwLmNvbS84ZmZmN2RjNi0yMTQ3LTQ0NWEtYjcwYy1kM2I3NTFiODlmYzIiLCJpc3MiOiJodHRwczovL3N0cy53aW5kb3dzLm5ldC84ZTExZGY5Zi00NjE1LTQzNGYtYTZjNi04YTBjYjRmZmViNmMvIiwiaWF0IjoxNTczNTU2MjM0LCJuYmYiOjE1NzM1NTYyMzQsImV4cCI6MTU3MzU2MDEzNCwiYWNyIjoiMSIsImFpbyI6IjQyVmdZTmhmcWxnNE1jTEFtdWtSay91YkJUNnZqdnZxYVppdUw0bnVPbk43V3RGTGsyMEEiLCJhbXIiOlsicHdkIl0sImFwcGlkIjoiYzUxNWFiMTItZDMzYy00MWEwLWFjMzEtNGMzZjUxN2RkYjUyIiwiYXBwaWRhY3IiOiIwIiwiZmFtaWx5X25hbWUiOiJOaWxzdXJpeWFrb24iLCJnaXZlbl9uYW1lIjoiUHJ1ayIsImlwYWRkciI6IjQ5LjIzMS4yMTIuMTQ4IiwibmFtZSI6IlBydWsgTmlsc3VyaXlha29uIiwib2lkIjoiMDRkNTJjOTctODM5Yy00MTI0LWI3NzMtYmY5OGFjYzdlODNjIiwib25wcmVtX3NpZCI6IlMtMS01LTIxLTM2ODgwOTM4NjItNDA4MjI2OTg0LTM1MTAwMzEyOTktMjMzMzYwIiwic2NwIjoidXNlcl9pbXBlcnNvbmF0aW9uIiwic3ViIjoibzdXR2JYelNmVW9JNzFGMG5VeGk3cU1ESGpNOWo2MlNZdUNkci1tVnJkNCIsInRpZCI6IjhlMTFkZjlmLTQ2MTUtNDM0Zi1hNmM2LThhMGNiNGZmZWI2YyIsInVuaXF1ZV9uYW1lIjoicHJ1ay5uaUBrYnRnLnRlY2giLCJ1cG4iOiJwcnVrLm5pQGtidGcudGVjaCIsInV0aSI6IlE5TVEyUDJxWUUtSEtLVHRmcWJzQUEiLCJ2ZXIiOiIxLjAifQ.jZhdln3qkS0l_VOxbf-LQ-cfgg0y_35-xYR6w_Ril286XXy-pO6EKQJJLUF0qT6qcRWmLU_drCRsdvCryzqmKR_BpsvnbZXcnW3GhJWJMg2BD_82ynUdYG4DYrFd6nX9upmasYqPX4qpVORkkRjKyKtiUq7RK64P6CJIBxzUqiCEEYG1tuOxRjnY0PJWMfXHAvglPWDGU2XiHapSkxbQf3utcubSFE2tCYyuLBs8YVDWiNJMVpYKZ6amWXCzbXucGwxe3XIJFu-3eMDTt-25nVUiBUZX0GQCvKDV93fFfWSx7R6sxi8vzAlhLgC5FfHO5iCFxH8nTpYzG6VRKrclLw"}
	header["Content-Type"] = []string{"application/x-www-form-urlencoded; charset=UTF-8"}
	req := http.Req{
		Method: "GET",
		Url:    "https://website-api.airvisual.com/v1/places/map?" + q.Encode(),
		Header: header,
	}
	byteRs, err := s.IHttpBackendService.DoRequest(req)
	if err != nil {
		return nil, err
	}
	var data http.AQIRes
	err = json.Unmarshal(byteRs, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (s *HttpBackendService) KPeopleGetData(kreq http.KPeopleReq) (*http.KPeopleRes, error) {
	data := url.Values{}
	data.Set("EMP_ID", kreq.EMPID)
	data.Set("CO_TP_CD", kreq.COTPCD)
	header := make(map[string][]string)
	header["Authorization"] = []string{"Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsIng1dCI6IkJCOENlRlZxeWFHckdOdWVoSklpTDRkZmp6dyIsImtpZCI6IkJCOENlRlZxeWFHckdOdWVoSklpTDRkZmp6dyJ9.eyJhdWQiOiJodHRwczovL2thc2lrb3JuYmFua2dyb3VwLmNvbS84ZmZmN2RjNi0yMTQ3LTQ0NWEtYjcwYy1kM2I3NTFiODlmYzIiLCJpc3MiOiJodHRwczovL3N0cy53aW5kb3dzLm5ldC84ZTExZGY5Zi00NjE1LTQzNGYtYTZjNi04YTBjYjRmZmViNmMvIiwiaWF0IjoxNTczNTU2MjM0LCJuYmYiOjE1NzM1NTYyMzQsImV4cCI6MTU3MzU2MDEzNCwiYWNyIjoiMSIsImFpbyI6IjQyVmdZTmhmcWxnNE1jTEFtdWtSay91YkJUNnZqdnZxYVppdUw0bnVPbk43V3RGTGsyMEEiLCJhbXIiOlsicHdkIl0sImFwcGlkIjoiYzUxNWFiMTItZDMzYy00MWEwLWFjMzEtNGMzZjUxN2RkYjUyIiwiYXBwaWRhY3IiOiIwIiwiZmFtaWx5X25hbWUiOiJOaWxzdXJpeWFrb24iLCJnaXZlbl9uYW1lIjoiUHJ1ayIsImlwYWRkciI6IjQ5LjIzMS4yMTIuMTQ4IiwibmFtZSI6IlBydWsgTmlsc3VyaXlha29uIiwib2lkIjoiMDRkNTJjOTctODM5Yy00MTI0LWI3NzMtYmY5OGFjYzdlODNjIiwib25wcmVtX3NpZCI6IlMtMS01LTIxLTM2ODgwOTM4NjItNDA4MjI2OTg0LTM1MTAwMzEyOTktMjMzMzYwIiwic2NwIjoidXNlcl9pbXBlcnNvbmF0aW9uIiwic3ViIjoibzdXR2JYelNmVW9JNzFGMG5VeGk3cU1ESGpNOWo2MlNZdUNkci1tVnJkNCIsInRpZCI6IjhlMTFkZjlmLTQ2MTUtNDM0Zi1hNmM2LThhMGNiNGZmZWI2YyIsInVuaXF1ZV9uYW1lIjoicHJ1ay5uaUBrYnRnLnRlY2giLCJ1cG4iOiJwcnVrLm5pQGtidGcudGVjaCIsInV0aSI6IlE5TVEyUDJxWUUtSEtLVHRmcWJzQUEiLCJ2ZXIiOiIxLjAifQ.jZhdln3qkS0l_VOxbf-LQ-cfgg0y_35-xYR6w_Ril286XXy-pO6EKQJJLUF0qT6qcRWmLU_drCRsdvCryzqmKR_BpsvnbZXcnW3GhJWJMg2BD_82ynUdYG4DYrFd6nX9upmasYqPX4qpVORkkRjKyKtiUq7RK64P6CJIBxzUqiCEEYG1tuOxRjnY0PJWMfXHAvglPWDGU2XiHapSkxbQf3utcubSFE2tCYyuLBs8YVDWiNJMVpYKZ6amWXCzbXucGwxe3XIJFu-3eMDTt-25nVUiBUZX0GQCvKDV93fFfWSx7R6sxi8vzAlhLgC5FfHO5iCFxH8nTpYzG6VRKrclLw"}
	header["Content-Type"] = []string{"application/x-www-form-urlencoded; charset=UTF-8"}
	req := http.Req{
		Method: "POST",
		Url:    "https://iservice.kworkplace.com:8445/KProfile/api/getData",
		Body:   strings.NewReader(data.Encode()),
		Header: header,
	}

	bytes, err := s.IHttpBackendService.DoRequest(req)
	if err != nil {
		return nil, err
	}
	var out http.KPeopleRes
	err = json.Unmarshal(bytes, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (s *HttpBackendService) DopaCheckLaser(kreq http.DopaReq) (*http.DopaRes, error) {
	jsReq, _ := json.Marshal(kreq)
	header := make(map[string][]string)
	header["Content-Type"] = []string{"multipart/form-data"}
	req := http.Req{
		Method: "POST",
		Url:    "https://epit.rd.go.th/EFILING/RegAjaxController?flag=chkLaserId",
		Body:   bytes.NewBuffer(jsReq),
		Header: header,
	}
	bytes, err := s.IHttpBackendService.DoRequest(req)
	if err != nil {
		return nil, err
	}
	var data http.DopaRes
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
