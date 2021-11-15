package main

const (
	boardSize = 8
)

const (
	white Color = 1
	black Color = 2
)

const (
	pawn   Rank = 1
	rook   Rank = 2
	knight Rank = 3
	bishop Rank = 4
	queen  Rank = 5
	king   Rank = 6
)

var lineUp = []Rank{
	rook,
	knight,
	bishop,
	queen,
	king,
	bishop,
	knight,
	rook,
}
