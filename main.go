package main

import (
	"log"
	"os"

	"github.com/FlamesX-128/aridia/src/misc"
	"github.com/FlamesX-128/aridia/src/routes"
	"github.com/joho/godotenv"
)

var (
	dirPath string
	rPort   string
	dUri    string
	err     error
)

func init() {
	if godotenv.Load(); err != nil {
		log.Panicln("Error loading .env file")
	}

	if misc.AuthConfig.ClientSecret = os.Getenv("CLIENT_SECRET"); misc.AuthConfig.ClientSecret == "" {
		log.Panicln("CLIENT_SECRET is not set")
	}

	if misc.AuthConfig.ClientID = os.Getenv("CLIENT_ID"); misc.AuthConfig.ClientID == "" {
		log.Panicln("CLIENT_ID is not set")
	}

	if dUri = os.Getenv("DATABASE_URL"); dUri == "" {
		dUri = "localhost:28015"
	}

	if rPort = os.Getenv("PORT"); rPort == "" {
		rPort = "8000"
	}

	if dirPath, err = os.Getwd(); err != nil {
		log.Panicln(err.Error())
	}

}

func main() {
	if err = routes.SetupRoutes(dirPath, rPort, dUri); err != nil {
		log.Panicln(err.Error())
	}

}
