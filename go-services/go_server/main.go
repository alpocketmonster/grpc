package main

import (
	pkg_grpc "example/go_server/pkg/grpc"
	"example/go_server/pkg/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	s := grpc.NewServer()
	adder_server := &pkg_grpc.AdderServer{}
	proto.RegisterAdderServer(s, adder_server)

	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("server listening at %v", listen.Addr())
	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
