package game

import (
	. "tictactoeweb/internal/domain/game"
)

type (
	ServGame struct {
		Game
		board ServBoard
	}
	State = string
)

var (
	UNDEFINED        = ""
	WAITING_FOR_MARK = "waiting_for_mark"
	WAITING_FOR_TURN = "waiting_for_turn"
	GAME_OVER_P1_WON = "game_over_p1_won"
	GAME_OVER_P2_WON = "game_over_p2_won"
	GAME_OVER_DRAW   = "game_over_draw"
)

// Public

// Props

func (g ServGame) Board() ServBoard {
	return g.board
}

// // if brd.IsWinner(plr.Mark()) {
// // 	 PrintWinner(plr)
// // 	return  Games.SetBoard(gam, brd), false
// // }
// // if !brd.HasEmpty() {
// // 	 PrintDraw()
// // 	return  Games.SetBoard(gam, brd), false
// // }
// // return  Games.SetBoard(gam, brd), true
