package api

import (
	"errors"
	"net/http"

	"github.com/mnf-group/openapimux"
	"github.com/rmartinus/go-rest/pkg/logger"
)

type getEmployee struct {
}

func (e getEmployee) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := openapimux.PathParam(r, "id")
	logger.FromContext(ctx).Infof("ID %s", id)

	if id == "1" {
		Success(w, map[string]interface{}{"id": id}, http.StatusOK)
		return
	}

	err := errors.New("DB error")
	logger.FromContext(ctx).Errorf("Error: %+v", err)
	Error(w, r, AppError{500, "Something went wrong.", err})
}
