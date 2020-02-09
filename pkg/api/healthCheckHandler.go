package api

import (
	"net/http"
)

// HealthCheckHandler standard handler for healthCheck endpoint - outputs OK.
type HealthCheckHandler struct{}

func (h HealthCheckHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	Success(w, nil, http.StatusOK)
}
