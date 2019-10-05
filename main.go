package main

import (
	"github.com/pruknil/goapp/router/http"
	"github.com/pruknil/goapp/router/protobuf"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	gin := http.NewGin()
	gin.Start()

	proto := protobuf.NewProtoBuf()
	proto.Start()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	gin.Shutdown()
	proto.Shutdown()
}
