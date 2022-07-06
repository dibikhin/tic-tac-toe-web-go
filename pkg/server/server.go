package server

import (
	"fmt"
	"log"
	"net"

	"tictactoe/pkg/api"
	"tictactoe/pkg/config"

	"google.golang.org/grpc"
)

func MakeServer() *grpc.Server {
	gr := NewGameRepo()
	s := NewGameService(gr)
	gs := grpc.NewServer()

	api.RegisterGameServer(gs, s)
	return gs
}

func StartListen(cfg config.Config) net.Listener {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", cfg.Server.Port))
	if err != nil {
		log.Fatalf("server: %v", err)
	}
	return lis
}

func RunServer(srv *grpc.Server, lis net.Listener) {
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("server: serve %v", err)
	}
	log.Println("server: stopped")
}
