package client

import (
	"errors"
	"fmt"
	"log"
	"tictactoeweb/internal/domain/game"

	api "tictactoeweb/api"

	. "tictactoeweb/internal"
)

func RunStatusLoop(ctx Ctx) error {
	for {
		args := &api.CommandRequest{Action: api.Actions_GET_STATUS}
		log.Printf("Calling remote: %v...", args)
		resp, err := Client().RunCommand(ctx, args)
		if resp == nil {
			return nilerr
		}
		if err != nil {
			return err
		}
		log.Printf("Remote call: Done. args: %T{%v}, res: %v", args, args, resp)

		err = react(ctx, resp)
		if err != nil {
			return err
		}
	}
}

func react(ctx Ctx, sr *api.StatusReply) error {
	switch sr.State {
	case api.Is_UNDEFINED:
		return errors.New("default state found: " + sr.State.String())
	case api.Is_WAITING:
		switch sr.For {
		case api.For_NOTHING:
			return errors.New("default 'for' found: " + sr.For.String())
		case api.For_MARK:
			fmt.Printf("%v %v \n", sr.State, sr.For)

			SetupMarks(ctx)
		case api.For_TURN:
			fmt.Printf("%v %v \n", sr.State, sr.For)

			Play(ctx)
		default:
			return errors.New("unknown 'for': " + sr.For.String())
		}
	case api.Is_GAME_OVER:
		switch sr.Outcome {
		case api.Outcome_DEFAULT:
			return errors.New("default outcome found: " + sr.Outcome.String())
		case api.Outcome_DRAW:
			fmt.Printf("%v %v \n", sr.Outcome, sr.Player)

			Domain.PrintDraw()
		case api.Outcome_WON:
			fmt.Printf("%v %v \n", sr.Outcome, sr.Player)

			p := game.NewPlayer(sr.Player.Mark, int32(sr.Player.Num))
			Domain.PrintWinner(p)
		default:
			return errors.New("unknown outcome: " + sr.Outcome.String())
		}
	default:
		return errors.New("unknown state: " + sr.State.String())
	}
	return nil
}

// fmt.Println(sr.Message)
// fmt.Println(sr.Actions)

// Play(ctx, sr)

// log.Printf("Calling remote: Run()...")
// cr := &api.CommandRequest{Action: api.Actions_START_GAME}
// log.Printf("Run() args: %v", cr)
// return Client().RunCommand(ctx, cr)

// case "do auth":
// 	fmt.Println("Are you Player1 or Player2?")
// case "ask mark":
// 	fmt.Println("Choose your mark: 'X' or 'O'")

// case "ask turn":
// 	fmt.Println("Player 123, choose turn from 1 to 9:")
