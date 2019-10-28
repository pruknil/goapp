package hsm

import (
	"fmt"
	"gopkg.in/fatih/pool.v3"
	"net"
	"time"
)

type Config struct {
	Host          string
	Port          string
	ConnTimeout   time.Duration
	ReadDeadline  time.Duration
	WriteDeadline time.Duration
}
type HSMConnection struct {
	connPool pool.Pool
	config   Config
}

func New(cfg Config) *HSMConnection {
	return &HSMConnection{config: cfg}
}

func (h *HSMConnection) Open() error {
	dialer := net.Dialer{Timeout: h.config.ConnTimeout}
	factory := func() (net.Conn, error) {
		return dialer.Dial("tcp", fmt.Sprintf("%s:%s", h.config.Host, h.config.Port))
	}
	p, err := pool.NewChannelPool(1, 2, factory)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	h.connPool = p
	return nil
}

func (h *HSMConnection) requestConnection() (net.Conn, error) {
	if h.connPool == nil {
		h.Open()
	}

	conn, err := h.connPool.Get()
	if err != nil {
		return nil, err
	}
	return conn, nil

}

func (h *HSMConnection) Close() {
	if h.connPool != nil {
		h.connPool.Close()
	}
}
