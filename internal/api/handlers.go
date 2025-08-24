package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/simonjoseph/go-status-api/internal/models"
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

func (router *Router) handleNextYear(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		router.handleMethodNotAllowed(w)
		return
	}

	// TODO: how to handle validation at this level ?
	request := &models.NextYearRequest{}
	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
		router.handleBadRequest(w, "Invalid request")
		return
	}

	nextYearResponse := models.NextYearResponse{
		Year: request.Year + 1,
	}
	response := models.Response{
		Status: "success",
		Data:   nextYearResponse,
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
