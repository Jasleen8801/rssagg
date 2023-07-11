package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Hello, world!")

	godotenv.Load(".env")

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT environment variable not set")
	}
	fmt.Println("PORT:", portString) // $env:PORT = "8000"

}

// go get github.com/joho/godotenv
// go mod vendor
// go mod tidy
// go mod vendor
