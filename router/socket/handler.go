package socket

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"github.com/ianlopshire/go-fixedwidth"
	"net"
)

type Config struct {
	Port string
}
type Socket struct {
	config Config
}

func NewSocket(cfg Config) *Socket {
	return &Socket{config: cfg}
}

func (r *Socket) Start() {
	go func() {
		PORT := ":" + r.config.Port
		l, err := net.Listen("tcp4", PORT)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer l.Close()
		for {
			c, err := l.Accept()
			if err != nil {
				fmt.Println(err)
				return
			}
			go r.handleConnection(c)
		}
	}()
}

func (r *Socket) Shutdown() {
	fmt.Println("Socket Shutdown")
}

func (r *Socket) handleConnection(c net.Conn) {
	fmt.Printf("Serving %s\n", c.RemoteAddr().String())
	for {
		buf := make([]byte, 64)
		bufio.NewReader(c).Read(buf)

		raw := hex.EncodeToString(buf)
		fmt.Println(raw)
		header := SocketMsg{}
		err := fixedwidth.Unmarshal([]byte(hex.EncodeToString(buf)), &header)
		if err != nil {
			fmt.Println(err)
		}

		if header.ResponseHeader == "0000000000" {
			break
		}
		//fmt.Printf("%+v", header)
		var byteReturn []byte

		//if strings.HasPrefix(header.Fn, "9b") {
		//	byteReturn = genCVV(raw[16:32], raw[32:36])
		//} else if strings.HasPrefix(header.Fn, "9c") {
		//	byteReturn = verCVV(raw[16:32], raw[32:36])
		//} else if strings.HasPrefix(header.Fn, "ee0801") {
		//	byteReturn = decryptFunc(raw[76:108])
		//} else if strings.HasPrefix(header.Fn, "ee0800") {
		//	byteReturn = encryptFunc(raw[76:92])
		//}

		c.Write(byteReturn)

	}
	c.Close()
	fmt.Println("Disconnect ", c.RemoteAddr().Network())
}

type SocketMsg struct {
	ResponseHeader string `fixed:"1,10"`
	ResponseLen    string `fixed:"11,12"`
	Fn             string `fixed:"13,18"`
	DataLen        string `fixed:"37,38"`
}
