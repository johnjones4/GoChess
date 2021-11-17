package chess

import (
	"fmt"
	"strings"
)

func (n *Node) string() string {
	str := strings.Builder{}
	n._string(&str, 0)
	return str.String()
}

func (n *Node) _string(str *strings.Builder, tabs int) {
	for i := 0; i < tabs; i++ {
		str.WriteString(" ")
	}
	str.WriteString("* ")
	str.WriteString(fmt.Sprint(n.Value))
	if n.IsLeaf {
		str.WriteString("-L")
	}
	if n.Edge != nil {
		str.WriteString(" " + n.Edge.String(n.Board))
	}
	str.WriteString("\n")
	for _, c := range n.Children {
		c._string(str, tabs+1)
	}
}

func (n *Node) evaluate() {
	n.Value = 0
	for _, p := range n.Board {
		if p.Rank == king && p.Stolen {
			n.Value = maxInt
			if n.IsOpponent {
				n.Value = -maxInt
			}
			return
		}
	}
	for _, p := range n.Board {
		if p.Color != n.Player && p.Stolen {
			n.Value += (int(p.Rank) ^ 2)
		}
	}
	if n.IsOpponent {
		n.Value = -n.Value
	}
}

func (n *Node) nextLevel(depth int) {
	if n.IsLeaf {
		return
	}
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

	for _, p := range *b {
		if p.Rank == king && p.Stolen {
			n.IsLeaf = true
		}
	}

	return n
}
