package ai

import "github.com/johnjones4/GoChess/chess/core"

func (n *Node) pcntPieceOnBoard() float64 {
	piecesOnBoard := 0.0
	piecesTotal := 0.0
	for _, p := range n.Board {
		if p.Color == n.Player {
			if !p.Stolen {
				piecesOnBoard += 1.0
			}
			piecesTotal += 1.0
		}
	}
	return piecesOnBoard / piecesTotal
}

func (n *Node) avgDistanceToKing() float64 {
	var kingCoord *core.Coordinate
	for _, p := range n.Board {
		if p.Color != n.Player && p.Rank == core.King {
			kingCoord = &p.Coord
			break
		}
	}
	if kingCoord == nil {
		panic("no coordinate for king")
	}
	distances := 0.0
	nDistances := 0.0
	for _, p := range n.Board {
		if p.Color == n.Player {
			distances += float64(core.Distance(p.Coord, *kingCoord))
			nDistances++
		}
	}
	return distances / nDistances
}
