package chess

func startMinimaxLocal(b Board, c Color) Move {
	root := newNode(&b, nil, 0, c, false)
	move, _ := minimaxLocal(root, minInt, maxInt, 3)
	return move
}

func minimaxLocal(n Node, alpha, beta int, maxDepth int) (Move, int) {
	n.nextLevel(n.Depth + 1)
	if n.IsLeaf || n.Depth == maxDepth {
		n.evaluate()
		return Move{}, n.Value
	}
	if !n.IsOpponent {
		v := minInt
		var m *Move = nil
		for _, child := range n.Children {
			_, v1 := minimaxLocal(child, alpha, beta, maxDepth)
			if v1 > v {
				v = v1
				m = child.Edge
			}
			if v >= beta {
				break
			}
			alpha = max(alpha, v)
		}
		if m == nil {
			panic("no move")
		}
		return *m, v
	} else {
		v := maxInt
		var m *Move = nil
		for _, child := range n.Children {
			_, v1 := minimaxLocal(child, alpha, beta, maxDepth)
			if v1 < v {
				v = v1
				m = child.Edge
			}
			if v <= alpha {
				break
			}
			beta = min(beta, v)
		}
		if m == nil {
			panic("no move")
		}
		return *m, v
	}
}
