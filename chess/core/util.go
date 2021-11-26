package core

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func Opposite(c Color) Color {
	if c == White {
		return Black
	}
	return White
}

func coordsEqual(a, b Coordinate) bool {
	return a.Row == b.Row && a.Col == b.Col
}

func Distance(a, b Coordinate) int {
	return Abs(a.Row-b.Row) + Abs(a.Col-b.Col)
}
