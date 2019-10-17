package main

import (
	"github.com/pruknil/goapp/router"
	"github.com/pruknil/goapp/router/http"
	"github.com/pruknil/goapp/router/socket"
	"github.com/pruknil/goapp/service"
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
	container.Provide(NewService)
	container.Provide(NewRouter)

	return container
}

func NewRouter(svc service.Service) []router.Router {
	var route []router.Router
	route = append(route, http.NewGin(http.Config{Port: "8080"}, svc))
	route = append(route, socket.NewSocket(socket.Config{Port: "1111"}))
	return route
}

func NewService() service.Service {
	return service.Service{}
}
