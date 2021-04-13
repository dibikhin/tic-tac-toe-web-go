package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "tictactoeweb/tttwebgrpc"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedGameServer
}

func (s *server) GetStatus(ctx context.Context, _ *pb.Empty) (*pb.StatusReply, error) {
	log.Print("Recieved: GetStatus()")
	return &pb.StatusReply{Status: "ok"}, nil
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
