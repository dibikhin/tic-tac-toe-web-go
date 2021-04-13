package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"

	pb "tictactoeweb/tttwebgrpc"
)

const (
	address = "localhost:50051"
)

func main() {
	log.Print("Starting client...")

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	log.Print("Dialed address")
	defer conn.Close()

	c := pb.NewGameClient(conn)
	log.Print("Connected client")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	log.Print("Running GetStatus()...")
	r, err := c.GetStatus(ctx, &pb.Empty{})
	if err != nil {
		log.Fatalf("could not get status: %v", err)
	}
	log.Printf("GetStatus(): %v", r)
}
