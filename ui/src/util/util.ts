import { Color, Rank } from './consts';
import { Coordinate, Move, Piece } from './types';

export const makeBoard = (): Coordinate[] => {
  const coords: Coordinate[] = []
  for (let row = 0; row < 8; row++) {
    for (let col = 0; col < 8; col++) {
      coords.push({row, col})
    } 
  }
  return coords
}

export const getPieceAtCoord = (b: Piece[], c: Coordinate): number => {
  return b.findIndex(p => !p.stolen && p.coord.row === c.row && p.coord.col === c.col)
}

export const stringForPiece = (p: Piece): string => {
  switch (p.color) {
    case Color.White:
      switch (p.rank) {
        case Rank.Pawn:
          return '♙'
        case Rank.Rook:
          return '♖'
        case Rank.Knight:
          return '♘'
        case Rank.Bishop:
          return '♗'
        case Rank.King:
          return '♔'
        case Rank.Queen:
          return '♕'
        default:
          return ''
      }
    case Color.Black:
      switch (p.rank) {
        case Rank.Pawn:
          return '♟︎'
        case Rank.Rook:
          return '♜'
        case Rank.Knight:
          return '♞'
        case Rank.Bishop:
          return '♝'
        case Rank.King:
          return '♚'
        case Rank.Queen:
          return '♛'
        default:
          return ''
      }
    default:
      return ''
    }
}

export const stringForMove = (b: Piece[], m: Move): string => {
  let str = `${stringForPiece(b[m.mover])} from ${stringForCoord(b[m.mover].coord)} to ${stringForCoord(m.coord)}`
  if (m.steal >= 0) {
    str += ` (takes ${stringForPiece(b[m.steal])})`
  }
  return str
}

export const stringForCoord = (c: Coordinate): string => {
  return `${c.col},${c.row}`
}
