package ai

import (
	"github.com/johnjones4/GoChess/chess/core"
)

type Node struct {
	Depth      int        `json:"depth"`
	Board      core.Board `json:"board"`
	Edge       *core.Move `json:"-"`
	Parent     *Node      `json:"-"`
	Children   []Node     `json:"children"`
	IsOpponent bool       `json:"isOpponent"`
	Player     core.Color `json:"player"`
	IsLeaf     bool       `json:"isLeaf"`
	Value      float64    `json:"value"`
}

// func (n *Node) string() string {
// 	str := strings.Builder{}
// 	n._string(&str, 0)
// 	return str.String()
// }

// func (n *Node) _string(str *strings.Builder, tabs int) {
// 	for i := 0; i < tabs; i++ {
// 		str.WriteString(" ")
// 	}
// 	str.WriteString("* ")
// 	str.WriteString(fmt.Sprint(n.Value))
// 	if n.IsLeaf {
// 		str.WriteString("-L")
// 	}
// 	if n.Edge != nil && n.Parent != nil {
// 		str.WriteString(" " + n.Edge.String(n.Parent.Board))
// 	}
// 	str.WriteString("\n")
// 	for _, c := range n.Children {
// 		c._string(str, tabs+1)
// 	}
// }

func (n *Node) Evaluate() {
	n.Value = 0
	for _, p := range n.Board {
		if p.Rank == core.King && p.Stolen {
			n.Value = 1.0
			if n.IsOpponent {
				n.Value = -1.0
			}
			return
		}
	}
	piecesOnBoard := n.pcntPieceOnBoard()
	avgDistanceToKing := 1.0 - (n.avgDistanceToKing() / 16.0)

	n.Value = (piecesOnBoard * 0.75) + (avgDistanceToKing * 0.25)
	if n.IsOpponent {
		n.Value = -n.Value
	}
}

func (n *Node) NextLevel() {
	if n.IsLeaf {
		return
	}
	moves := n.Board.Moves(n.Player)
	n.Children = make([]Node, len(moves))
	n.IsLeaf = len(moves) == 0
	for i, m := range moves {
		m1 := m
		n.Children[i] = NewNode(n, n.Board, &m1, n.Depth+1, core.Opposite(n.Player), !n.IsOpponent)
	}
}

// func (n *Node) build(toDepth int) {
// 	if n.Depth == toDepth {
// 		return
// 	}
// 	n.NextLevel()
// 	for i := range n.Children {
// 		n.Children[i].build(toDepth)
// 	}
// }

func NewNode(p *Node, b core.Board, m *core.Move, depth int, c core.Color, opponent bool) Node {
	b1 := b.Copy()
	if m != nil {
		b1 = b1.DoMove(*m)
	}

	n := Node{
		Depth:      depth,
		Board:      b1,
		Edge:       m,
		Parent:     p,
		IsOpponent: opponent,
		Player:     c,
		IsLeaf:     false,
		Value:      0,
	}

	for _, p := range b {
		if p.Rank == core.King && p.Stolen {
			n.IsLeaf = true
		}
	}

	return n
}
