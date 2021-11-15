package main

import (
	"fmt"
	"strings"
)

func (m Move) string(b Board) string {
	str := fmt.Sprintf("%s from %s to %s", b[m.Mover].string(), b[m.Mover].Coord.string(), m.Coord.string())
	if m.Steal >= 0 {
		str += fmt.Sprintf(" (takes %s)", b[m.Steal].string())
	}
	return str
}

func (b Board) string() string {
	rows := make([]string, boardSize)
	for r := 0; r < boardSize; r++ {
		cols := make([]string, boardSize)
		for c := 0; c < boardSize; c++ {
			p := pieceAtCoordinate(b, Coordinate{Row: r, Col: c})
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

func (c Coordinate) string() string {
	return fmt.Sprintf("%d,%d", c.Col, c.Row)
}

func (c Color) string() string {
	switch c {
	case white:
		return "white"
	case black:
		return "black"
	default:
		return ""
	}
}

func (p Piece) string() string {
	switch p.Color {
	case white:
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
	case black:
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
