package chess

const (
	boardSize = 8
)

const (
	White Color = 1
	Black Color = 2
)

const (
	pawn   Rank = 1
	rook   Rank = 2
	knight Rank = 3
	bishop Rank = 4
	queen  Rank = 5
	king   Rank = 10
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

const (
	maxInt = 10000000
	minInt = -10000000
)
