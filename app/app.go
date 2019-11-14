package app

import (
	http2 "github.com/pruknil/goapp/backends/http"
	"github.com/pruknil/goapp/backends/socket/hsm"
	"github.com/pruknil/goapp/router/http"
	"github.com/pruknil/goapp/router/socket"
)

type Config struct {
	Backend
	Router
}

type Router struct {
	Http   http.Config
	Socket socket.Config
}

type Backend struct {
	Hsm  hsm.Config
	Http http2.Config
}
