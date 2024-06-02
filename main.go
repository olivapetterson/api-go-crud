package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv" // lib para ler .env
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	fmt.Println(os.Getenv("TEST"))
}
