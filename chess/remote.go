package chess

import "sync"

func startMinimaxRemote(b Board, c Color) Move {
	root := newNode(nil, b, nil, 0, c, false)
	root.nextLevel()
	values := make([]int, len(root.Children))
	waitGroup := sync.WaitGroup{}
	for i, child := range root.Children {
		go minimaxRemote(&waitGroup, child, 6, &values[i])
	}
	waitGroup.Wait()
	v := minInt
	var m *Move = nil
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

func minimaxRemote(wg *sync.WaitGroup, n Node, maxDepth int, result *int) {
	wg.Add(1)
	node, _ := minimaxLocal(&n, minInt, maxInt, maxDepth)
	*result = node.Value
	wg.Done()
}
