package main

import (
	"math/rand"
)

const (
	MaxInt = 10000000
	MinInt = -10000000
)

func bestMinimaxMove(b board, c color) move {
	root := newNode(&b, nil, c, false)
	move, _ := minimax(root, 3)
	// fmt.Println(v, "\t", move.string(b))
	return move
}

func randomMove(b board, c color) move {
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

func opposite(c color) color {
	if c == white {
		return black
	}
	return white
}

func (n *node) evaluate() {
	n.value = 0
	minDistanceToKing := MaxInt
	var kingCoord *coordinate = nil
	for _, p := range n.board {
		if p.color != n.player && p.rank == king {
			kingCoord = &p.coord
			break
		}
	}
	if kingCoord == nil {
		panic("no king!")
	}
	for _, p := range n.board {
		if p.color != n.player && p.stolen {
			n.value += (int(p.rank) ^ 2)
		}
		if p.color == n.player {
			distance := abs(p.coord.col-kingCoord.col) + abs(p.coord.row-kingCoord.row)
			if distance < minDistanceToKing {
				minDistanceToKing = distance
			}
		}
	}
	n.value -= (minDistanceToKing * 2)
}

func (n *node) makeChildren() {
	moves := n.board.moves(n.player)
	n.isLeaf = len(moves) == 0
	for _, m := range moves {
		n.children[m] = newNode(&n.board, &m, opposite(n.player), !n.isOpponent)
	}
}

func newNode(b *board, m *move, c color, opponent bool) node {
	b1 := b.copy()
	if m != nil {
		b1 = b1.doMove(*m)
	}

	n := node{
		board:      b1,
		isOpponent: opponent,
		player:     c,
		children:   make(map[move]node),
		isLeaf:     false,
		value:      0,
		edge:       m,
		parent:     b,
	}

	return n
}

func minimax(n node, depth int) (move, int) {
	n.makeChildren()

	if n.isLeaf || depth == 0 {
		n.evaluate()
		return move{}, n.value
	}
	if !n.isOpponent {
		v := MinInt
		var m *move = nil
		for m1, child := range n.children {
			_, v1 := minimax(child, depth-1)
			if v1 > v {
				v = v1
				m = &m1
			}
		}
		if m == nil {
			panic("no move")
		}
		return *m, v
	} else {
		v := MaxInt
		var m *move = nil
		for m1, child := range n.children {
			_, v1 := minimax(child, depth-1)
			if v1 < v {
				v = v1
				m = &m1
			}
		}
		if m == nil {
			panic("no move")
		}
		return *m, v
	}
}
