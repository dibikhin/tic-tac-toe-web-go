package main

import (
	"fmt"
	"log"
	"net"

	"tictactoe/api"
	"tictactoe/server"

	"google.golang.org/grpc"
)

func MakeServer() (*grpc.Server, func()) {
	gr := server.MakeGameRepo()
	gs := server.NewGameService(gr)
	s := grpc.NewServer()

	api.RegisterGameServer(s, gs)
	teardown := func() {
		log.Println("app: gracefully stopping...")
		s.GracefulStop()
	}
	return s, teardown
}

func Listen(port uint16) net.Listener {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatalf("server: %v", err)
	}
	return lis
}

func RunServer(srv *grpc.Server, lis net.Listener) {
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("server: serve: %v", err)
	}
	log.Println("server: stopped")
}
