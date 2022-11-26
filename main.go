package main

import (
	"log"
	"os"

	"github.com/FlamesX-128/aridia/src/routes"
)

var (
	dirPath string
	rPort   string
	dUri    string
	err     error
)

func main() {
	if dUri = os.Getenv("DATABASE_URL"); dUri == "" {
		dUri = "localhost:28015"
	}

	if rPort = os.Getenv("PORT"); rPort == "" {
		rPort = "8000"
	}

	if dirPath, err = os.Getwd(); err != nil {
		log.Panicln(err.Error())
	}

	if err = routes.SetupRoutes(dirPath, rPort, dUri); err != nil {
		log.Panicln(err.Error())
	}

}
