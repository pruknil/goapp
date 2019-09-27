package socket

import "fmt"

type SocketRoute struct{}

func (SocketRoute) Start() {
	fmt.Println("Socket Start")
}

func (SocketRoute) Shutdown() {
	fmt.Println("Socket Shutdown")
}
