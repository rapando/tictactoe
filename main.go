package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/rapando/tictactoe/controllers"
)

func main() {
	var base controllers.Base
	var err error
	if err = godotenv.Load(); err != nil {
		log.Panic(err)
	}

	if err = base.Init(); err != nil {
		log.Panic(err)
	}
	base.RunServer()
}
