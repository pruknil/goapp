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
	srv     *http.Server
	config  Config
	service service.Service
}

func NewGin(cfg Config, sv service.Service) *Gin {
	return &Gin{config: cfg, service: sv}
}

func (g *Gin) Start() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		g.service.Echo(service.ReqMsg{
			Header: service.ReqHeader{
				FuncNm:       "Echo",
				RqUID:        "",
				RqDt:         "",
				RqAppID:      "123",
				UserLangPref: "th",
			},
			Body: nil,
		})
		c.String(http.StatusOK, "Welcome Gin Server ")
	})

	g.srv = &http.Server{
		Addr:           ":" + g.config.Port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		IdleTimeout:    10 * time.Second,
		MaxHeaderBytes: 1 << 20,
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
