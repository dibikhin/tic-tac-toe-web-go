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

// func (s *server) GetStatus(ctx Ctx, m *api.Empty) (*api.StatusReply, error) {
// 	log.Printf("Recieved: GetStatus(), args: %v", m)
// 	sr := &api.StatusReply{
// 		State: api.State_WAITING,
// 		For:   api.For_MARK,
// 		Actions: []api.Actions{
// 			api.Actions_SET_MARK,
// 			api.Actions_GET_STATUS,
// 		},
// 		Message: "You can:",
// 	}
// 	log.Print(sr)
// 	return sr, nil
// }

func (s *server) RunCommand(ctx Ctx, cr *api.CommandRequest) (*api.StatusReply, error) {
	log.Printf("Recieved: Run(), args: %v", cr)
	sr := &api.StatusReply{
		State: api.State_WAITING,
		For:   api.For_MARK,
		Actions: []api.Actions{
			api.Actions_GET_STATUS,
		},
		Message: "You can:",
	}
	log.Print(sr)
	return sr, nil
}

// // Commands

// func (s *server) ArrangePlayers(ctx Ctx, cr *api.CommandRequest) (*api.StatusReply, error) {
// 	return &api.StatusReply{}, nil
// }

// func (s *server) Turn(ctx Ctx, cr *api.CommandRequest) (*api.StatusReply, error) {
// 	return &api.StatusReply{}, nil
// }

// // Querys

// func (s *server) IsFilled(ctx Ctx, cr *api.CommandRequest) (*api.StatusReply, error) {
// 	return &api.StatusReply{}, nil
// }
