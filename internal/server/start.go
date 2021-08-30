package server

import (
	"log"
	"net"

	"google.golang.org/grpc"

	api "tictactoeweb/api"

	cfg "tictactoeweb/configs"
)

func Start() (func(), error) {
	log.Print("Starting server...")

	lis, err := net.Listen("tcp", cfg.Port)
	if err != nil {
		return nil, err
	}
	log.Print("Listening...")

	s := grpc.NewServer()
	teardown := func() {
		s.GracefulStop()
	}
	api.RegisterGameServer(s, &server{})
	log.Print("Serving gRPC...")

	if err := s.Serve(lis); err != nil {
		return nil, err
	}
	return teardown, nil
}
