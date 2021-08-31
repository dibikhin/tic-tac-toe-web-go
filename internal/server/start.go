package server

import (
	"log"
	"net"
	"sync"

	"google.golang.org/grpc"

	api "tictactoeweb/api"

	cfg "tictactoeweb/configs"

	. "tictactoeweb/internal"
)

func Start(wg sync.WaitGroup, teardown func()) error {
	lis, err := net.Listen("tcp", cfg.Port)
	if err != nil {
		return err
	}
	log.Print("Net: listening.")

	s := grpc.NewServer()
	log.Print("gRPC: created.")

	done := func() {
		log.Print("gRPC: stopping...")
		s.GracefulStop()
		log.Print("gRPC: stopped.")

		log.Print("App: tearing down...")
		teardown()
		// log.Print("App: teared down.")
	}
	defer done()

	wg.Add(1)
	OnExit(wg, done)

	api.RegisterGameServer(s, &server{})

	log.Print("gRPC: serving...")
	if err := s.Serve(lis); err != nil {
		return err
	}
	return nil
}
