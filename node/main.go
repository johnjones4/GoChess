package main

import (
	"os"

	"github.com/johnjones4/GoChess/chess/distributed"
)

func main() {
	distributed.StartNodeServer(os.Getenv("HOST"))
}
