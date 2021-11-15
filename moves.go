package main

func pawnMoves(board board, pi int) []move {
	p := board[pi]
	direction := 1
	if p.color == white {
		direction = -1
	}
	deltas := make([]delta, 0)
	d := delta{direction, 0}
	c := moveDelta(p.coord, d)
	targetPiece := pieceAtCoordinate(board, c)
	if isValidCoordinate(c) && targetPiece < 0 {
		deltas = append(deltas, d)

		if (p.color == black && p.coord.row == 1) || (p.color == white && p.coord.row == 6) {
			d = delta{direction * 2, 0}
			c = moveDelta(p.coord, d)
			targetPiece = pieceAtCoordinate(board, c)
			if isValidCoordinate(c) && targetPiece < 0 {
				deltas = append(deltas, d)
			}
		}
	}

	if (p.color == black && p.coord.row == 1) || (p.color == white && p.coord.row == boardSize-2) {
		deltas = append(deltas, delta{direction * 2, 0})
	}
	if p.coord.row+direction >= 0 && p.coord.row+direction < boardSize {
		if p.coord.col-1 >= 0 {
			d := delta{direction, -1}
			c1 := moveDelta(p.coord, d)
			targetPiece := pieceAtCoordinate(board, c1)
			if targetPiece >= 0 && board[targetPiece].color != p.color {
				deltas = append(deltas, d)
			}
		}
		if p.coord.col+1 < 8 {
			d := delta{direction, 1}
			c1 := moveDelta(p.coord, d)
			targetPiece := pieceAtCoordinate(board, c1)
			if targetPiece >= 0 && board[targetPiece].color != p.color {
				deltas = append(deltas, d)
			}
		}
	}
	moves := make([]move, 0)
	for _, d := range deltas {
		targetCoord := coordinate{
			row: p.coord.row + d.dr,
			col: p.coord.col + d.dc,
		}
		targetPiece := pieceAtCoordinate(board, targetCoord)
		if targetPiece < 0 || board[targetPiece].color != p.color {
			moves = append(moves, move{
				mover: pi,
				coord: targetCoord,
				steal: targetPiece,
			})
		}
	}
	return moves
}

func rookMoves(board board, pi int) []move {
	deltas := []delta{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}
	return searchDirections(board, pi, deltas)
}

func knightMoves(board board, pi int) []move {
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

func bishopMoves(board board, pi int) []move {
	deltas := []delta{
		{1, 1},
		{-1, -1},
		{-1, 1},
		{1, -1},
	}
	return searchDirections(board, pi, deltas)
}

func queenMoves(board board, pi int) []move {
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

func kingMoves(board board, pi int) []move {
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

func pieceAtCoordinate(board board, coord coordinate) int {
	for i := range board {
		p := &board[i]
		if p.coord.row == coord.row && p.coord.col == coord.col && !p.stolen {
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

func moveDelta(c coordinate, d delta) coordinate {
	return coordinate{
		row: c.row + d.dr,
		col: c.col + d.dc,
	}
}

func isValidCoordinate(c coordinate) bool {
	return c.row >= 0 && c.row < boardSize && c.col >= 0 && c.col < boardSize
}

func intInArray(a []int, i int) bool {
	for _, v := range a {
		if v == i {
			return true
		}
	}
	return false
}

func searchDirections(board board, pi int, deltas []delta) []move {
	p := board[pi]
	moves := make([]move, 0)
	skipDeltas := make([]int, 0)
	for c := 0; c < boardSize; c++ {
		deltas1 := make([]delta, len(deltas))
		for i, d := range deltas {
			deltas1[i] = extendDelta(d, c)
		}
		for i, d := range deltas1 {
			if !intInArray(skipDeltas, i) {
				targetCoord := moveDelta(p.coord, d)
				if isValidCoordinate(targetCoord) {
					targetPiece := pieceAtCoordinate(board, targetCoord)
					if targetPiece < 0 || board[targetPiece].color != p.color {
						moves = append(moves, move{
							mover: pi,
							coord: targetCoord,
							steal: targetPiece,
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

func searchDeltas(board board, pi int, deltas []delta) []move {
	p := board[pi]
	moves := make([]move, 0)
	for _, d := range deltas {
		targetCoord := moveDelta(p.coord, d)
		if isValidCoordinate(targetCoord) {
			targetPiece := pieceAtCoordinate(board, targetCoord)
			if targetPiece < 0 || board[targetPiece].color != p.color {
				moves = append(moves, move{
					mover: pi,
					coord: targetCoord,
					steal: targetPiece,
				})
			}
		}
	}
	return moves
}
