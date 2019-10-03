package main

import (
	"github.com/pruknil/goapp/router/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	gin := http.NewGin()
	gin.Start()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	gin.Shutdown()
}
