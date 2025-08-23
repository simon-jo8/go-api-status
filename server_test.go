package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestHandler(t *testing.T) {
	tests := []struct {
		name           string
		path           string
		expectedCode   int
		expectedStatus string
		checkData      func(t *testing.T, data interface{})
	}{
		{
			name:           "Status endpoint returns success with date",
			path:           "/status",
			expectedCode:   http.StatusOK,
			expectedStatus: "success",
			checkData: func(t *testing.T, data interface{}) {
				dataMap, ok := data.(map[string]interface{})
				if !ok {
					t.Errorf("data is not a map[string]interface{}, got %T", data)
					return
				}
				dateStr, ok := dataMap["date"].(string)
				if !ok {
					t.Errorf("date is not a string, got %T", dataMap["date"])
					return
				}
				_, err := time.Parse(time.RFC3339, dateStr)
				if err != nil {
					t.Errorf("data is not a valid RFC3339 date: %v", err)
				}
			},
		},
		{
			name:           "Non-status endpoint returns 404",
			path:           "/invalid",
			expectedCode:   http.StatusNotFound,
			expectedStatus: "error",
			checkData: func(t *testing.T, data interface{}) {
				dataMap, ok := data.(map[string]interface{})
				if !ok {
					t.Errorf("data is not a map[string]interface{}, got %T", data)
					return
				}
				msg, ok := dataMap["message"].(string)
				if !ok {
					t.Errorf("message is not a string, got %T", dataMap["message"])
					return
				}
				if msg != "404 not found" {
					t.Errorf("unexpected error message: got %v want '404 not found'", msg)
				}
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Create a request to pass to our handler
			req, err := http.NewRequest("GET", test.path, nil)
			if err != nil {
				t.Fatal(err)
			}

			// Create a ResponseRecorder to record the response
			responseRecorder := httptest.NewRecorder()
			handler := http.HandlerFunc(handler)

			// Call the handler with our request
			handler.ServeHTTP(responseRecorder, req)

			// Check the status code
			if status := responseRecorder.Code; status != test.expectedCode {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, test.expectedCode)
			}

			// Check the Content-Type header
			contentType := responseRecorder.Header().Get("Content-Type")
			if contentType != "application/json" {
				t.Errorf("handler returned wrong content type: got %v want application/json",
					contentType)
			}

			// Parse the JSON response
			var response Response
			if err := json.NewDecoder(responseRecorder.Body).Decode(&response); err != nil {
				t.Fatalf("Could not decode response body: %v", err)
			}

			// Check the response status
			if response.Status != test.expectedStatus {
				t.Errorf("handler returned wrong status: got %v want %v",
					response.Status, test.expectedStatus)
			}

			// Check the data using the provided check function
			test.checkData(t, response.Data)
		})
	}
}