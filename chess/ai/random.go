package ai

import (
	"math/rand"

	"github.com/johnjones4/GoChess/chess/core"
)

func RandomMove(b core.Board, c core.Color) core.Move {
	moves := b.Moves(c)
	return moves[rand.Intn(len(moves))]
}
