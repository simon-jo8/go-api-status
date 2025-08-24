package main

import (
	"log"
	"net/http"

	"github.com/simonjoseph/go-status-api/internal/api"
)

func main() {
	router := api.NewRouter()

	log.Println("Server is running on http://localhost:8080")
	log.Println("Available routes:")
	log.Println("- GET /status")
	log.Println("- POST /nextYear")

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}
