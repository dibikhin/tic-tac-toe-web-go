package server

import (
	"log"

	api "tictactoeweb/api"

	. "tictactoeweb/internal"
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
		Actions: []api.Actions{
			api.Actions_GET_STATUS,
		},
	}
	log.Print(sr)
	return sr, nil
}

func (s *server) RunQuery(ctx Ctx, cr *api.QueryRequest) (*api.StatusReply, error) {
	log.Printf("Recieved command: args: %v", cr)
	sr := &api.StatusReply{
		State: api.Is_WAITING,
		For:   api.For_TURN,
		Actions: []api.Actions{
			api.Actions_GET_STATUS,
		},
	}
	log.Print(sr)
	return sr, nil
}
