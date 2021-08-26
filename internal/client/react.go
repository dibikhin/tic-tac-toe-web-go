package client

import (
	"errors"
	"fmt"
	"log"

	api "tictactoeweb/api"

	. "tictactoeweb/internal"
)

func RunStatusLoop(ctx Ctx) error {
	for {
		log.Print("Calling remote: GetStatus()...")
		args := &api.Empty{}
		resp, err := Client.GetStatus(ctx, args)
		if err != nil {
			return err
		}
		log.Printf("GetStatus(): Done. args: %T{%v}, res: %v", args, args, resp)
		_, err = react(resp)
		if err != nil {
			return err
		}
	}
}

func react(sr *api.StatusReply) (*api.StatusReply, error) {
	var err error
	var r *api.StatusReply

	switch sr.State {
	case api.State_WAITING:
		switch sr.For {
		case api.For_MARK:
			SetupMarks()
			fmt.Printf("%v %v \n", sr.State, sr.For)
		case api.For_TURN:
			Play()
			fmt.Printf("%v %v \n", sr.State, sr.For)
		default:
			return nil, errors.New("unknown 'for': " + sr.For.String())
		}
	default:
		return nil, errors.New("unknown state: " + sr.State.String())
	}
	return r, err
}

// fmt.Println(sr.Message)
// fmt.Println(sr.Actions)

// Play(ctx, sr)

// log.Printf("Calling remote: Run()...")
// cr := &api.CommandRequest{Action: api.Actions_START_GAME}
// log.Printf("Run() args: %v", cr)
// return Client.Run(ctx, cr)

// case "do auth":
// 	fmt.Println("Are you Player1 or Player2?")
// case "ask mark":
// 	fmt.Println("Choose your mark: 'X' or 'O'")

// case "ask turn":
// 	fmt.Println("Player 123, choose turn from 1 to 9:")
