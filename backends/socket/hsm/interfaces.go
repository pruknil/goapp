package hsm

import "net"

type IConnection interface {
	Open() error
	Close()
	RequestConnection() (net.Conn, error)
}

type IHSMService interface {
	CheckStatus() string
}
