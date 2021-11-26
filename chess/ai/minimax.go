package ai

import (
	"math"

	"github.com/johnjones4/GoChess/chess/core"
)

func MinimaxLocal(b core.Board, c core.Color) core.Move {
	root := NewNode(nil, b, nil, 0, c, false)
	node, _ := Minimax(&root, -1.0, 1.0, 5)
	return *node.Edge
}

func Minimax(n *Node, alpha, beta float64, maxDepth int) (*Node, float64) {
	if !n.IsLeaf && n.Depth != maxDepth {
		n.NextLevel()
	}
	if n.IsLeaf || n.Depth == maxDepth {
		n.Evaluate()
		return nil, n.Value
	}
	if !n.IsOpponent {
		bestV := -1.0
		var bestChild *Node = nil
		for i := range n.Children {
			child := &n.Children[i]
			_, v := Minimax(child, alpha, beta, maxDepth)
			if bestChild == nil || v >= bestV {
				bestChild = child
				bestV = v
			}
			if bestV >= beta {
				break
			}
			alpha = math.Max(alpha, bestV)
		}
		if bestChild == nil {
			panic("no move")
		}
		n.Value = bestV
		return bestChild, bestV
	} else {
		var worstChild *Node = nil
		var worstV = 1.0
		for i := range n.Children {
			child := &n.Children[i]
			_, v := Minimax(child, alpha, beta, maxDepth)
			if worstChild == nil || v <= worstV {
				worstV = v
				worstChild = child
			}
			if worstV <= alpha {
				break
			}
			beta = math.Min(beta, worstV)
		}
		if worstChild == nil {
			panic("no move")
		}
		n.Value = worstV
		return worstChild, worstV
	}
}
