package gameserver

import (
	"context"
	"reflect"
	"testing"

	"tictactoe/pkg/api"

	"github.com/stretchr/testify/assert"
)

func Test_gameService(t *testing.T) {
	repo := MakeGameRepo()
	s := NewService(repo)

	t.Run("Game loop", func(t *testing.T) {
		name1 := "name1"
		name2 := "name2"

		got1, err := s.GetGame(context.TODO(), &api.GameRequest{PlayerName: name1})
		if (err != nil) != false {
			t.Errorf("gameService.GetGame() error = %v, wantErr %v", err, false)
			return
		}
		r1 := &api.GameResponse{Status: api.GameStatus_NOT_STARTED}
		if !assert.Equal(t, r1, got1) {
			t.Errorf("gameService.GetGame() = %v, want %v", got1, r1)
		}

		got2, err := s.StartGame(context.TODO(), &api.GameRequest{PlayerName: name1})
		if (err != nil) != false {
			t.Errorf("gameService.StartGame() error = %v, wantErr %v", err, false)
			return
		}
		r2 := &api.EmptyResponse{}
		if !assert.Equal(t, r2, got2) {
			t.Errorf("gameService.StartGame() = %v, want %v", got2, r2)
		}

		got3, err := s.GetGame(context.TODO(), &api.GameRequest{PlayerName: name1})
		if (err != nil) != false {
			t.Errorf("gameService.GetGame() error = %v, wantErr %v", err, false)
			return
		}
		if got3.Player1.Mark != "X" || got3.Player1.Name != name1 || got3.Status != api.GameStatus_WAITING_P2_JOIN {
			t.Errorf("gameService.GetGame() = %v, want %v %v %v", got3, name1, "X", api.GameStatus_WAITING_P2_JOIN)
		}

		got4, err := s.StartGame(context.TODO(), &api.GameRequest{PlayerName: name2})
		if (err != nil) != false {
			t.Errorf("gameService.StartGame() error = %v, wantErr %v", err, false)
			return
		}
		r4 := &api.EmptyResponse{}
		if !assert.Equal(t, r4, got4) {
			t.Errorf("gameService.StartGame() = %v, want %v", got4, r4)
		}

		got5, err := s.GetGame(context.TODO(), &api.GameRequest{PlayerName: name2})
		if (err != nil) != false {
			t.Errorf("gameService.GetGame() error = %v, wantErr %v", err, false)
			return
		}
		if got5.Player1.Mark != "X" || got5.Player1.Name != name1 ||
			got5.Player2.Mark != "O" || got5.Player2.Name != name2 ||
			got5.Status != api.GameStatus_WAITING_P1_TO_TURN {
			t.Errorf("gameService.GetGame() = %v, want %v %v %v %v %v", got5, name1, name2, "X", "O", api.GameStatus_WAITING_P1_TO_TURN)
		}

		got6, err := s.Turn(context.TODO(), &api.TurnRequest{PlayerName: name1, Turn: "1"})
		if (err != nil) != false {
			t.Errorf("gameService.Turn() error = %v, wantErr %v", err, false)
			return
		}
		r6 := &api.EmptyResponse{}
		if !assert.Equal(t, r6, got6) {
			t.Errorf("gameService.Turn() = %v, want %v", got6, r6)
		}

		got7, err := s.GetGame(context.TODO(), &api.GameRequest{PlayerName: name2})
		if (err != nil) != false {
			t.Errorf("gameService.GetGame() error = %v, wantErr %v", err, false)
			return
		}
		if got7.Status != api.GameStatus_WAITING_P2_TO_TURN {
			t.Errorf("gameService.GetGame() = %v, want %v", got7, api.GameStatus_WAITING_P2_TO_TURN)
		}

		got8, err := s.Turn(context.TODO(), &api.TurnRequest{PlayerName: name2, Turn: "2"})
		if (err != nil) != false {
			t.Errorf("gameService.Turn() error = %v, wantErr %v", err, false)
			return
		}
		r7 := &api.EmptyResponse{}
		if !assert.Equal(t, r7, got8) {
			t.Errorf("gameService.Turn() = %v, want %v", got8, r7)
		}

		got9, err := s.GetGame(context.TODO(), &api.GameRequest{PlayerName: name2})
		if (err != nil) != false {
			t.Errorf("gameService.GetGame() error = %v, wantErr %v", err, false)
			return
		}
		if got9.Status != api.GameStatus_WAITING_P1_TO_TURN {
			t.Errorf("gameService.GetGame() = %v, want %v", got9, api.GameStatus_WAITING_P1_TO_TURN)
		}

		got10, err := s.Turn(context.TODO(), &api.TurnRequest{PlayerName: name2, Turn: "111"})
		if (err != nil) != false {
			t.Errorf("gameService.Turn() error = %v, wantErr %v", err, false)
			return
		}
		r8 := &api.EmptyResponse{}
		if !assert.Equal(t, r8, got10) {
			t.Errorf("gameService.Turn() = %v, want %v", got10, r8)
		}

		got11, err := s.StartGame(context.TODO(), &api.GameRequest{PlayerName: name1})
		if (err != nil) != false {
			t.Errorf("gameService.StartGame() error = %v, wantErr %v", err, false)
			return
		}
		r9 := &api.EmptyResponse{}
		if !assert.Equal(t, r9, got11) {
			t.Errorf("gameService.StartGame() = %v, want %v", got11, r9)
		}

		gs1, _ := repo.GetAll()
		g1 := gs1[0]
		g1.board = Board{
			{__, __, "X"},
			{__, "X", __},
			{__, __, __},
		}
		repo.UpdateByID(g1.id, g1)

		got13, err := s.Turn(context.TODO(), &api.TurnRequest{PlayerName: name1, Turn: "7"})
		if (err != nil) != false {
			t.Errorf("gameService.Turn() error = %v, wantErr %v", err, false)
			return
		}
		r11 := &api.EmptyResponse{}
		if !assert.Equal(t, r11, got13) {
			t.Errorf("gameService.Turn() = %v, want %v", got13, r11)
		}

		gs2, _ := repo.GetAll()
		g2 := gs2[0]
		g2.status = api.GameStatus_DRAW
		repo.UpdateByID(g2.id, g2)

		got12, err := s.StartGame(context.TODO(), &api.GameRequest{PlayerName: name1})
		if (err != nil) != false {
			t.Errorf("gameService.StartGame() error = %v, wantErr %v", err, false)
			return
		}
		r10 := &api.EmptyResponse{}
		if !assert.Equal(t, r10, got12) {
			t.Errorf("gameService.StartGame() = %v, want %v", got12, r10)
		}
	})
}

func Test_gameService_Turn(t *testing.T) {
	type args struct {
		req *api.TurnRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *api.EmptyResponse
		wantErr bool
	}{
		{"Player has no game", args{&api.TurnRequest{PlayerName: "name3", Turn: "1"}}, (*api.EmptyResponse)(nil), true},
	}
	gamesDB := []Game{}
	repo := MakeGameRepo(gamesDB...)
	s := NewService(repo)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.Turn(context.TODO(), tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("gameService.Turn() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("gameService.Turn() = %v, want %v", got, tt.want)
			}
		})
	}
}
