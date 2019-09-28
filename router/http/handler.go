package http

import "fmt"

type Route struct {
}

func New() Route {
	return Route{}
}

func (Route) Start() {
	fmt.Println("HttpRoute Start")
}

func (Route) Shutdown() {
	fmt.Println("HttpRoute Shutdown")
}
