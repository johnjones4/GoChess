package main

func (b board) doMove(m move) board {
	if b[m.mover].stolen {
		panic("cannot move a stolen piece!")
	}
	if m.steal >= 0 && b[m.steal].stolen {
		panic("cannot steal a stolen piece!")
	}
	b1 := b.copy()
	b1[m.mover].coord = m.coord
	if m.steal >= 0 {
		b1[m.steal].stolen = true
	}
	return b1
}

func (b board) copy() board {
	b1 := make(board, boardSize*4)
	copy(b1, b)
	return b1
}

func (board board) moves(c color) []move {
	moves := make([]move, 0)
	for i, p := range board {
		if p.color == c && !p.stolen {
			switch p.rank {
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
