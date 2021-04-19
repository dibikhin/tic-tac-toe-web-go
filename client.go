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
	r, err = act(c, ctx, r)
	if err != nil {
		return err
	}
	return nil
}

func act(c pb.GameClient, ctx context.Context, sr *pb.StatusReply) (*pb.StatusReply, error) {
	var err error
	var r *pb.StatusReply

	switch sr.State {
	case pb.State_IDLE:
		fmt.Println(sr.Message)
		fmt.Println(sr.Actions)

		cr := &pb.CommandRequest{Action: pb.Actions_START_GAME}
		log.Printf("Running: Run()..., args: %v", cr)
		return c.Run(ctx, cr)

	// case "do auth":
	// 	fmt.Println("Are you Player1 or Player2?")
	// case "ask mark":
	// 	fmt.Println("Choose your mark: 'X' or 'O'")
	// case "ask turn":
	// 	fmt.Println("Player 123, choose turn from 1 to 9:")
	default:
		return nil, errors.New("unknown action: " + sr.State.String())
	}
	return r, err
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
