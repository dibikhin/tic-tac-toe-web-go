package server

import (
	"log"
	"net"

	"tictactoe/pkg/api"
	"tictactoe/pkg/config"

	"google.golang.org/grpc"
)

func MakeServer() *grpc.Server {
	var games []Game // In-mem storage
	gr := NewGameRepo(games)
	s := NewGameService(gr)
	gs := grpc.NewServer()

	api.RegisterGameServer(gs, s)
	return gs
}

func StartListen(cfg config.Config) net.Listener {
	lis, err := net.Listen("tcp", cfg.Server.Port)
	if err != nil {
		log.Fatalf("net: %v", err)
	}
	return lis
}

func RunServer(srv *grpc.Server, lis net.Listener) {
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("srv: %v", err)
	}
}
