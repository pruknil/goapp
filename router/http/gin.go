package http

import (
	"context"
	"github.com/pruknil/goapp/service"
	"log"
	"net/http"
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
	g.register("AQIStatus", g.httpService)
	g.router.POST("/hsm", g.serviceLocator)
}

func (g *Gin) serviceLocator(c *gin.Context) {
	var reqMsg service.ReqMsg
	c.BindJSON(&reqMsg)
	/*
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
	*/
	c.JSON(http.StatusOK, g.httpService.HSMStatus(reqMsg))
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
