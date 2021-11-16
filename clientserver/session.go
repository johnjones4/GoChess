package clientserver

import (
	"errors"
	"main/chess"

	"github.com/google/uuid"
)

type SessionManager struct {
	Games map[string]chess.Game
}

type GameArgs struct {
	Color chess.Color `json:"color"`
}

type GameResponse struct {
	Session string     `json:"session"`
	Game    chess.Game `json:"game"`
}

func (s *SessionManager) NewGame(args GameArgs, res *GameResponse) error {
	res.Game = chess.NewGame(args.Color)
	if args.Color == chess.Black {
		res.Game.TakeComputerTurn()
	}
	res.Session = uuid.NewString()
	s.Games[res.Session] = res.Game
	return nil
}

type MoveArgs struct {
	Session string     `json:"session"`
	Move    chess.Move `json:"move"`
}

func (s *SessionManager) Move(args MoveArgs, game *chess.Game) error {
	if g, ok := s.Games[args.Session]; ok {
		*game = g
		//TODO validate move
		if !game.TakeTurn(args.Move) {
			game.TakeComputerTurn()
		}
		s.Games[args.Session] = *game
	} else {
		return errors.New("no session with that ID")
	}
	return nil
}

type GetMoveArgs struct {
	Session string `json:"session"`
	Piece   int    `json:"piece"`
}

type GetMoveResponse struct {
	Moves []chess.Move `json:"moves"`
}

func (s *SessionManager) GetMoves(args GetMoveArgs, resp *GetMoveResponse) error {
	if g, ok := s.Games[args.Session]; ok {
		moves, err := g.Board.MovesForPiece(args.Piece)
		if err != nil {
			return err
		}
		resp.Moves = moves
	} else {
		return errors.New("no session with that ID")
	}
	return nil
}
