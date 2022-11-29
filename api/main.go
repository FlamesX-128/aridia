package main

import (
	"fmt"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error on loading .env file.")
	}

}

func main() {

}
