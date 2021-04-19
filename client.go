package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"google.golang.org/grpc"

	pb "tictactoeweb/tttwebgrpc"
)

const (
	address = "localhost:50051"
)

func main() {
	log.Print("Starting client...")

	onExit(sayBye)

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	log.Print("Dialed address")
	defer conn.Close()

	cli := pb.NewGameClient(conn)
	log.Print("Connected client")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	err = statusLoop(cli, ctx)
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
	err = act(r)
	if err != nil {
		return err
	}
	return nil
}

func act(r *pb.StatusReply) error {
	switch r.Action {
	case "do auth":
		fmt.Println("Are you Player1 or Player2?")
	case "ask mark":
		fmt.Println("Choose your mark: 'X' or 'O'")
	case "ask turn":
		fmt.Println("Player 123, choose turn from 1 to 9:")
	default:
		return errors.New("unknown action: " + r.Action)
	}
	return nil
}

func onExit(done func()) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func(f func()) {
		<-c
		f()
		os.Exit(0)
	}(done)
}

func sayBye() {
	fmt.Println()
	fmt.Println("Bye!")
}
