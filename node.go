package main

func (n *Node) evaluate() {
	n.Value = 0
	minDistanceToKing := MaxInt
	var kingCoord *Coordinate = nil
	for _, p := range n.Board {
		if p.Color != n.Player && p.Rank == king {
			kingCoord = &p.Coord
			break
		}
	}
	if kingCoord == nil {
		panic("no king!")
	}
	for _, p := range n.Board {
		if p.Color != n.Player && p.Stolen {
			n.Value += (int(p.Rank) ^ 2)
		}
		if p.Color == n.Player {
			distance := abs(p.Coord.Col-kingCoord.Col) + abs(p.Coord.Row-kingCoord.Row)
			if distance < minDistanceToKing {
				minDistanceToKing = distance
			}
		}
	}
	n.Value -= (minDistanceToKing * 2)
}

func (n *Node) nextLevel() {
	moves := n.Board.moves(n.Player)
	n.Children = make([]Node, len(moves))
	n.IsLeaf = len(moves) == 0
	for i, m := range moves {
		n.Children[i] = newNode(&n.Board, &m, opposite(n.Player), !n.IsOpponent)
	}
}

func newNode(b *Board, m *Move, c Color, opponent bool) Node {
	b1 := b.copy()
	if m != nil {
		b1 = b1.doMove(*m)
	}

	n := Node{
		Board:      b1,
		Edge:       m,
		IsOpponent: opponent,
		Player:     c,
		IsLeaf:     false,
		Value:      0,
	}

	return n
}

func minimax(n Node, alpha, beta, depth int) (Move, int) {
	n.nextLevel()
	if n.IsLeaf || depth == 0 {
		n.evaluate()
		return Move{}, n.Value
	}
	if !n.IsOpponent {
		v := MinInt
		var m *Move = nil
		for _, child := range n.Children {
			_, v1 := minimax(child, alpha, beta, depth-1)
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
		v := MaxInt
		var m *Move = nil
		for _, child := range n.Children {
			_, v1 := minimax(child, alpha, beta, depth-1)
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
