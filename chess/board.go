package chess

import "errors"

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
			moves1, _ := board.MovesForPiece(i)
			moves = append(moves, moves1...)
		}
	}
	return moves
}

func (board Board) MovesForPiece(i int) ([]Move, error) {
	if i >= len(board) {
		return nil, errors.New("illegal piece index")
	}
	p := board[i]
	if p.Stolen {
		return nil, errors.New("cannot move stolen piece")
	}
	switch p.Rank {
	case pawn:
		return pawnMoves(board, i), nil
	case rook:
		return rookMoves(board, i), nil
	case knight:
		return knightMoves(board, i), nil
	case bishop:
		return bishopMoves(board, i), nil
	case queen:
		return queenMoves(board, i), nil
	case king:
		return kingMoves(board, i), nil
	default:
		return nil, errors.New("bad rank")
	}
}

func (board Board) MoveIsValid(color Color, move Move) bool {
	if move.Mover >= len(board) {
		return false
	}
	p := board[move.Mover]
	if p.Color != color {
		return false
	}
	moves, err := board.MovesForPiece(move.Mover)
	if err != nil {
		return false
	}
	for _, m := range moves {
		if coordsEqual(m.Coord, move.Coord) && m.Mover == move.Mover && m.Steal == move.Steal {
			return true
		}
	}
	return false
}
