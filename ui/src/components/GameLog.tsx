import React from 'react'
import { Game } from '../util/types'
import { stringForMove } from '../util/util'
import './GameLog.css'

interface GameLogProps {
  game: Game
}

const GameLog = (props: GameLogProps) => {
  return (
    <div className='GameLog'>
      <ul>
        { props.game.log.map((logItem, i) => (<li key={i}>{stringForMove(logItem.board, logItem.move)}</li>)) }
      </ul>
    </div>
  )
}

export default GameLog
