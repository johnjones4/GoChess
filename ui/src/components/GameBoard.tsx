import React, {Component} from 'react'
import { Color, Method } from '../util/consts'
import JSONRPC from '../util/JSONRPC'
import { Game, GameArgs, GameResponse, GetMoveArgs, GetMoveResponse, Move, MoveArgs } from '../util/types'
import Board from './Board'

interface GameProps {

}

interface GameState {
  game: Game | null
  session: string | null
}

export default class GameBoard extends Component<GameProps, GameState> {
  private rpc: JSONRPC

  constructor(props: GameProps) {
    super(props)
    this.rpc = new JSONRPC('ws://'+window.location.host+'/jsonrpc')
    this.state = {
      game: null,
      session: null
    }
  }

  componentDidMount() {
    this.rpc.ws.onopen = async () => {
      try {
        const response = await this.rpc.call<GameArgs, GameResponse>(Method.NewGame, {color: Color.White})
        this.setState({
          game: response.game,
          session: response.session
        })
      } catch (e) {
        console.error(e)
      }
    }
  }

  render() {
    return this.state.game && (
      <Board
        game={this.state.game}
        onMove={ async (move: Move) => {
          const game = await this.rpc.call<MoveArgs, Game>(Method.Move, {session: this.state.session as string, move})
          this.setState({game})
        }}
        onSelection={ async (piece: number) : Promise<Move[]> => {
          const resp = await this.rpc.call<GetMoveArgs, GetMoveResponse>(Method.GetMoves, {session: this.state.session as string, piece})
          return resp.moves
        }}
      />
    )
  }
}
