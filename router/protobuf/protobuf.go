package protobuf

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

type ProtoBuf struct {
	srv *grpc.Server
}

func NewProtoBuf() ProtoBuf {
	return ProtoBuf{}
}

func (g *ProtoBuf) Start() {
	// create a listener on TCP port 7777
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 7777))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// create a server instance
	//s := api.Server{}
	// create a gRPC server object
	grpcServer := grpc.NewServer()
	// attach the Ping service to the server
	//api.RegisterPingServer(grpcServer, &s)
	// start the server

	go func() {
		// service connections
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %s", err)
		}
	}()
}

func (g *ProtoBuf) Shutdown() {
	log.Println("Shutdown ProtoBuf Server ...")
	g.srv.Stop()
	log.Println("Server ProtoBuf exiting")
}
