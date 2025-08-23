package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type Response struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

type NextYearRequest struct {
	Year int `json:"year"`
}

type NextYearResponse struct {
	Year int `json:"year"`
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	if(r.Method != "GET") {
		response := Response{
			Status: "error",
			Data: map[string]string{
				"message": "Method not allowed",
			},
		}
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(response)
		return
	}
	response := Response{
		Status: "success",
		Data: map[string]string{
			"date": time.Now().Format(time.RFC3339),
		},
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func nextYearHandler(w http.ResponseWriter, r *http.Request) {
	if(r.Method != "POST") {
		response := Response{
			Status: "error",
			Data: map[string]string{
				"message": "Method not allowed",
			},
		}
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(response)
		return
	}
	// TODO: how to validate the request?
	request := &NextYearRequest{}
	err := json.NewDecoder(r.Body).Decode(request) 
	if err != nil {
		response := Response{
			Status: "error",
			Data: map[string]string{
				"message": "Invalid request",
			},
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	response := NextYearResponse{
		Year: request.Year + 1,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Status: "error",
		Data: map[string]string{
			"message": "404 not found",
		},
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(response)
}

func main() {
	// List of routes
	http.HandleFunc("/status", statusHandler)
	http.HandleFunc("/nextYear", nextYearHandler)
	http.HandleFunc("/", notFoundHandler)

	println("Server is running on http://localhost:8080")
	println("Available routes:")
	println("- GET /status")
	println("- GET /nextYear")
	
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}