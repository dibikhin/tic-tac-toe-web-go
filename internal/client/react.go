package client

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	api "tictactoeweb/api"
)

func StatusLoop(c api.GameClient, ctx context.Context) error {
	log.Print("Calling remote: GetStatus()...")
	args := &api.Empty{}
	res, err := c.GetStatus(ctx, args)
	if err != nil {
		return err
	}
	log.Printf("GetStatus(): Done. args: %T{%v}, res: %v", args, args, res)
	_, err = react(c, ctx, res)
	if err != nil {
		return err
	}
	return nil
}

func react(c api.GameClient, ctx context.Context, sr *api.StatusReply) (*api.StatusReply, error) {
	var err error
	var r *api.StatusReply

	switch sr.State {
	case api.State_IDLE:
		fmt.Println(sr.Message)
		fmt.Println(sr.Actions)

		// f := func() string {
		// 	// NOTE: it's easier to create it in place on demand vs. to store
		// 	// and to initialize it somewhere. The `NewScanner` is very cheap inside actually
		// 	s := bufio.NewScanner(os.Stdin)
		// 	s.Scan()
		// 	return strings.TrimSpace(s.Text())

		// 	// TODO: have to check and propagate _scanner.Err() ?
		// }
		// f()
		Play(ctx, sr)

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
