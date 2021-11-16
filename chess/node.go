package chess

func (n *Node) evaluate() {
	n.Value = 0
	minDistanceToKing := maxInt
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
	if n.IsOpponent {
		n.Value = -n.Value
	}
}

func (n *Node) nextLevel(depth int) {
	moves := n.Board.moves(n.Player)
	n.Children = make([]Node, len(moves))
	n.IsLeaf = len(moves) == 0
	for i, m := range moves {
		n.Children[i] = newNode(&n.Board, &m, depth, opposite(n.Player), !n.IsOpponent)
	}
}

func newNode(b *Board, m *Move, depth int, c Color, opponent bool) Node {
	b1 := b.copy()
	if m != nil {
		b1 = b1.doMove(*m)
	}

	n := Node{
		Depth:      depth,
		Board:      b1,
		Edge:       m,
		IsOpponent: opponent,
		Player:     c,
		IsLeaf:     false,
		Value:      0,
	}

	return n
}
