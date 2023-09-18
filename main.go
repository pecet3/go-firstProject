package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	port := os.Getenv("PORT")

	fmt.Println(port)

	if port == "" {
		log.Fatal("Port is not found in .env")

	}

	router := chi.NewRouter()

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	log.Printf("Server is Listening on PORT %v", port)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
