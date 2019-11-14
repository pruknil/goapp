package http

import (
	"fmt"
	breaker "github.com/sony/gobreaker"
	"io/ioutil"
	"net/http"
)

const baseURL string = "https://website-api.airvisual.com/v1/places/map?bbox=100.1972940089172,13.47902387099105,100.78038479108272,13.979292336476917&units.temperature=celsius&units.distance=kilometer&AQI=US&language=th"

type Client struct {
	Username string
	Password string
	*breaker.CircuitBreaker
	Config
}

func New(c Config) IHTTPService {
	var st breaker.Settings
	st.Name = "HTTP"
	st.Timeout = 3
	st.ReadyToTrip = func(counts breaker.Counts) bool {
		failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
		return counts.Requests >= 3 && failureRatio >= 0.6
	}
	cb := breaker.NewCircuitBreaker(st)
	return &Client{CircuitBreaker: cb, Config: c}
}

type Config struct {
}

func (s *Client) doRequest(req *http.Request) ([]byte, error) {
	req.SetBasicAuth(s.Username, s.Password)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if http.StatusOK != resp.StatusCode {
		return nil, fmt.Errorf("%s", body)
	}
	return body, nil
}
