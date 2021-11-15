package main

type color int

type rank int

type piece struct {
	color  color
	rank   rank
	coord  coordinate
	stolen bool
}

type coordinate struct {
	row int
	col int
}

type move struct {
	mover int
	coord coordinate
	steal int
}

type delta struct {
	dr int
	dc int
}

type board []piece

type game struct {
	board board
	turn  int
}

type node struct {
	board      board
	children   map[move]node
	isOpponent bool
	player     color
	isLeaf     bool
	value      int
	edge       *move
	parent     *board
}
