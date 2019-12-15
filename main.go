package main

import (
	"github.com/pruknil/goapp/app"
	httpbackend "github.com/pruknil/goapp/backends/http"
	service2 "github.com/pruknil/goapp/backends/http/service"
	"github.com/pruknil/goapp/backends/smtp"
	"github.com/pruknil/goapp/backends/socket/hsm"
	"github.com/pruknil/goapp/logger"
	"github.com/pruknil/goapp/router"
	"github.com/pruknil/goapp/router/http"
	"github.com/pruknil/goapp/router/socket"
	"github.com/pruknil/goapp/service"
	"go.uber.org/dig"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	container := buildContainer()
	err := invokeContainer(container)
	if err != nil {
		log.Fatal("Invoke Container error")
	}
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
func errorWrap(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
func buildContainer() *dig.Container {
	container := dig.New()
	errorWrap(container.Provide(NewConfig))
	errorWrap(container.Provide(NewLogger))

	errorWrap(container.Provide(NewHttpService))
	errorWrap(container.Provide(NewSocketService))

	errorWrap(container.Provide(NewHSMConn))
	errorWrap(container.Provide(NewHSM))
	errorWrap(container.Provide(NewHttp))
	errorWrap(container.Provide(NewHttpBackend))
	errorWrap(container.Provide(NewSmtp))

	errorWrap(container.Provide(NewRouter))
	return container
}

func NewLogger() logger.AppLog {
	al := logger.New()
	al.Error = al.NewLog("error", "debug")
	al.Perf = al.NewLog("perf", "debug")
	al.Trace = al.NewLog("trace", "debug")
	al.Rest = al.NewLog("rest", "debug")
	return al
}

//================= Start BACKEND Section =================
func NewHSMConn(cfg app.Config, log logger.AppLog) hsm.IConnection {
	return hsm.New(cfg.Hsm, log)
}

func NewHSM(b hsm.IConnection, cfg app.Config) hsm.IHSMService {
	return hsm.NewHSM(b, cfg.Hsm)
}

func NewHttp(cfg app.Config, log logger.AppLog) httpbackend.IHttpBackendService {
	return httpbackend.New(cfg.Backend.Http, log)
}

func NewHttpBackend(s httpbackend.IHttpBackendService) service2.IHttpBackend {
	return service2.New(s)
}

func NewSmtp(cfg app.Config) smtp.IMailService {
	return smtp.New(&smtp.Smtp{}, cfg.Backend.Smtp)
}

//================= End BACKEND Section =================

//Create all router here eg.. rest, socket, mq
func NewRouter(httpService service.IHttpService, socketService service.ISocketService, conf app.Config, log logger.AppLog) []router.IRouter {
	var route []router.IRouter
	route = append(route, http.NewGin(conf.Router.Http, httpService))
	route = append(route, socket.New(conf.Router.Socket, socketService, log))
	return route
}

//Http service
func NewHttpService(hsmService hsm.IHSMService, httpService service2.IHttpBackend, mailService smtp.IMailService) service.IHttpService {
	routes := make(map[string]service.IServiceTemplate)
	routes["AirQualityService"] = &service.AirQualityService{IHttpBackend: httpService}
	routes["KPeopleService"] = &service.KPeopleService{IHttpBackend: httpService}
	routes["HsmService"] = &service.HsmService{IHSMService: hsmService}
	routes["DopaService"] = &service.DopaService{IHttpBackend: httpService}
	routes["MailService"] = &service.MailService{IMailService: mailService}
	routes["QRService"] = &service.QRService{}
	return &service.HttpService{Routes: routes}
}

//Socket service
func NewSocketService() service.ISocketService {
	//routes := make(map[string]service.IServiceTemplate)
	return &service.SocketService{}
}

func NewConfig() app.Config {
	five, _ := time.ParseDuration("5s")
	ten, _ := time.ParseDuration("10s")
	return app.Config{
		Backend: app.Backend{
			Hsm: hsm.Config{
				//Host:          "172.30.154.84",
				//Port:          "2048",
				Host:          "localhost",
				Port:          "1111",
				ConnTimeout:   five,
				ReadDeadline:  five,
				WriteDeadline: five,
				PoolMin:       5,
				PoolMax:       5,
			},
			Smtp: smtp.Config{
				From:     "p_nilsuriyakon@hotmail.com",
				Password: "Aoom1346",
				Host:     "smtp.office365.com:587",
			},
		},
		Router: app.Router{
			Http: http.Config{
				Port:         "8080",
				ReadTimeout:  ten,
				WriteTimeout: ten,
				IdleTimeout:  ten,
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
