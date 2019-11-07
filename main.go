package main

import (
	"github.com/pruknil/goapp/app"
	"github.com/pruknil/goapp/backends/socket/hsm"
	"github.com/pruknil/goapp/logger"
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
	container.Provide(NewLogger)

	container.Provide(NewHSMConn)

	container.Provide(NewService)
	container.Provide(NewSocketService)

	container.Provide(NewHSM)
	container.Provide(NewRouter)
	return container
}

func NewLogger(cfg app.Config) logger.AppLog {
	al := logger.New()
	al.Error = al.NewLog("error", "info")
	al.Perf = al.NewLog("perf", "info")
	al.Trace = al.NewLog("trace", "info")
	return al
}

func NewHSMConn(cfg app.Config, log logger.AppLog) hsm.IConnection {
	return hsm.New(cfg.Hsm, log)
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

/*
func ydb() {
	// Set global node ["^hello", "world"] to "Go World"
	err := yottadb.SetValE(yottadb.NOTTP, nil, "Go World", "^hello", []string{"world"})
	if err != nil {
		panic(err)
	}

	// Retrieve the value that was set
	r, err := yottadb.ValE(yottadb.NOTTP, nil, "^hello", []string{"world"})
	if err != nil {
		panic(err)
	}
	if r != "Go World" {
		panic("Value not what was expected; did someone else set something?")
	}

	// Set a few more nodes so we can iterate through them
	err = yottadb.SetValE(yottadb.NOTTP, nil, "Go Middle Earth", "^hello", []string{"shire"})
	if err != nil {
		panic(err)
	}
	err = yottadb.SetValE(yottadb.NOTTP, nil, "Go Westeros", "^hello", []string{"Winterfell"})
	if err != nil {
		panic(err)
	}

	var cur_sub = ""
	for true {
		cur_sub, err = yottadb.SubNextE(yottadb.NOTTP, nil, "^hello", []string{cur_sub})
		if err != nil {
			error_code := yottadb.ErrorCode(err)
			if error_code == yottadb.YDB_ERR_NODEEND {
				break
			} else {
				panic(err)
			}
		}
		fmt.Printf("%s ", cur_sub)
	}

}*/
