package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthCheck(t *testing.T) {
	handler := HealthCheckHandler{}

	res := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/v1/healthCheck", nil)

	handler.ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("wrong status code: got %v want %v", res.Code, http.StatusOK)
	}
}
