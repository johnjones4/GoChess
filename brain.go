package main

import (
	"math/rand"
)

const (
	MaxInt = 10000000
	MinInt = -10000000
)

func bestMinimaxMove(b Board, c Color) Move {
	root := newNode(&b, nil, c, false)
	move, _ := minimax(root, MinInt, MaxInt, 3)
	return move
}

func randomMove(b Board, c Color) Move {
	moves := b.moves(c)
	return moves[rand.Intn(len(moves))]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func opposite(c Color) Color {
	if c == white {
		return black
	}
	return white
}
