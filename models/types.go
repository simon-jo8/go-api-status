package models

// Response represents a standard API response
type Response struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

// PlusOneRequest represents the request body for the plusOne endpoint
type PlusOneRequest struct {
	Number int `json:"number"`
}

// PlusOneResponse represents the response body for the plusOne endpoint
type PlusOneResponse struct {
	Number int `json:"number"`
}

type GoldenHourRequest struct {
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Date string `json:"date,omitempty"`
}

type GoldenHourResponse struct {
	GoldenHour string `json:"goldenHour"`
}
