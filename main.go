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
	container.Invoke(func(route []router.Router, hsm hsm.IConnection) {
		/*
			if err := hsm.Open(); err != nil {
				panic(`
					// ------------------------------------------------------------------------------
					//! Fail to connect HSM: ` + err.Error() + `
					// ------------------------------------------------------------------------------
					`)
			}
		*/
		for _, v := range route {
			v.Start()
		}
		//go backgroundTask()
		quit := make(chan os.Signal)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit
		hsm.Close()
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
	container.Provide(NewHSMConn)
	container.Provide(NewHSM)
	container.Provide(NewConfig)
	return container
}

func NewHSMConn(cfg app.Config) hsm.IConnection {
	return hsm.New(cfg.Hsm)
}

func NewHSM(b hsm.IConnection) hsm.IHSMService {
	return hsm.NewHSM(b)
}

func NewRouter(svc service.Service) []router.Router {
	var route []router.Router
	route = append(route, http.NewGin(http.Config{Port: "8080"}, svc))
	route = append(route, socket.NewSocket(socket.Config{Port: "1111"}, svc))
	return route
}

func NewService(h hsm.IHSMService) service.Service {
	return &service.DemoService{IHSMService: h}
}
func NewConfig() app.Config {
	connTmout, _ := time.ParseDuration("5s")
	return app.Config{
		Backend: app.Backend{
			Hsm: hsm.Config{
				Host:        "localhost",
				Port:        "1111",
				ConnTimeout: connTmout,
			},
		},
	}
}

//
//func backgroundTask() {
//	ticker := time.NewTicker(1 * time.Second)
//	for _ = range ticker.C {
//		fmt.Println("Tock")
//	}
//}
