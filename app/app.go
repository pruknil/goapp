package app

import (
	"github.com/pruknil/goapp/backends/socket/hsm"
	"github.com/pruknil/goapp/router/http"
)

type Config struct {
	Backend
	Router
}

type Router struct {
	Http http.Config
}

type Backend struct {
	Hsm hsm.Config
}
