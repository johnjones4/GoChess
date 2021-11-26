package distributed

import (
	"log"
	"net/rpc"
	"os"
	"strings"
	"sync"

	"github.com/johnjones4/GoChess/chess/ai"

	"github.com/johnjones4/GoChess/chess/core"
)

type connection struct {
	host   string
	client *rpc.Client
	lock   sync.Mutex
}

var connections []connection
var connectionIndex = 0
var connectionLock = sync.Mutex{}

func InitNodes() error {
	nodeAddrs := strings.Split(os.Getenv("NODE_ADDRS"), ",")
	connections = make([]connection, len(nodeAddrs))
	for i, addr := range nodeAddrs {
		client, err := rpc.DialHTTP("tcp", addr)
		if err != nil {
			return err
		}
		connections[i] = connection{addr, client, sync.Mutex{}}
		log.Printf("Connected to %s\n", addr)
	}
	return nil
}

func nextConnection() *connection {
	connectionLock.Lock()
	c := &connections[connectionIndex]
	connectionIndex++
	if connectionIndex >= len(connections) {
		connectionIndex = 0
	}
	connectionLock.Unlock()
	return c
}

func MinimaxRemote(b core.Board, c core.Color) core.Move {
	root := ai.NewNode(nil, b, nil, 0, c, false)
	root.NextLevel()
	values := make([]float64, len(root.Children))
	waitGroup := sync.WaitGroup{}
	for i, child := range root.Children {
		go _minimaxRemote(&waitGroup, child, 6, &values[i])
	}
	waitGroup.Wait()
	v := -1.0
	var m *core.Move = nil
	for i, v1 := range values {
		if v1 > v {
			v = v1
			m = root.Children[i].Edge
		}
	}
	if m == nil {
		panic("no move")
	}
	return *m
}

func _minimaxRemote(wg *sync.WaitGroup, n ai.Node, maxDepth int, result *float64) {
	wg.Add(1)
	c := nextConnection()
	c.lock.Lock()
	n.Parent = nil
	args := ExploreTreeArgs{n}
	var resp ExploreTreeResponse
	log.Printf("Dispatching child to %s\n", c.host)
	err := c.client.Call("NodeServer.ExploreTree", args, &resp)
	log.Printf("Child has value: %f\n", resp.Value)
	if err != nil {
		log.Panic(err)
	} else {
		*result = resp.Value
	}
	c.lock.Unlock()
	wg.Done()
}
