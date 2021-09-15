package client

import (
	"errors"
	"fmt"
	"log"
	"tictactoeweb/internal/domain/game"

	api "tictactoeweb/api"

	. "tictactoeweb/internal"
)

// Constants

func ErrNilArgument() error {
	return errors.New("status loop: cannot react on nil argument")
}

func RunStatusLoop(ctx Ctx) error {
	log.Print("Client: running status loop...")
	for {
		args := &api.QueryRequest{Query: api.Querys_GET_STATUS}
		log.Printf("Client: calling remote. args: %v...", args)
		resp, err := Api().RunQuery(ctx, args)
		if err != nil {
			return err
		}
		if resp == nil {
			return ErrNilArgument()
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
		return Play(ctx, *sr.Player(), *sr.Board()) // TODO: parse players
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
		if sr.Player == nil {
			return ErrNilArgument()
		}
		p := game.NewPlayer(sr.Player.Mark, int32(sr.Player.Num))
		Domain.PrintWinner(p)
		return nil
	default:
		return errors.New("unknown outcome: " + sr.Outcome.String())
	}
}
