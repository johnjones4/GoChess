package clientserver

import (
	"main/chess"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"

	"golang.org/x/net/websocket"
)

func jsonrpcHandler(ws *websocket.Conn) {
	jsonrpc.ServeConn(ws)
}

func StartServer(host string) error {
	sm := SessionManager{
		Games: make(map[string]chess.Game),
	}
	rpc.Register(&sm)
	http.Handle("/jsonrpc", websocket.Handler(jsonrpcHandler))
	return http.ListenAndServe(host, nil)
}
