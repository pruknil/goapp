package hsm

import (
	"encoding/hex"
	"fmt"
	"gopkg.in/fatih/pool.v3"
	"log"
	"net"
	"time"
)

type Config struct {
	Host          string
	Port          string
	ConnTimeout   time.Duration
	ReadDeadline  time.Duration
	WriteDeadline time.Duration
	PoolMin       int
	PoolMax       int
}
type HSMConnection struct {
	connPool pool.Pool
	config   Config
}

func New(cfg Config) *HSMConnection {
	instance := &HSMConnection{config: cfg}
	instance.ping()
	return instance
}

func (h *HSMConnection) ping() {
	t := time.NewTicker(15 * time.Second)
	go func(ticker *time.Ticker) {
		for range ticker.C {
			if h.connPool != nil {
			START:
				fmt.Println("connPool >>", h.connPool.Len())
				c, err := h.connPool.Get()
				if err == nil && c != nil {
					err = c.SetDeadline(time.Now().Add(time.Second * 5))
					w, _ := hex.DecodeString("01010000000101")
					c.Write(w)
					result := make([]byte, 48)
					n, err := c.Read(result)
					if err != nil || n < 4 {
						log.Printf("read data error: %v, size: %d\n", err, n)
					}
					replyHexString := hex.EncodeToString(result)
					log.Printf("got data: %s\n", replyHexString)

					if err != nil {
						if pc, ok := c.(*pool.PoolConn); ok {
							pc.MarkUnusable()
						}
						goto START
					}
					c.Close()
				}
			}
		}
	}(t)
}

func (h *HSMConnection) Open() error {
	dialer := net.Dialer{Timeout: h.config.ConnTimeout}
	factory := func() (net.Conn, error) {
		return dialer.Dial("tcp", fmt.Sprintf("%s:%s", h.config.Host, h.config.Port))
	}
	p, err := pool.NewChannelPool(h.config.PoolMin, h.config.PoolMax, factory)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	h.connPool = p
	return nil
}

func (h *HSMConnection) requestConnection() (net.Conn, error) {
	if h.connPool == nil {
		err := h.Open()
		if err != nil {
			return nil, err
		}
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
