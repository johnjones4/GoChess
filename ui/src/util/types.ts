export interface Piece {
  color: number
  rank: number
  coord: Coordinate
  stolen: boolean
}

export interface Coordinate {
  row: number
  col: number
}

export interface Move {
  mover: number
  coord: Coordinate
  steal: number
}

export interface LogItem {
  board: Piece[]
  move: Move
}

export interface Game {
  board: Piece[]
  turn: number
  userPlayer: number
  winner: number
  log: LogItem[]
}

export interface GameArgs {
  color: number
}

export interface GameResponse {
  session: string
  game: Game
}

export interface MoveArgs {
  session: string
  move: Move
}

export interface GetMoveArgs {
  session: string
  piece: number
}

export interface GetMoveResponse {
  moves: Move[]
}
