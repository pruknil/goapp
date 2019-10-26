package app

import "github.com/pruknil/goapp/backends/socket/hsm"

type Config struct {
	Backend
}
type Backend struct {
	Hsm hsm.Config
}
