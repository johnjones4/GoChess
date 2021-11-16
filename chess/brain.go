package chess

func bestMinimaxMove(b Board, c Color) Move {
	root := newNode(&b, nil, 0, c, false)
	move, _ := minimax(root, minInt, maxInt, 3)
	return move
}
