package models

// Response represents a standard API response
type Response struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

// NextYearRequest represents the request body for the nextYear endpoint
type NextYearRequest struct {
	Year int `json:"year"`
}

// NextYearResponse represents the response body for the nextYear endpoint
type NextYearResponse struct {
	Year int `json:"year"`
}
