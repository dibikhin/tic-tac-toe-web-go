package server

import (
	"log"
	"net"

	"google.golang.org/grpc"

	api "tictactoeweb/api"

	cfg "tictactoeweb/configs"
)

func Prepare() (net.Listener, *grpc.Server, error) {
	log.Print("Net: starting...")
	lis, err := net.Listen("tcp", cfg.Port)
	if err != nil {
		return nil, nil, err
	}
	log.Print("Net: listening.")

	log.Print("gRPC: creating...")
	s := grpc.NewServer()
	log.Print("gRPC: created.")
	return lis, s, nil
}

// NOTE: srv == nil will crash on start anyway
func Start(srv *grpc.Server, lis net.Listener) error {
	log.Print("gRPC: starting...")
	api.RegisterGameServer(srv, &server{})

	log.Print("gRPC: serving.")
	if err := srv.Serve(lis); err != nil {
		return err
	}
	return nil
}

func Stop(s *grpc.Server) {
	log.Print("gRPC: stopping...")
	if s == nil {
		log.Print("gRPC: server == nil, ignored")
	} else {
		s.GracefulStop()
		log.Print("gRPC: stopped.")
	}
}
