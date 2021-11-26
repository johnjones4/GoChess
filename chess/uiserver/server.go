package uiserver

import (
	"errors"
	"log"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"

	"github.com/johnjones4/GoChess/chess/core"
	"github.com/johnjones4/GoChess/chess/distributed"

	"github.com/google/uuid"
	"golang.org/x/net/websocket"
)

type SessionManager struct {
	Games map[string]core.Game
}

type GameArgs struct {
	Color core.Color `json:"color"`
}

type GameResponse struct {
	Session string    `json:"session"`
	Game    core.Game `json:"game"`
}

func (s *SessionManager) NewGame(args GameArgs, res *GameResponse) error {
	game, err := core.NewGame(args.Color, distributed.MinimaxRemote)
	if err != nil {
		return err
	}
	res.Game = game
	if args.Color == core.Black {
		res.Game.TakeComputerTurn()
	}
	res.Session = uuid.NewString()
	s.Games[res.Session] = res.Game
	log.Printf("Started game with session ID %s\n", res.Session)
	return nil
}

type MoveArgs struct {
	Session string    `json:"session"`
	Move    core.Move `json:"move"`
}

func (s *SessionManager) Move(args MoveArgs, game *core.Game) error {
	if g, ok := s.Games[args.Session]; ok {
		*game = g
		if !game.Board.MoveIsValid(game.UserPlayer, args.Move) {
			return errors.New("move is illegal")
		}
		winner, err := game.TakeTurn(args.Move)
		if err != nil {
			return err
		}
		if !winner {
			game.TakeComputerTurn()
		}
		s.Games[args.Session] = *game

		log.Printf("Game (%s): %s\n", args.Session, args.Move.String(game.Board))
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
	Moves []core.Move `json:"moves"`
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

func jsonrpcHandler(ws *websocket.Conn) {
	jsonrpc.ServeConn(ws)
}

func StartUIServer(host string) error {
	sm := SessionManager{
		Games: make(map[string]core.Game),
	}
	rpc.Register(&sm)
	http.Handle("/jsonrpc", websocket.Handler(jsonrpcHandler))
	return http.ListenAndServe(host, nil)
}
