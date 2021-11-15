package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	wins := map[color]int{
		white: 0,
		black: 0,
	}
	rand.Seed(time.Now().Unix())
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		g := newGame()
		winner := g.run()
		wins[winner]++
	}
	for player, wins := range wins {
		fmt.Printf("%s: %d\n", player.string(), wins)
	}
}
