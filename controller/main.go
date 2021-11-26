package main

import (
	"log"
	"os"

	"github.com/johnjones4/GoChess/chess/core"
	"github.com/johnjones4/GoChess/chess/distributed"
	"github.com/johnjones4/GoChess/chess/uiserver"
)

func main() {
	err := distributed.InitNodes()
	if err != nil {
		log.Fatal(err)
	}
	err = core.InitGameStorage()
	if err != nil {
		log.Fatal(err)
	}
	uiserver.StartUIServer(os.Getenv("HOST"))
}
