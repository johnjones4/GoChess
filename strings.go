package main

import (
	"fmt"
	"strings"
)

func (m move) string(b board) string {
	str := fmt.Sprintf("%s from %s to %s", b[m.mover].string(), b[m.mover].coord.string(), m.coord.string())
	if m.steal >= 0 {
		str += fmt.Sprintf(" (takes %s)", b[m.steal].string())
	}
	return str
}

func (b board) string() string {
	rows := make([]string, boardSize)
	for r := 0; r < boardSize; r++ {
		cols := make([]string, boardSize)
		for c := 0; c < boardSize; c++ {
			p := pieceAtCoordinate(b, coordinate{row: r, col: c})
			if p >= 0 {
				cols[c] = b[p].string()
			} else {
				cols[c] = " "
			}
		}
		rows[r] = strings.Join(cols, "")
	}
	return strings.Join(rows, "\n")
}

func (c coordinate) string() string {
	return fmt.Sprintf("%d,%d", c.col, c.row)
}

func (c color) string() string {
	switch c {
	case white:
		return "white"
	case black:
		return "black"
	default:
		return ""
	}
}

func (p piece) string() string {
	switch p.color {
	case white:
		switch p.rank {
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
	case black:
		switch p.rank {
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
