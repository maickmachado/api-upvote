package test

import (
	"github.com/maickmachado/upvote-api/controllers"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthCheck(t *testing.T) {
	tt := []struct {
		name   string
		method string
		//input      *Pizzas
		//want       string
		statusCode int
	}{
		{
			name:   "status ok",
			method: http.MethodGet,
			//input:      &Pizzas{},
			//want:       "Error: No pizzas found",
			statusCode: http.StatusOK,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			request := httptest.NewRequest(tc.method, "/healthcheck", nil)
			responseRecorder := httptest.NewRecorder()

			controllers.HealthCheck(responseRecorder, request)

			if responseRecorder.Code != tc.statusCode {
				t.Errorf("Want status '%d', got '%d'", tc.statusCode, responseRecorder.Code)
			}
		})
	}
}
