package main

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

type Game struct {
	Board Board
	Turn  int
}

type Node struct {
	Board      Board
	Edge       *Move
	Children   []Node
	IsOpponent bool
	Player     Color
	IsLeaf     bool
	Value      int
}
