package server

// func (_Games) ArrangePlayers(m Mark) (Player, Player) {
// 	if strings.ToUpper(m) == X {
// 		return NewPlayer(X, 1), NewPlayer(O, 2)
// 	}
// 	return NewPlayer(O, 1), NewPlayer(X, 2)
// }

// // if brd.IsWinner(plr.Mark()) {
// // 	domain.PrintWinner(plr)
// // 	return domain.Games.SetBoard(gam, brd), false
// // }
// // if !brd.HasEmpty() {
// // 	domain.PrintDraw()
// // 	return domain.Games.SetBoard(gam, brd), false
// // }
// // return domain.Games.SetBoard(gam, brd), true

// func (_Games) SetBoard(g Game, b Board) Game {
// 	// TODO: send to server
// 	return /*updated*/ Game{}
// }

// // Other

// func (k Key) IsKey() bool {
// 	n, err := strconv.Atoi(string(k))
// 	if err != nil {
// 		return false
// 	}
// 	return n >= 1 && n <= 9
// }

// // Constants, Private

// func _coords() coords {
// 	return coords{
// 		"1": {0, 0}, "2": {0, 1}, "3": {0, 2},
// 		"4": {1, 0}, "5": {1, 1}, "6": {1, 2},
// 		"7": {2, 0}, "8": {2, 1}, "9": {2, 2},
// 	}
// }

// func (b Board) SetBoard(gr grid) Board {
// 	b.grid = gr
// 	return b
// }

// func (b Board) IsFilled(c Cell) bool {
// 	// WARN: possible out of range
// 	return b.grid[c.Row()][c.Col()] != Gap
// }

// // Checks

// // Party:Server
// func (b Board) IsEmpty() Empty {
// 	grd := b.grid
// 	return b == Board{} || b == Dead() || len(grd) != Size ||
// 		len(grd[0]) != Size ||
// 		len(grd[1]) != Size ||
// 		len(grd[2]) != Size
// }

// // Party:Server
// func (b Board) HasEmpty() bool {
// 	for _, row := range b.grid {
// 		for _, m := range row {
// 			if m == Gap {
// 				return true
// 			}
// 		}
// 	}
// 	return false
// }

// // Party:Server
// func (b Board) IsWinner(m Mark) bool {
// 	grd := b.grid
// 	// Horizontal
// 	h0 := grd[0][0] == m && grd[0][1] == m && grd[0][2] == m // 1 1 1 -> 7
// 	h1 := grd[1][0] == m && grd[1][1] == m && grd[1][2] == m // - - -
// 	h2 := grd[2][0] == m && grd[2][1] == m && grd[2][2] == m // - - -
// 	// Vertical
// 	v0 := grd[0][0] == m && grd[1][0] == m && grd[2][0] == m
// 	v1 := grd[0][1] == m && grd[1][1] == m && grd[2][1] == m
// 	v2 := grd[0][2] == m && grd[1][2] == m && grd[2][2] == m
// 	// Diagonal
// 	d0 := grd[0][0] == m && grd[1][1] == m && grd[2][2] == m
// 	d1 := grd[0][2] == m && grd[1][1] == m && grd[2][0] == m

// 	return h0 || h1 || h2 || v0 || v1 || v2 || d0 || d1
// }

// // type grid [Size][Size]Mark // Party:Server

// type grid = string // Party:Client
// type Row = [Size]string

// // Constants, Public

// const (
// 	Size = 3
// 	Gap  = __
// 	X    = "X"
// 	O    = "O"

// 	X_x = "X_x"
// )

// const __ = "-" // Private

// func BlankBoard() Board {
// 	return Board{
// 		id: __,
// 		// Party:Client
// 		grid: strings.Join(Row{
// 			strings.Join(Row{__, __, __}, " "),
// 			strings.Join(Row{__, __, __}, " "),
// 			strings.Join(Row{__, __, __}, " "),
// 		}, "\n"),

// 		// Party:Server
// 		// grid{
// 		// 	{__, __, __},
// 		// 	{__, __, __},
// 		// 	{__, __, __},
// 		// },
// 	}
// }

// func Dead() Board {
// 	return Board{
// 		id: X_x,
// 		// Party:Client
// 		grid: strings.Join(Row{
// 			strings.Join(Row{X_x, X_x, X_x}, " "),
// 			strings.Join(Row{X_x, X_x, X_x}, " "),
// 			strings.Join(Row{X_x, X_x, X_x}, " "),
// 		}, "\n"),

// 		// Party:Server
// 		// grid{
// 		// 	{X_x, X_x, X_x},
// 		// 	{X_x, X_x, X_x},
// 		// 	{X_x, X_x, X_x},
// 		// },
// 	}
// }

// func (g Game) IsReady() bool {
// 	return 	!g.player1.IsEmpty() &&
// 		!g.player2.IsEmpty() &&
// 		!g.board.IsEmpty()
// }

// // Other, Public

// // Party:Server
// func (b Board) String() string {
// 	var dump Row
// 	for _, row := range b.grid {
// 		s := strings.Join(row[:], " ")
// 		dump = append(dump, s)
// 	}
// 	return strings.Join(dump, "\n")
// }

// func (k Key) ToCell() Cell {
// 	return _coords()[k] // TODO: detect and propagate errors?
// }
// // Party:Server
// type Game struct {
// 	id Id

// 	// board   Board // Party:Server
// 	board Board

// 	player1 Player
// 	player2 Player
// 	// Party:Client
// 	reader Reader
// }

// Cell struct {
// 	row, col Len
// }
// // Private

// type (
// 	coords map[Key]Cell
// )

// // Props

// func (c Cell) Row() int {
// 	return c.row
// }

// func (c Cell) Col() int {
// 	return c.col
// }
// Len = int  // 1..3
