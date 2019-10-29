package hsm

import "net"

type IConnection interface {
	Open() error
	Close()
	requestConnection() (net.Conn, error)
}

type IHSMService interface {
	CheckStatus() *HSMStatusResponse
}
