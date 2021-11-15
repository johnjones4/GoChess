package main

func NewGame() Game {
	g := Game{
		Board: make(Board, boardSize*4),
		Turn:  0,
	}
	j := 0
	for i := 0; i < boardSize; i++ {
		g.Board[j] = Piece{
			Color: white,
			Rank:  pawn,
			Coord: Coordinate{
				Row: 6,
				Col: i,
			},
			Stolen: false,
		}
		j++

		g.Board[j] = Piece{
			Color: white,
			Rank:  lineUp[i],
			Coord: Coordinate{
				Row: 7,
				Col: i,
			},
			Stolen: false,
		}
		j++

		g.Board[j] = Piece{
			Color: black,
			Rank:  pawn,
			Coord: Coordinate{
				Row: 1,
				Col: i,
			},
			Stolen: false,
		}
		j++

		g.Board[j] = Piece{
			Color: black,
			Rank:  lineUp[i],
			Coord: Coordinate{
				Row: 0,
				Col: i,
			},
			Stolen: false,
		}
		j++
	}
	return g
}

func (g *Game) nextTurn(c Color) bool {
	var move Move
	if c == black {
		move = bestMinimaxMove(g.Board, c)
	} else {
		move = randomMove(g.Board, c)
	}
	g.Board = g.Board.doMove(move)
	if move.Steal >= 0 {
		if g.Board[move.Steal].Rank == king {
			return true
		}
	}
	return false
}

func (g *Game) run() Color {
	turns := []Color{white, black}
	for {
		t := turns[g.Turn%len(turns)]
		if g.nextTurn(t) {
			return t
		}
		g.Turn++
	}
}
