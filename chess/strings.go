package chess

import (
	"fmt"
	"strings"
)

func (m Move) String(b Board) string {
	str := fmt.Sprintf("%s from %s to %s", b[m.Mover].String(), b[m.Mover].Coord.String(), m.Coord.String())
	if m.Steal >= 0 {
		str += fmt.Sprintf(" (takes %s)", b[m.Steal].String())
	}
	return str
}

func (b Board) String() string {
	rows := make([]string, boardSize)
	for r := 0; r < boardSize; r++ {
		cols := make([]string, boardSize)
		for c := 0; c < boardSize; c++ {
			p := pieceAtCoordinate(b, Coordinate{Row: r, Col: c})
			if p >= 0 {
				cols[c] = b[p].String()
			} else {
				cols[c] = " "
			}
		}
		rows[r] = strings.Join(cols, "")
	}
	return strings.Join(rows, "\n")
}

func (c Coordinate) String() string {
	return fmt.Sprintf("%d,%d", c.Col, c.Row)
}

func (c Color) String() string {
	switch c {
	case White:
		return "White"
	case Black:
		return "Black"
	default:
		return ""
	}
}

func (p Piece) String() string {
	switch p.Color {
	case White:
		switch p.Rank {
		case pawn:
			return "♟︎"
		case rook:
			return "♜"
		case knight:
			return "♞"
		case bishop:
			return "♝"
		case king:
			return "♚"
		case queen:
			return "♛"
		default:
			return ""
		}
	case Black:
		switch p.Rank {
		case pawn:
			return "♙"
		case rook:
			return "♖"
		case knight:
			return "♘"
		case bishop:
			return "♗"
		case king:
			return "♔"
		case queen:
			return "♕"
		default:
			return ""
		}
	default:
		return ""
	}
}
