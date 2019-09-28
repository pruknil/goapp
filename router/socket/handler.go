package socket

import "fmt"

type Route struct{}

func New() {
	fmt.Println("Socket Create")
}

func (Route) Start() {
	fmt.Println("Socket Start")
}

func (Route) Shutdown() {
	fmt.Println("Socket Shutdown")
}
