package core

func NewGame(userPlayer Color, computerPlayerFn func(Board, Color) Move) (Game, error) {
	g := Game{
		Board:            make(Board, boardSize*4),
		Turn:             0,
		UserPlayer:       userPlayer,
		Winner:           -1,
		Log:              make([]LogItem, 0),
		computerPlayerFn: computerPlayerFn,
	}
	j := 0
	for i := 0; i < boardSize; i++ {
		g.Board[j] = Piece{
			Color: White,
			Rank:  Pawn,
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
			Rank:  Pawn,
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

	err := insertGame(&g)
	if err != nil {
		return Game{}, err
	}

	return g, nil
}

func (g *Game) TakeComputerTurn() (bool, error) {
	move := g.computerPlayerFn(g.Board, Opposite(g.UserPlayer))
	return g.TakeTurn(move)
}

func (g *Game) TakeTurn(move Move) (bool, error) {
	g.Log = append(g.Log, LogItem{
		Board: g.Board.Copy(),
		Move:  move,
	})
	g.Board = g.Board.DoMove(move)
	if move.Steal >= 0 {
		if g.Board[move.Steal].Rank == King {
			g.Winner = g.Board[move.Mover].Color
			err := updateGame(g)
			return true, err
		}
	}
	g.Turn++
	err := updateGame(g)
	return false, err
}

func (g *Game) ComputerGame() (Color, error) {
	players := []Color{White, Black}
	for {
		player := players[g.Turn%len(players)]
		move := g.computerPlayerFn(g.Board, player)
		winner, err := g.TakeTurn(move)
		if err != nil {
			return -1, err
		}
		if winner {
			return player, nil
		}
	}
}
