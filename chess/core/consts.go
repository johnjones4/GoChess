package core

const (
	boardSize = 8
)

const (
	White Color = 1
	Black Color = 2
)

const (
	Pawn   Rank = 1
	Rook   Rank = 2
	Knight Rank = 3
	Bishop Rank = 4
	Queen  Rank = 5
	King   Rank = 10
)

var lineUp = []Rank{
	Rook,
	Knight,
	Bishop,
	Queen,
	King,
	Bishop,
	Knight,
	Rook,
}
