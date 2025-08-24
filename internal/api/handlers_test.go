package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/simonjoseph/go-status-api/internal/models"
)

func TestRouter(t *testing.T) {
	router := NewRouter()

	tests := []struct {
		name           string
		method         string
		path           string
		body           interface{}
		expectedCode   int
		expectedStatus string
		checkResponse  func(t *testing.T, response *models.Response)
	}{
		{
			name:           "GET /status returns success with date",
			method:         http.MethodGet,
			path:           "/status",
			expectedCode:   http.StatusOK,
			expectedStatus: "success",
			checkResponse: func(t *testing.T, response *models.Response) {
				dataMap, ok := response.Data.(map[string]interface{})
				if !ok {
					t.Errorf("data is not a map[string]interface{}, got %T", response.Data)
					return
				}
				dateStr, ok := dataMap["date"].(string)
				if !ok {
					t.Errorf("date is not a string, got %T", dataMap["date"])
					return
				}
				_, err := time.Parse(time.RFC3339, dateStr)
				if err != nil {
					t.Errorf("date is not a valid RFC3339 date: %v", err)
				}
			},
		},
		{
			name:           "POST /status returns method not allowed",
			method:         http.MethodPost,
			path:           "/status",
			expectedCode:   http.StatusMethodNotAllowed,
			expectedStatus: "error",
			checkResponse: func(t *testing.T, response *models.Response) {
				checkErrorMessage(t, response, "Method not allowed")
			},
		},
		{
			name:           "POST /nextYear returns next year",
			method:         http.MethodPost,
			path:           "/nextYear",
			body:           models.NextYearRequest{Year: 2024},
			expectedCode:   http.StatusOK,
			expectedStatus: "success",
			checkResponse: func(t *testing.T, response *models.Response) {
				var nextYearResponse models.NextYearResponse
				responseData, err := json.Marshal(response.Data)
				if err != nil {
					t.Errorf("failed to marshal response data: %v", err)
					return
				}
				if err := json.Unmarshal(responseData, &nextYearResponse); err != nil {
					t.Errorf("failed to unmarshal response data: %v", err)
					return
				}
				if nextYearResponse.Year != 2025 {
					t.Errorf("expected year 2025, got %d", nextYearResponse.Year)
				}
			},
		},
		{
			name:           "GET /nextYear returns method not allowed",
			method:         http.MethodGet,
			path:           "/nextYear",
			expectedCode:   http.StatusMethodNotAllowed,
			expectedStatus: "error",
			checkResponse: func(t *testing.T, response *models.Response) {
				checkErrorMessage(t, response, "Method not allowed")
			},
		},
		{
			name:           "POST /nextYear with invalid JSON returns bad request",
			method:         http.MethodPost,
			path:           "/nextYear",
			body:           "invalid json",
			expectedCode:   http.StatusBadRequest,
			expectedStatus: "error",
			checkResponse: func(t *testing.T, response *models.Response) {
				checkErrorMessage(t, response, "Invalid request")
			},
		},
		{
			name:           "GET /invalid returns 404",
			method:         http.MethodGet,
			path:           "/invalid",
			expectedCode:   http.StatusNotFound,
			expectedStatus: "error",
			checkResponse: func(t *testing.T, response *models.Response) {
				checkErrorMessage(t, response, "404 not found")
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var body bytes.Buffer
			if test.body != nil {
				if jsonStr, ok := test.body.(string); ok {
					body.WriteString(jsonStr)
				} else {
					if err := json.NewEncoder(&body).Encode(test.body); err != nil {
						t.Fatalf("Failed to encode request body: %v", err)
					}
				}
			}

			req := httptest.NewRequest(test.method, test.path, &body)
			rec := httptest.NewRecorder()

			router.ServeHTTP(rec, req)

			// Check status code
			if rec.Code != test.expectedCode {
				t.Errorf("expected status code %d, got %d", test.expectedCode, rec.Code)
			}

			// Check Content-Type
			contentType := rec.Header().Get("Content-Type")
			if contentType != "application/json" {
				t.Errorf("expected Content-Type application/json, got %s", contentType)
			}

			// Parse response
			var response models.Response
			if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
				t.Fatalf("Failed to decode response body: %v", err)
			}

			// Check response status
			if response.Status != test.expectedStatus {
				t.Errorf("expected status %s, got %s", test.expectedStatus, response.Status)
			}

			// Run custom response checks
			test.checkResponse(t, &response)
		})
	}
}

func checkErrorMessage(t *testing.T, response *models.Response, expectedMessage string) {
	dataMap, ok := response.Data.(map[string]interface{})
	if !ok {
		t.Errorf("data is not a map[string]interface{}, got %T", response.Data)
		return
	}
	message, ok := dataMap["message"].(string)
	if !ok {
		t.Errorf("message is not a string, got %T", dataMap["message"])
		return
	}
	if message != expectedMessage {
		t.Errorf("expected message %q, got %q", expectedMessage, message)
	}
}
