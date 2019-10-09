package main

import (
	"github.com/pruknil/goapp/router/http"
	"github.com/pruknil/goapp/router/socket"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	gin := http.NewGin()
	gin.Start()

	so := socket.NewSocket()
	so.Start()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	gin.Shutdown()
	so.Shutdown()
}
