package client

import (
	"context"
	"errors"
	"fmt"
	"log"

	api "tictactoeweb/api"
)

type Ctx = context.Context

func StatusLoop(ctx Ctx) error {
	log.Print("Calling remote: GetStatus()...")
	args := &api.Empty{}
	res, err := _cli.GetStatus(ctx, args)
	if err != nil {
		return err
	}
	log.Printf("GetStatus(): Done. args: %T{%v}, res: %v", args, args, res)
	_, err = react(ctx, asdf, res)
	if err != nil {
		return err
	}
	return nil
}

func react(ctx Ctx, cli api.GameClient, sr *api.StatusReply) (*api.StatusReply, error) {
	var err error
	var r *api.StatusReply

	switch sr.State {
	case api.State_IDLE:
		fmt.Println(sr.Message)
		fmt.Println(sr.Actions)

		Play(ctx, sr)

		log.Printf("Calling remote: Run()...")
		cr := &api.CommandRequest{Action: api.Actions_START_GAME}
		log.Printf("Run() args: %v", cr)
		return cli.Run(ctx, cr)

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
