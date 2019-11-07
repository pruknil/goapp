package socket

import (
	"bufio"
	"encoding/hex"
	"github.com/ianlopshire/go-fixedwidth"
	"github.com/pruknil/goapp/logger"
	"github.com/pruknil/goapp/service"
	"net"
)

type Config struct {
	Port string
}
type Socket struct {
	config  Config
	log     logger.AppLog
	service service.ISocketService
}

func New(cfg Config, sv service.ISocketService, log logger.AppLog) *Socket {
	return &Socket{
		config:  cfg,
		service: sv,
		log:     log,
	}
}

func (r *Socket) Start() {
	go func() {
		l, err := net.Listen("tcp4", ":"+r.config.Port)
		if err != nil {
			r.log.Error.Error(err)
			return
		}
		defer l.Close()
		for {
			c, err := l.Accept()
			if err != nil {
				r.log.Error.Error(err)
				return
			}
			go r.handleConnection(c)
		}
	}()
}

func (r *Socket) Shutdown() {
	r.log.Trace.Info("Socket Shutdown")
}

func (r *Socket) handleConnection(c net.Conn) {
	r.log.Trace.Infof("Serving %s\n", c.RemoteAddr().String())
	for {
		buf := make([]byte, 64)
		bufio.NewReader(c).Read(buf)

		raw := hex.EncodeToString(buf)
		r.log.Trace.Debug(raw)
		header := Msg{}
		err := fixedwidth.Unmarshal([]byte(hex.EncodeToString(buf)), &header)
		if err != nil {
			r.log.Error.Error(err)
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
		//by, _ := hex.DecodeString("01010000001d0100000000000000000000030000002a000001a6084d3039393939394500000000000000000000000000")
		//byteReturn = []byte(by)
		byteReturn = r.dispatchService(raw)
		c.Write(byteReturn)

	}
	c.Close()
	r.log.Trace.Info("Disconnect ", c.RemoteAddr().Network())
}

func (r *Socket) dispatchService(raw string) []byte {
	req := service.ReqMsg{}
	response := r.service.HSMFunc01(req)
	return response.Body.([]byte)
}

type Msg struct {
	ResponseHeader string `fixed:"1,10"`
	ResponseLen    string `fixed:"11,12"`
	Fn             string `fixed:"13,18"`
	DataLen        string `fixed:"37,38"`
}
