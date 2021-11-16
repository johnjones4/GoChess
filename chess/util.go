package chess

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

func opposite(c Color) Color {
	if c == White {
		return Black
	}
	return White
}

func coordsEqual(a, b Coordinate) bool {
	return a.Row == b.Row && a.Col == b.Col
}
