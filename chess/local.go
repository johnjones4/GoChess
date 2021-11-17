package chess

import "fmt"

func startMinimaxLocal(b Board, c Color) Move {
	root := newNode(&b, nil, 0, c, false)
	node := minimaxLocal(&root, minInt, maxInt, 2)
	fmt.Println(root.string())
	return *node.Edge
}

func minimaxLocal(n *Node, alpha, beta int, maxDepth int) *Node {
	if !n.IsLeaf && n.Depth != maxDepth {
		n.nextLevel(n.Depth + 1)
	}
	if n.IsLeaf || n.Depth == maxDepth {
		n.evaluate()
		return nil
	}
	if !n.IsOpponent {
		var bestNode *Node = nil
		for i := range n.Children {
			child := &n.Children[i]
			bestGrandchild := minimaxLocal(child, alpha, beta, maxDepth)
			if bestGrandchild != nil {
				child.Value = bestGrandchild.Value
			}
			if bestNode == nil || child.Value >= bestNode.Value {
				bestNode = child
			}
			if bestNode.Value >= beta {
				break
			}
			alpha = max(alpha, bestNode.Value)
		}
		if bestNode == nil {
			panic("no move")
		}
		n.Value = bestNode.Value
		return bestNode
	} else {
		var bestNode *Node = nil
		for i := range n.Children {
			child := &n.Children[i]
			bestGrandchild := minimaxLocal(child, alpha, beta, maxDepth)
			if bestGrandchild != nil {
				child.Value = bestGrandchild.Value
			}
			if bestNode == nil || child.Value <= bestNode.Value {
				bestNode = child
			}
			if bestNode.Value <= alpha {
				break
			}
			beta = min(beta, bestNode.Value)
		}
		if bestNode == nil {
			panic("no move")
		}
		n.Value = bestNode.Value
		return bestNode
	}
}
