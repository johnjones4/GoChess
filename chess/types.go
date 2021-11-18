package chess

type Color int

type Rank int

type Piece struct {
	Color  Color      `json:"color"`
	Rank   Rank       `json:"rank"`
	Coord  Coordinate `json:"coord"`
	Stolen bool       `json:"stolen"`
}

type Coordinate struct {
	Row int `json:"row"`
	Col int `json:"col"`
}

type Move struct {
	Mover int        `json:"mover"`
	Coord Coordinate `json:"coord"`
	Steal int        `json:"steal"`
}

type delta struct {
	dr int
	dc int
}

type Board []Piece

type LogItem struct {
	Board []Piece `json:"board"`
	Move  Move    `json:"move"`
}

type Game struct {
	Board      Board     `json:"board"`
	Turn       int       `json:"turn"`
	UserPlayer Color     `json:"userPlayer"`
	Winner     Color     `json:"winner"`
	Log        []LogItem `json:"log"`
}

type Node struct {
	Depth      int
	Board      Board
	Edge       *Move
	Parent     *Node
	Children   []Node
	IsOpponent bool
	Player     Color
	IsLeaf     bool
	Value      int
}
