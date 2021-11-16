import React, {Component} from 'react'
import { Coordinate, Game, Move } from '../util/types'
import { getPieceAtCoord, makeBoard, stringForPiece } from '../util/util'
import './Board.css'

interface BoardProps {
  game: Game
  onMove(m: Move): void
  onSelection(index: number): Promise<Move[]>
}

interface Selection {
  coord: Coordinate
  index: number
  validMoves: Move[] | null
}

interface BoardState {
  selected: Selection | null
}

export default class Board extends Component<BoardProps, BoardState>{
  constructor(props: BoardProps) {
    super(props)
    this.state = {
      selected: null
    }
  }

  async selectTile(coord: Coordinate) {
    const pieceIndex = getPieceAtCoord(this.props.game.board, coord)
    const piece = pieceIndex >= 0 ? this.props.game.board[pieceIndex] : null
    if (this.state.selected) {
      if (piece && this.state.selected.coord.row === piece.coord.row && this.state.selected.coord.col === piece.coord.col) {
        this.setState({
          selected: null
        })
      } else if (!piece || piece.color !== this.props.game.board[this.state.selected.index].color) {
        this.props.onMove({
          mover: this.state.selected.index,
          coord,
          steal: piece ? pieceIndex : -1
        })
        this.setState({
          selected: null
        })
      }
    } else if (piece && piece.color === this.props.game.userPlayer) {
      this.setState({
        selected: {
          index: pieceIndex,
          coord,
          validMoves: await this.props.onSelection(pieceIndex)
        }
      })
    }
  }

  isValidMoveForSelection(coord: Coordinate): boolean {
    if (this.state.selected === null || this.state.selected.validMoves === null) {
      return false
    }
    return this.state.selected.validMoves.findIndex(move => {
      return move.coord.row === coord.row && move.coord.col === coord.col
    }) >= 0
  }

  render() {
    return (
      <div className='Board'>
        { makeBoard().map((coord, i) => {
          const pieceIndex = getPieceAtCoord(this.props.game.board, coord)
          const piece = pieceIndex >= 0 ? this.props.game.board[pieceIndex] : null
          return (
            <button 
              key={i}
              grid-column-start={coord.col}
              grid-column-end={coord.col}
              grid-row-start={coord.row}
              grid-row-end={coord.row}
              disabled={this.props.game.winner >= 0 || (this.state.selected === null && piece?.color !== this.props.game.userPlayer) || (this.state.selected !== null && !this.isValidMoveForSelection(coord))}
              onClick={() => this.selectTile(coord)}
            >
              { piece && stringForPiece(piece) }
            </button>
          )
        }) }
      </div>
    )
  }
}
