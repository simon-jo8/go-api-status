package api

import (
	"net/http"
)

// Router handles HTTP requests
type Router struct{}

// NewRouter creates a new Router instance
func NewRouter() *Router {
	return &Router{}
}

// ServeHTTP implements the http.Handler interface
func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.URL.Path {
	case "/status":
		router.handleStatus(w, r)
	case "/nextYear":
		router.handleNextYear(w, r)
	default:
		router.handleNotFound(w, r)
	}
}
