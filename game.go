package main

func newGame() game {
	g := game{
		board: make(board, boardSize*4),
		turn:  0,
	}
	j := 0
	for i := 0; i < boardSize; i++ {
		g.board[j] = piece{
			color: white,
			rank:  pawn,
			coord: coordinate{
				row: 6,
				col: i,
			},
			stolen: false,
		}
		j++

		g.board[j] = piece{
			color: white,
			rank:  lineUp[i],
			coord: coordinate{
				row: 7,
				col: i,
			},
			stolen: false,
		}
		j++

		g.board[j] = piece{
			color: black,
			rank:  pawn,
			coord: coordinate{
				row: 1,
				col: i,
			},
			stolen: false,
		}
		j++

		g.board[j] = piece{
			color: black,
			rank:  lineUp[i],
			coord: coordinate{
				row: 0,
				col: i,
			},
			stolen: false,
		}
		j++
	}
	return g
}

func (g *game) nextTurn(c color) bool {
	var move move
	if c == black {
		move = bestMinimaxMove(g.board, c)
	} else {
		move = randomMove(g.board, c)
	}
	g.board = g.board.doMove(move)
	if move.steal >= 0 {
		if g.board[move.steal].rank == king {
			return true
		}
	}
	return false
}

func (g *game) run() color {
	turns := []color{white, black}
	for {
		t := turns[g.turn%len(turns)]
		if g.nextTurn(t) {
			return t
		}
		g.turn++
	}
}
