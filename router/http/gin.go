package http

import (
	"context"
	"fmt"
	"github.com/pruknil/goapp/service"
	"log"
	"net/http"
	"reflect"
	"time"

	"github.com/gin-gonic/gin"
)

type Config struct {
	Port         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
}

type Gin struct {
	srv         *http.Server
	config      Config
	httpService service.IHttpService
	router      *gin.Engine
	routes      map[string]interface{}
}

func NewGin(cfg Config, service service.IHttpService) *Gin {
	return &Gin{
		config:      cfg,
		httpService: service,
		routes:      make(map[string]interface{}),
	}
}

func (g *Gin) initializeRoutes() {
	g.register("HSMStatus", g.httpService)
	g.router.POST("/hsm", g.serviceLocator)
}

func (g *Gin) serviceLocator(c *gin.Context) {
	var reqMsg service.ReqMsg
	c.BindJSON(&reqMsg)
	route, ok := g.routes[reqMsg.Header.FuncNm]
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "notfound"})
		return
	}
	a, err := invoke(route, reqMsg.Header.FuncNm, reqMsg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, a.Interface().(service.ResMsg))
}
func invoke(any interface{}, name string, args ...interface{}) (reflect.Value, error) {
	method := reflect.ValueOf(any).MethodByName(name)
	methodType := method.Type()
	numIn := methodType.NumIn()
	if numIn > len(args) {
		return reflect.ValueOf(nil), fmt.Errorf("Method %s must have minimum %d params. Have %d", name, numIn, len(args))
	}
	if numIn != len(args) && !methodType.IsVariadic() {
		return reflect.ValueOf(nil), fmt.Errorf("Method %s must have %d params. Have %d", name, numIn, len(args))
	}
	in := make([]reflect.Value, len(args))
	for i := 0; i < len(args); i++ {
		var inType reflect.Type
		if methodType.IsVariadic() && i >= numIn-1 {
			inType = methodType.In(numIn - 1).Elem()
		} else {
			inType = methodType.In(i)
		}
		argValue := reflect.ValueOf(args[i])
		if !argValue.IsValid() {
			return reflect.ValueOf(nil), fmt.Errorf("Method %s. Param[%d] must be %s. Have %s", name, i, inType, argValue.String())
		}
		argType := argValue.Type()
		if argType.ConvertibleTo(inType) {
			in[i] = argValue.Convert(inType)
		} else {
			return reflect.ValueOf(nil), fmt.Errorf("Method %s. Param[%d] must be %s. Have %s", name, i, inType, argType)
		}
	}
	return method.Call(in)[0], nil
}

func (g *Gin) register(route string, controller interface{}) {
	g.routes[route] = controller
}

//func (g *Gin) callHsm(c *gin.Context) {
//	var u service.ReqMsg
//	c.BindJSON(&u)
//	c.JSON(http.StatusOK, g.httpService.HSMStatus(u))
//}

func (g *Gin) Start() {
	g.router = gin.Default()
	g.initializeRoutes()
	g.srv = &http.Server{
		Addr:         ":" + g.config.Port,
		Handler:      g.router,
		ReadTimeout:  g.config.ReadTimeout,
		WriteTimeout: g.config.WriteTimeout,
		IdleTimeout:  g.config.IdleTimeout,
		//MaxHeaderBytes: 1 << 20,
	}

	go func() {
		// service connections
		if err := g.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
}

func (g *Gin) Shutdown() {
	log.Println("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := g.srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	}
	log.Println("Server exiting")
}
