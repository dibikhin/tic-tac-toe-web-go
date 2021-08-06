package server

import (
	"log"

	api "tictactoeweb/api"
)

type server struct {
	api.UnimplementedGameServer
}

// Public

func (s *server) GetStatus(ctx Ctx, m *api.Empty) (*api.StatusReply, error) {
	log.Printf("Recieved: GetStatus(), args: %v", m)
	sr := &api.StatusReply{
		State: api.State_IDLE,
		Actions: []api.Actions{
			api.Actions_START_GAME,
			api.Actions_GET_STATUS,
		},
		Message: "You can:",
	}
	log.Print(sr)
	return sr, nil
}

func (s *server) Run(ctx Ctx, cr *api.CommandRequest) (*api.StatusReply, error) {
	log.Printf("Recieved: Run(), args: %v", cr)
	sr := &api.StatusReply{
		State: api.State_WAITING,
		For:   api.For_AUTH,
		Actions: []api.Actions{
			api.Actions_GET_STATUS,
		},
		Message: "You can:",
	}
	log.Print(sr)
	return sr, nil
}
