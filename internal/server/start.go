package server

import (
	"log"
	"net"

	"google.golang.org/grpc"

	api "tictactoeweb/api"

	cfg "tictactoeweb/configs"

	. "tictactoeweb/internal"
)

func Serve(teardown func()) error {
	lis, srv, err := prepareServer()
	if err != nil {
		return err
	}
	defer stopServer(srv, teardown)
	// because have to stop server here
	OnExit(func() {
		stopServer(srv, teardown)
	})
	return startServer(srv, lis)
}

func prepareServer() (net.Listener, *grpc.Server, error) {
	lis, err := net.Listen("tcp", cfg.Port)
	if err != nil {
		return nil, nil, err
	}
	log.Print("Net: listening.")

	s := grpc.NewServer()
	log.Print("gRPC: created.")
	return lis, s, nil
}

func stopServer(s *grpc.Server, teardown func()) {
	log.Print("gRPC: stopping...")
	s.GracefulStop()
	log.Print("gRPC: stopped.")

	log.Print("App: tearing down...")
	teardown()
	log.Print("App: teared down.")
}

func startServer(srv *grpc.Server, lis net.Listener) error {
	api.RegisterGameServer(srv, &server{})

	log.Print("gRPC: serving...")
	if err := srv.Serve(lis); err != nil {
		return err
	}
	return nil
}
