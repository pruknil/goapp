package http

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func (s *Client) CheckStatus() (*AutoGenerated, error) {
	const baseURL string = "https://website-api.airvisual.com/v1/places/map?bbox=100.1972940089172,13.47902387099105,100.78038479108272,13.979292336476917&units.temperature=celsius&units.distance=kilometer&AQI=US&language=th"
	req, err := http.NewRequest("GET", baseURL, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := s.doRequest(req)
	if err != nil {
		return nil, err
	}
	var data AutoGenerated
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (s *Client) KPeopleGetData(kreq KPeopleReq) (*KPeopleRes, error) {
	const baseURL string = "https://iservice.kworkplace.com:8445/KProfile/api/getData"
	formStr := []byte(kreq.String())
	req, err := http.NewRequest("POST", baseURL, bytes.NewBuffer(formStr))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsIng1dCI6IkJCOENlRlZxeWFHckdOdWVoSklpTDRkZmp6dyIsImtpZCI6IkJCOENlRlZxeWFHckdOdWVoSklpTDRkZmp6dyJ9.eyJhdWQiOiJodHRwczovL2thc2lrb3JuYmFua2dyb3VwLmNvbS84ZmZmN2RjNi0yMTQ3LTQ0NWEtYjcwYy1kM2I3NTFiODlmYzIiLCJpc3MiOiJodHRwczovL3N0cy53aW5kb3dzLm5ldC84ZTExZGY5Zi00NjE1LTQzNGYtYTZjNi04YTBjYjRmZmViNmMvIiwiaWF0IjoxNTczNTU2MjM0LCJuYmYiOjE1NzM1NTYyMzQsImV4cCI6MTU3MzU2MDEzNCwiYWNyIjoiMSIsImFpbyI6IjQyVmdZTmhmcWxnNE1jTEFtdWtSay91YkJUNnZqdnZxYVppdUw0bnVPbk43V3RGTGsyMEEiLCJhbXIiOlsicHdkIl0sImFwcGlkIjoiYzUxNWFiMTItZDMzYy00MWEwLWFjMzEtNGMzZjUxN2RkYjUyIiwiYXBwaWRhY3IiOiIwIiwiZmFtaWx5X25hbWUiOiJOaWxzdXJpeWFrb24iLCJnaXZlbl9uYW1lIjoiUHJ1ayIsImlwYWRkciI6IjQ5LjIzMS4yMTIuMTQ4IiwibmFtZSI6IlBydWsgTmlsc3VyaXlha29uIiwib2lkIjoiMDRkNTJjOTctODM5Yy00MTI0LWI3NzMtYmY5OGFjYzdlODNjIiwib25wcmVtX3NpZCI6IlMtMS01LTIxLTM2ODgwOTM4NjItNDA4MjI2OTg0LTM1MTAwMzEyOTktMjMzMzYwIiwic2NwIjoidXNlcl9pbXBlcnNvbmF0aW9uIiwic3ViIjoibzdXR2JYelNmVW9JNzFGMG5VeGk3cU1ESGpNOWo2MlNZdUNkci1tVnJkNCIsInRpZCI6IjhlMTFkZjlmLTQ2MTUtNDM0Zi1hNmM2LThhMGNiNGZmZWI2YyIsInVuaXF1ZV9uYW1lIjoicHJ1ay5uaUBrYnRnLnRlY2giLCJ1cG4iOiJwcnVrLm5pQGtidGcudGVjaCIsInV0aSI6IlE5TVEyUDJxWUUtSEtLVHRmcWJzQUEiLCJ2ZXIiOiIxLjAifQ.jZhdln3qkS0l_VOxbf-LQ-cfgg0y_35-xYR6w_Ril286XXy-pO6EKQJJLUF0qT6qcRWmLU_drCRsdvCryzqmKR_BpsvnbZXcnW3GhJWJMg2BD_82ynUdYG4DYrFd6nX9upmasYqPX4qpVORkkRjKyKtiUq7RK64P6CJIBxzUqiCEEYG1tuOxRjnY0PJWMfXHAvglPWDGU2XiHapSkxbQf3utcubSFE2tCYyuLBs8YVDWiNJMVpYKZ6amWXCzbXucGwxe3XIJFu-3eMDTt-25nVUiBUZX0GQCvKDV93fFfWSx7R6sxi8vzAlhLgC5FfHO5iCFxH8nTpYzG6VRKrclLw")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")

	bytes, err := s.doRequest(req)
	if err != nil {
		return nil, err
	}
	var data KPeopleRes
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
