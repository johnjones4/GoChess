package chess

func startMinimaxLocal(b Board, c Color) Move {
	root := newNode(nil, b, nil, 0, c, false)
	node, _ := minimaxLocal(&root, minInt, maxInt, 6)
	return *node.Edge
}

func minimaxLocal(n *Node, alpha, beta int, maxDepth int) (*Node, int) {
	if !n.IsLeaf && n.Depth != maxDepth {
		n.nextLevel()
	}
	if n.IsLeaf || n.Depth == maxDepth {
		n.evaluate()
		return nil, n.Value
	}
	if !n.IsOpponent {
		bestV := minInt
		var bestChild *Node = nil
		for i := range n.Children {
			child := &n.Children[i]
			_, v := minimaxLocal(child, alpha, beta, maxDepth)
			if bestChild == nil || v >= bestV {
				bestChild = child
				bestV = v
			}
			if bestV >= beta {
				break
			}
			alpha = max(alpha, bestV)
		}
		if bestChild == nil {
			panic("no move")
		}
		n.Value = bestV
		return bestChild, bestV
	} else {
		var worstChild *Node = nil
		var worstV = maxInt
		for i := range n.Children {
			child := &n.Children[i]
			_, v := minimaxLocal(child, alpha, beta, maxDepth)
			if worstChild == nil || v <= worstV {
				worstV = v
				worstChild = child
			}
			if worstV <= alpha {
				break
			}
			beta = min(beta, worstV)
		}
		if worstChild == nil {
			panic("no move")
		}
		n.Value = worstV
		return worstChild, worstV
	}
}
