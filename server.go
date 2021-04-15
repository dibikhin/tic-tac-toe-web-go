package main

import (
	"context"
	"log"
	"net"

	pb "tictactoeweb/tttwebgrpc"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedGameServer
}

func (s *server) GetStatus(ctx context.Context, _ *pb.Empty) (*pb.StatusReply, error) {
	log.Print("Repcieved: GetStatus()")
	sr := &pb.StatusReply{Status: "waiting turn p1", Action: "ask turn"}
	// sr := &pb.StatusReply{Status: "not started", Action: "do auth"}
	// sr := &pb.StatusReply{Status: "not started", Action: "do auth"}

	log.Print(sr)
	return sr, nil
}

func main() {
	log.Print("Starting server...")

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	log.Print("Listening...")

	s := grpc.NewServer()
	pb.RegisterGameServer(s, &server{})

	log.Print("Serving gRPC...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
