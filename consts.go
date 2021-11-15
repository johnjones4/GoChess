package main

const (
	boardSize = 8
)

const (
	white color = 1
	black color = 2
)

const (
	pawn   rank = 1
	rook   rank = 2
	knight rank = 3
	bishop rank = 4
	queen  rank = 5
	king   rank = 6
)

var lineUp = []rank{
	rook,
	knight,
	bishop,
	queen,
	king,
	bishop,
	knight,
	rook,
}
