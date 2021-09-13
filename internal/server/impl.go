package server

import (
	"log"

	api "tictactoeweb/api"

	. "tictactoeweb/internal"
	. "tictactoeweb/internal/server/game"
)

type server struct {
	api.UnimplementedGameServer
}

// Public

func (s *server) RunCommand(ctx Ctx, cr *api.CommandRequest) (*api.StatusReply, error) {
	log.Printf("Recieved command: args: %v", cr)

	switch cr.Action {
	// case api.Actions_NOOP:
	// 	return errors.New("default state found: " + sr.State.String())
	case api.Actions_SET_MARK:
	}
	// default:
	sr := &api.StatusReply{
		State: api.Is_WAITING,
		For:   api.For_MARK,
		// Actions: []api.Actions{
		// 	api.Actions_GET_STATUS,
		// },
	}
	log.Print(sr)
	return sr, nil
}

func (s *server) RunQuery(ctx Ctx, cr *api.QueryRequest) (*api.StatusReply, error) {
	log.Printf("Recieved query: args: %v", cr)

	switch cr.Query {
	// case api.Querys_EMPTY:
	// 	return errors.New("default query found: " + cr.Query.String())

	case api.Querys_IS_FILLED:
		isFilled, err := Domain.Boards.IsFilled(cr.BoardId, ServKey(cr.Key))
		sr := &api.StatusReply{
			IsFilled: isFilled,
			State:    api.Is_WAITING,
			For:      api.For_TURN,
			Querys: []api.Querys{
				api.Querys_GET_STATUS,
			},
			Actions: []api.Actions{
				api.Actions_DO_TURN,
			},
		}
		log.Print(sr)
		return sr, err
	case api.Querys_GET_STATUS:

	}

	// default:
	// unknown

	// sr := &api.StatusReply{
	// 	State: api.Is_WAITING,
	// 	For:   api.For_TURN,
	// 	Querys: []api.Querys{
	// 		api.Querys_GET_STATUS,
	// 	},
	// 	Actions: []api.Actions{
	// 		api.Actions_DO_TURN,
	// 	},
	// // }
	// log.Print(sr)
	// return sr, nil
	return nil, nil
}
