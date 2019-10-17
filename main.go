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
	container := buildContainer()
	invokeContainer(container)
}
func invokeContainer(container *dig.Container) error {
	container.Invoke(func(route []router.Router) {
		for _, v := range route {
			v.Start()
		}
		quit := make(chan os.Signal)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit
		for _, v := range route {
			v.Shutdown()
		}
	})
	return nil
}

func buildContainer() *dig.Container {
	container := dig.New()
	container.Provide(NewConfig)
	container.Provide(NewRouter)

	return container
}

func NewRouter() []router.Router {
	var route []router.Router
	route = append(route, http.NewGin(http.Config{Port: "8080"}))
	route = append(route, socket.NewSocket())
	return route
}

type Config struct {
	Port string
}

func NewConfig() Config {
	return Config{
		Port: "8000",
	}
}
