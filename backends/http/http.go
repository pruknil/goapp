package http

import (
	"fmt"
	breaker "github.com/sony/gobreaker"
	"io/ioutil"
	"net/http"
)

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
