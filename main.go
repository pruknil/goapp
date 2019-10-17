package main

import (
	"github.com/pruknil/goapp/router"
	"github.com/pruknil/goapp/router/http"
	"github.com/pruknil/goapp/router/socket"
	"go.uber.org/dig"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	BuildContainer()
	gin := http.NewGin()
	so := socket.NewSocket()
	var route []router.Router
	route = append(route, gin)
	route = append(route, so)
	for _, v := range route {
		v.Start()
	}
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	for _, v := range route {
		v.Shutdown()
	}
}

func BuildContainer() *dig.Container {
	container := dig.New()
	container.Provide(NewConfig)
	return container
}

type Config struct {
	Port string
}

func NewConfig() *Config {
	return &Config{
		Port: "8000",
	}
}
