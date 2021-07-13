package client

import (
	"context"
	"errors"
	"fmt"
	"log"

	api "tictactoeweb/api"
)

func StatusLoop(c api.GameClient, ctx context.Context) error {
	log.Print("Calling remote: GetStatus()...")
	args := &api.Empty{}
	res, err := c.GetStatus(ctx, args)
	if err != nil {
		return err
	}
	log.Printf("GetStatus() args: %v, res: %v", args, res)
	res, err = act(c, ctx, res)
	if err != nil {
		return err
	}
	return nil
}

func act(c api.GameClient, ctx context.Context, sr *api.StatusReply) (*api.StatusReply, error) {
	var err error
	var r *api.StatusReply

	switch sr.State {
	case api.State_IDLE:
		fmt.Println(sr.Message)
		fmt.Println(sr.Actions)

		log.Printf("Calling remote: Run()...")
		cr := &api.CommandRequest{Action: api.Actions_START_GAME}
		log.Printf("Run() args: %v", cr)
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
