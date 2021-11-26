package main

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/johnjones4/GoChess/chess/ai"
	"github.com/johnjones4/GoChess/chess/core"
	"github.com/johnjones4/GoChess/chess/distributed"
)

func main() {
	gamesToPlay, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	err = distributed.InitNodes()
	if err != nil {
		log.Fatal(err)
	}
	winners := map[core.Color]int{
		core.White: 0,
		core.Black: 0,
	}
	start := time.Now()
	for i := 0; i < gamesToPlay; i++ {
		log.Printf("Starting game %d\n", i+1)
		g := core.NewGame(-1, func(b core.Board, c core.Color) core.Move {
			if c == core.White {
				return distributed.MinimaxRemote(b, c)
			}
			return ai.RandomMove(b, c)
		})
		winner := g.ComputerGame()
		winners[winner]++
		log.Printf("Finished game %d. Winner: %s\n", i+1, winner.String())
	}
	elapsed := time.Now().Unix() - start.Unix()
	log.Printf("Finished all games in %d seconds (%f per second).\n", elapsed, float64(elapsed)/float64(gamesToPlay))
	for color, wins := range winners {
		log.Printf("%s: %d\n", color.String(), wins)
	}
}
