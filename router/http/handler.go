package http

import "fmt"

type HttpRoute struct{}

func (HttpRoute) Start() {
	fmt.Println("HttpRoute Start")
}

func (HttpRoute) Shutdown() {
	fmt.Println("HttpRoute Shutdown")
}
