package main

func (b Board) doMove(m Move) Board {
	if b[m.Mover].Stolen {
		panic("cannot move a stolen piece!")
	}
	if m.Steal >= 0 && b[m.Steal].Stolen {
		panic("cannot steal a stolen piece!")
	}
	b1 := b.copy()
	b1[m.Mover].Coord = m.Coord
	if m.Steal >= 0 {
		b1[m.Steal].Stolen = true
	}
	return b1
}

func (b Board) copy() Board {
	b1 := make(Board, boardSize*4)
	copy(b1, b)
	return b1
}

func (board Board) moves(c Color) []Move {
	moves := make([]Move, 0)
	for i, p := range board {
		if p.Color == c && !p.Stolen {
			switch p.Rank {
			case pawn:
				moves = append(moves, pawnMoves(board, i)...)
			case rook:
				moves = append(moves, rookMoves(board, i)...)
			case knight:
				moves = append(moves, knightMoves(board, i)...)
			case bishop:
				moves = append(moves, bishopMoves(board, i)...)
			case queen:
				moves = append(moves, queenMoves(board, i)...)
			case king:
				moves = append(moves, kingMoves(board, i)...)
			}
		}
	}
	return moves
}
