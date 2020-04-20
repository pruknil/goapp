package main

import (
	"github.com/go-stomp/stomp"
	"time"
)

type ActiveMQ struct {
	Addr string
}

//New activeMQ with addr[eg:localhost:61613] as host address.
func NewActiveMQ(addr string) *ActiveMQ {
	if addr == "" {
		addr = "localhost:61613"
	}
	return &ActiveMQ{addr}
}

// Used for health check
func (a *ActiveMQ) Check() error {
	conn, err := a.Connect()
	if err == nil {
		defer conn.Disconnect()
		return nil
	} else {
		return err
	}
}

// Connect to activeMQ
func (a *ActiveMQ) Connect() (*stomp.Conn, error) {
	var options = []func(*stomp.Conn) error{
		stomp.ConnOpt.HeartBeat(0*time.Second, 0*time.Second),
	}
	return stomp.Dial("tcp", a.Addr, options...)
}

// Send msg to destination
func (a *ActiveMQ) Send(destination string, msg string) error {
	conn, err := a.Connect()
	if err != nil {
		return err
	}
	defer conn.Disconnect()
	return conn.Send(
		destination,  // destination
		"text/plain", // content-type
		[]byte(msg))  // body
}

// Subscribe Message from destination
// func handler handle msg reveived from destination
func (a *ActiveMQ) Subscribe(done chan bool, destination string, handler func(err error, msg string)) error {
	conn, err := a.Connect()
	if err != nil {
		return err
	}
	sub, err := conn.Subscribe(destination, stomp.AckAuto)
	if err != nil {
		return err
	}
	for {
		select {
		case m := <-sub.C:
			if m != nil && m.Body != nil {
				handler(m.Err, string(m.Body))
			}
		case <-done:
			sub.Unsubscribe()
			conn.Disconnect()
		}
	}
	return nil
}

type AMQ struct {
}

func (a *AMQ) Start() {

}

func (a *AMQ) Shutdown() {

}
