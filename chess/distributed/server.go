package distributed

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"time"

	"github.com/johnjones4/GoChess/chess/ai"
)

type NodeServer struct {
}

type ExploreTreeArgs struct {
	Node ai.Node `json:"node"`
}

type ExploreTreeResponse struct {
	Value float64 `json:"value"`
}

func (n *NodeServer) ExploreTree(args ExploreTreeArgs, resp *ExploreTreeResponse) error {
	log.Println("Received tree traversal requests")
	start := time.Now()
	node, _ := ai.Minimax(&args.Node, -1.0, 1.0, 5)
	total := float64(time.Now().UnixNano()-start.UnixNano()) / 1000000000.0
	log.Printf("Explored tree in %fs\n", total)
	resp.Value = node.Value
	return nil
}

func StartNodeServer(host string) error {
	rpc.Register(new(NodeServer))
	rpc.HandleHTTP()
	log.Printf("Starting node server on %s", host)
	l, err := net.Listen("tcp", host)
	if err != nil {
		return err
	}
	return http.Serve(l, nil)
}
