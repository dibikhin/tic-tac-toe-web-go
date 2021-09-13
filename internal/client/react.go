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
		args := &api.QueryRequest{Query: api.Querys_GET_STATUS}
		log.Printf("Client: calling remote. args: %v...", args)
		resp, err := GameClient().RunQuery(ctx, args)
		// TODO:
		// if resp == nil {
		// 	return nilerr
		// }
		if err != nil {
			return err
		}
		log.Printf("Client: remote call done. args: %T{%v}, res: %v", args, args, resp)

		log.Print("Game: reacting on state...")
		err = react(ctx, resp)
		if err != nil {
			return err
		}
	}
}

func react(ctx Ctx, sr *api.StatusReply) error {
	fmt.Printf("%v %v \n", sr.State, sr.For)

	switch sr.State {
	case api.Is_UNDEFINED:
		return errors.New("default state found: " + sr.State.String())
	case api.Is_WAITING:
		return play(ctx, sr)
	case api.Is_GAME_OVER:
		fmt.Printf("%v %v \n", sr.Outcome, sr.Player)
		return gameOver(sr)
	default:
		return errors.New("unknown state: " + sr.State.String())
	}
}

func play(ctx Ctx, sr *api.StatusReply) error {
	switch sr.For {
	case api.For_NOTHING:
		return errors.New("default 'for' found: " + sr.For.String())
	case api.For_MARK:
		_, err := SetupMarks(ctx)
		return err
	case api.For_TURN:
		return Play(ctx, sr) // TODO: parse players
	default:
		return errors.New("unknown 'for': " + sr.For.String())
	}
}

func gameOver(sr *api.StatusReply) error {
	switch sr.Outcome {
	case api.Outcome_DEFAULT:
		return errors.New("default outcome found: " + sr.Outcome.String())
	case api.Outcome_DRAW:
		Domain.PrintDraw()
		return nil
	case api.Outcome_WON:
		// TODO:
		// if sr.Player == nil {
		// 	return nilerr
		// }
		p := game.NewPlayer(sr.Player.Mark, int32(sr.Player.Num))
		Domain.PrintWinner(p)
		return nil
	default:
		return errors.New("unknown outcome: " + sr.Outcome.String())
	}
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
