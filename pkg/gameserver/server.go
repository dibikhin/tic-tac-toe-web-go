package gameserver

import (
	"fmt"
	"log"
	"net"

	"tictactoe/app"
	"tictactoe/pkg/api"

	"google.golang.org/grpc"
)

func Make() *grpc.Server {
	gr := MakeGameRepo()
	s := NewService(gr)
	gs := grpc.NewServer()

	api.RegisterGameServer(gs, s)
	return gs
}

func Listen(cfg app.Config) net.Listener {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", cfg.GameServer.Port))
	if err != nil {
		log.Fatalf("server: %v", err)
	}
	return lis
}

func Run(srv *grpc.Server, lis net.Listener) {
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("server: serve: %v", err)
	}
	log.Println("server: stopped")
}
