package server

import (
	"log"
	"net"

	"google.golang.org/grpc"

	api "tictactoeweb/api"

	cfg "tictactoeweb/configs"
)

func Start() error {
	log.Print("Starting server...")

	lis, err := net.Listen("tcp", cfg.Port)
	if err != nil {
		return err
	}
	log.Print("Listening...")

	s := grpc.NewServer()
	api.RegisterGameServer(s, &server{})

	log.Print("Serving gRPC...")
	if err := s.Serve(lis); err != nil {
		return err
	}
	return nil
}
