package main

import (
	"context"
	"fmt"
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

	err = statusLoop(c, ctx)
	if err != nil {
		log.Fatalf("could not get status: %v", err)
	}
}

func statusLoop(c pb.GameClient, ctx context.Context) error {
	log.Print("Running GetStatus()...")
	r, err := c.GetStatus(ctx, &pb.Empty{})
	if err != nil {
		return err
	}
	log.Printf("GetStatus(): %v", r)
	act(r)
	return nil
}

func act(r *pb.StatusReply) {
	switch r.Action {
	case "do auth":
		fmt.Println("Are you Player1 or Player2?")
	case "ask mark":
		fmt.Println("Choose your mark: 'X' or 'O'")
	case "ask turn":
		fmt.Println("Player 123, choose turn from 1 to 9:")
	}
}
