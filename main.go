package main

import (
	"github.com/pruknil/goapp/app"
	"github.com/pruknil/goapp/backends/socket/hsm"
	"github.com/pruknil/goapp/router"
	"github.com/pruknil/goapp/router/http"
	"github.com/pruknil/goapp/router/socket"
	"github.com/pruknil/goapp/service"
	"go.uber.org/dig"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	container := buildContainer()
	invokeContainer(container)
}

func invokeContainer(container *dig.Container) error {
	return container.Invoke(func(route []router.IRouter, hsm hsm.IConnection) {
		for _, v := range route {
			v.Start()
		}
		quit := make(chan os.Signal)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit
		hsm.Close()
		for _, v := range route {
			v.Shutdown()
		}
	})
}

func buildContainer() *dig.Container {
	container := dig.New()
	container.Provide(NewConfig)

	container.Provide(NewHSMConn)

	container.Provide(NewService)
	container.Provide(NewSocketService)

	container.Provide(NewHSM)
	container.Provide(NewRouter)
	return container
}

func NewHSMConn(cfg app.Config) hsm.IConnection {
	return hsm.New(cfg.Hsm)
}

func NewHSM(b hsm.IConnection, cfg app.Config) hsm.IHSMService {
	return hsm.NewHSM(b, cfg.Hsm)
}

func NewRouter(svc service.IHSMService, socketService service.ISocketService, conf app.Config) []router.IRouter {
	var route []router.IRouter
	route = append(route, http.NewGin(conf.Router.Http, svc))
	route = append(route, socket.New(conf.Router.Socket, socketService))
	return route
}

func NewService(h hsm.IHSMService) service.IHSMService {
	return &service.HSMService{IHSMService: h}
}

func NewSocketService() service.ISocketService {
	return &service.SocketService{}
}

func NewConfig() app.Config {
	fiveSec, _ := time.ParseDuration("5s")
	tenSec, _ := time.ParseDuration("10s")
	return app.Config{
		Backend: app.Backend{
			Hsm: hsm.Config{
				//Host:          "172.30.154.84",
				//Port:          "2048",
				Host:          "localhost",
				Port:          "1111",
				ConnTimeout:   fiveSec,
				ReadDeadline:  fiveSec,
				WriteDeadline: fiveSec,
				PoolMin:       5,
				PoolMax:       5,
			},
		},
		Router: app.Router{
			Http: http.Config{
				Port:         "8080",
				ReadTimeout:  tenSec,
				WriteTimeout: tenSec,
				IdleTimeout:  tenSec,
			},
			Socket: socket.Config{
				Port: "1111",
			},
		},
	}
}
