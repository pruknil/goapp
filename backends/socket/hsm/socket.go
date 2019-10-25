package hsm

import (
	"fmt"
	"gopkg.in/fatih/pool.v3"
	"net"
	"time"
)

type HSMConnection struct {
	connPool pool.Pool
}

func New() *HSMConnection {
	return &HSMConnection{}
}

func (h *HSMConnection) Open() error {
	t, _ := time.ParseDuration("5s")
	dialer := net.Dialer{Timeout: t}
	factory := func() (net.Conn, error) { return dialer.Dial("tcp", "172.30.154.84:2048") }
	p, err := pool.NewChannelPool(1, 2, factory)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	h.connPool = p
	return nil
}

func (h *HSMConnection) RequestConnection() (net.Conn, error) {
	conn, err := h.connPool.Get()

	if err != nil {
		return nil, err
	}
	return conn, nil

}

func (h *HSMConnection) Close() {
	h.connPool.Close()
}
