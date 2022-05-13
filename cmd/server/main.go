package main

import (
	"awesomeProject/pkg/api"
	"awesomeProject/pkg/bookinfo"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	s := grpc.NewServer()
	srv := bookinfo.GRPCServer{}
	api.RegisterBookStorageServer(s, &srv)

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
