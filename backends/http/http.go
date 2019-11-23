package http

import (
	"fmt"
	breaker "github.com/sony/gobreaker"
	"io/ioutil"
	"net/http"
)

type Client struct {
	*breaker.CircuitBreaker
	Config
}

func New(c Config) IHttpBackendService {
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

func (s *Client) DoRequest(req *http.Request) ([]byte, error) {
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
	fmt.Println("response Body:", string(body))
	if http.StatusOK != resp.StatusCode {
		return nil, fmt.Errorf("%s", body)
	}
	return body, nil
}
