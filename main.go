package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	port := os.Getenv("PORT")

	fmt.Println(port)

	if port == "" {
		log.Fatal("Port is not found in .env")

	}
}
