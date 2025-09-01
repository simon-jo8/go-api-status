package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/simonjoseph/go-status-api/internal"
	"github.com/simonjoseph/go-status-api/models"
)

func (router *Router) handleStatus(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		router.handleMethodNotAllowed(w)
		return
	}

	response := models.Response{
		Status: "success",
		Data: map[string]string{
			"date": time.Now().Format(time.RFC3339),
		},
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (router *Router) handlePlusOne(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		router.handleMethodNotAllowed(w)
		return
	}

	request := &models.PlusOneRequest{}
	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
		router.handleBadRequest(w, "Invalid request")
		return
	}

	plusOneResponse := models.PlusOneResponse{
		Number: internal.PlusOne(request.Number),
	}
	response := models.Response{
		Status: "success",
		Data:   plusOneResponse,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (router *Router) handleGoldenHour(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		router.handleMethodNotAllowed(w)
		return
	}

	request := &models.GoldenHourRequest{}
	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
		router.handleBadRequest(w, "Invalid request")
		return
	}

	goldenHourResponse := models.GoldenHourResponse{
		GoldenHour: internal.GoldenHour(request.Latitude, request.Longitude, request.Date),
	}
	response := models.Response{
		Status: "success",
		Data:   goldenHourResponse,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (router *Router) handleNotFound(w http.ResponseWriter) {
	response := models.Response{
		Status: "error",
		Data: map[string]string{
			"message": "404 not found",
		},
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(response)
}

func (router *Router) handleMethodNotAllowed(w http.ResponseWriter) {
	response := models.Response{
		Status: "error",
		Data: map[string]string{
			"message": "Method not allowed",
		},
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
	json.NewEncoder(w).Encode(response)
}

func (router *Router) handleBadRequest(w http.ResponseWriter, message string) {
	response := models.Response{
		Status: "error",
		Data: map[string]string{
			"message": message,
		},
	}
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(response)
}
