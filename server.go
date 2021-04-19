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

func (s *server) GetStatus(ctx context.Context, m *pb.Empty) (*pb.StatusReply, error) {
	log.Printf("Recieved: GetStatus(), args: %v", m)
	sr := &pb.StatusReply{
		State: pb.State_IDLE,
		Actions: []pb.Actions{
			pb.Actions_START_GAME,
			pb.Actions_GET_STATUS,
		},
		Message: "You can:",
	}
	log.Print(sr)
	return sr, nil
}

func (s *server) Run(ctx context.Context, cr *pb.CommandRequest) (*pb.StatusReply, error) {
	log.Printf("Recieved: Run(), args: %v", cr)
	sr := &pb.StatusReply{
		State: pb.State_WAITING,
		For:   pb.For_AUTH,
		Actions: []pb.Actions{
			pb.Actions_GET_STATUS,
		},
		Message: "You can:",
	}
	log.Print(sr)
	return sr, nil
}

func main() {
	err := bootstrap()
	if err != nil {
		log.Fatalf("Failed to bootstrap: %v", err)
	}
	log.Print("Started ok.")
}

func bootstrap() error {
	log.Print("Starting server...")

	lis, err := net.Listen("tcp", port)
	if err != nil {
		return err
	}
	log.Print("Listening...")

	s := grpc.NewServer()
	pb.RegisterGameServer(s, &server{})

	log.Print("Serving gRPC...")
	if err := s.Serve(lis); err != nil {
		return err
	}
	return nil
}
