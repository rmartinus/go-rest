package api

import (
	"net/http"
)

// Handlers return a map of api handlers.
func Handlers() map[string]http.Handler {
	return map[string]http.Handler{
		"getEmployee": getEmployee{},
	}
}
