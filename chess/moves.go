package chess

func pawnMoves(board Board, pi int) []Move {
	p := board[pi]
	direction := 1
	if p.Color == White {
		direction = -1
	}
	deltas := make([]delta, 0)
	d := delta{direction, 0}
	c := moveDelta(p.Coord, d)
	targetPiece := pieceAtCoordinate(board, c)
	if isValidCoordinate(c) && targetPiece < 0 {
		deltas = append(deltas, d)

		if (p.Color == Black && p.Coord.Row == 1) || (p.Color == White && p.Coord.Row == 6) {
			d = delta{direction * 2, 0}
			c = moveDelta(p.Coord, d)
			targetPiece = pieceAtCoordinate(board, c)
			if isValidCoordinate(c) && targetPiece < 0 {
				deltas = append(deltas, d)
			}
		}
	}

	if (p.Color == Black && p.Coord.Row == 1) || (p.Color == White && p.Coord.Row == boardSize-2) {
		deltas = append(deltas, delta{direction * 2, 0})
	}
	if p.Coord.Row+direction >= 0 && p.Coord.Row+direction < boardSize {
		if p.Coord.Col-1 >= 0 {
			d := delta{direction, -1}
			c1 := moveDelta(p.Coord, d)
			targetPiece := pieceAtCoordinate(board, c1)
			if targetPiece >= 0 && board[targetPiece].Color != p.Color {
				deltas = append(deltas, d)
			}
		}
		if p.Coord.Col+1 < 8 {
			d := delta{direction, 1}
			c1 := moveDelta(p.Coord, d)
			targetPiece := pieceAtCoordinate(board, c1)
			if targetPiece >= 0 && board[targetPiece].Color != p.Color {
				deltas = append(deltas, d)
			}
		}
	}
	moves := make([]Move, 0)
	for _, d := range deltas {
		targetCoord := Coordinate{
			Row: p.Coord.Row + d.dr,
			Col: p.Coord.Col + d.dc,
		}
		targetPiece := pieceAtCoordinate(board, targetCoord)
		if targetPiece < 0 || board[targetPiece].Color != p.Color {
			moves = append(moves, Move{
				Mover: pi,
				Coord: targetCoord,
				Steal: targetPiece,
			})
		}
	}
	return moves
}

func rookMoves(board Board, pi int) []Move {
	deltas := []delta{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}
	return searchDirections(board, pi, deltas)
}

func knightMoves(board Board, pi int) []Move {
	deltas := []delta{
		{2, 1},
		{1, 2},
		{-2, 1},
		{-1, 2},
		{2, -1},
		{1, -2},
		{-2, -1},
		{-1, -2},
	}
	return searchDeltas(board, pi, deltas)
}

func bishopMoves(board Board, pi int) []Move {
	deltas := []delta{
		{1, 1},
		{-1, -1},
		{-1, 1},
		{1, -1},
	}
	return searchDirections(board, pi, deltas)
}

func queenMoves(board Board, pi int) []Move {
	deltas := []delta{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
		{1, 1},
		{-1, -1},
		{-1, 1},
		{1, -1},
	}
	return searchDirections(board, pi, deltas)
}

func kingMoves(board Board, pi int) []Move {
	deltas := []delta{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
		{1, 1},
		{-1, -1},
		{-1, 1},
		{1, -1},
	}
	return searchDeltas(board, pi, deltas)
}

func pieceAtCoordinate(board Board, coord Coordinate) int {
	for i := range board {
		p := &board[i]
		if p.Coord.Row == coord.Row && p.Coord.Col == coord.Col && !p.Stolen {
			return i
		}
	}
	return -1
}

func extendDelta(d delta, a int) delta {
	return delta{
		dr: d.dr * a,
		dc: d.dc * a,
	}
}

func moveDelta(c Coordinate, d delta) Coordinate {
	return Coordinate{
		Row: c.Row + d.dr,
		Col: c.Col + d.dc,
	}
}

func isValidCoordinate(c Coordinate) bool {
	return c.Row >= 0 && c.Row < boardSize && c.Col >= 0 && c.Col < boardSize
}

func intInArray(a []int, i int) bool {
	for _, v := range a {
		if v == i {
			return true
		}
	}
	return false
}

func searchDirections(board Board, pi int, deltas []delta) []Move {
	p := board[pi]
	moves := make([]Move, 0)
	skipDeltas := make([]int, 0)
	for c := 1; c < boardSize; c++ {
		deltas1 := make([]delta, len(deltas))
		for i, d := range deltas {
			deltas1[i] = extendDelta(d, c)
		}
		for i, d := range deltas1 {
			if !intInArray(skipDeltas, i) {
				targetCoord := moveDelta(p.Coord, d)
				if isValidCoordinate(targetCoord) {
					targetPiece := pieceAtCoordinate(board, targetCoord)
					if targetPiece < 0 || board[targetPiece].Color != p.Color {
						moves = append(moves, Move{
							Mover: pi,
							Coord: targetCoord,
							Steal: targetPiece,
						})
					}
					if targetPiece >= 0 {
						skipDeltas = append(skipDeltas, i)
					}
				}
			}
		}
	}
	return moves
}

func searchDeltas(board Board, pi int, deltas []delta) []Move {
	p := board[pi]
	moves := make([]Move, 0)
	for _, d := range deltas {
		targetCoord := moveDelta(p.Coord, d)
		if isValidCoordinate(targetCoord) {
			targetPiece := pieceAtCoordinate(board, targetCoord)
			if targetPiece < 0 || board[targetPiece].Color != p.Color {
				moves = append(moves, Move{
					Mover: pi,
					Coord: targetCoord,
					Steal: targetPiece,
				})
			}
		}
	}
	return moves
}
