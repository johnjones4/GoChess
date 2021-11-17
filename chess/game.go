package chess

func NewGame(userPlayer Color) Game {
	g := Game{
		Board:      make(Board, boardSize*4),
		Turn:       0,
		UserPlayer: userPlayer,
		Winner:     -1,
	}
	j := 0
	for i := 0; i < boardSize; i++ {
		g.Board[j] = Piece{
			Color: White,
			Rank:  pawn,
			Coord: Coordinate{
				Row: 6,
				Col: i,
			},
			Stolen: false,
		}
		j++

		g.Board[j] = Piece{
			Color: White,
			Rank:  lineUp[i],
			Coord: Coordinate{
				Row: 7,
				Col: i,
			},
			Stolen: false,
		}
		j++

		g.Board[j] = Piece{
			Color: Black,
			Rank:  pawn,
			Coord: Coordinate{
				Row: 1,
				Col: i,
			},
			Stolen: false,
		}
		j++

		g.Board[j] = Piece{
			Color: Black,
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

func (g *Game) TakeComputerTurn() bool {
	move := startMinimaxLocal(g.Board, opposite(g.UserPlayer))
	return g.TakeTurn(move)
}

func (g *Game) TakeTurn(move Move) bool {
	g.Board = g.Board.doMove(move)
	if move.Steal >= 0 {
		if g.Board[move.Steal].Rank == king {
			g.Winner = g.Board[move.Mover].Color
			return true
		}
	}
	g.Turn++
	return false
}
