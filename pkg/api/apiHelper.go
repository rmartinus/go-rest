package api

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/mnf-group/openapimux"
)

// AppError struct for handled error.
type AppError struct {
	Code    int
	Message string
	Err     error
}

// NewRouter returns a router.
func NewRouter(handlers map[string]http.Handler, apis ...string) (*openapimux.OpenAPIMux, error) {
	if handlers == nil {
		handlers = make(map[string]http.Handler)
	}

	routers := make(openapi3filter.Routers, len(apis))
	handlers["healthCheck"] = HealthCheckHandler{}

	openapi3.DefineStringFormat("uuid", openapi3.FormatOfStringForUUIDOfRFC4122)
	loader := openapi3.NewSwaggerLoader()
	loader.IsExternalRefsAllowed = true

	for i, path := range apis {
		data, err := ioutil.ReadFile(path)
		if err != nil {
			return nil, err
		}

		swagger, err := loader.LoadSwaggerFromDataWithPath(data, &url.URL{
			Path: path,
		})

		if err != nil {
			return nil, err
		}

		routers[i] = openapi3filter.NewRouter().WithSwagger(swagger)
	}

	r := &openapimux.OpenAPIMux{
		Routers:       &routers,
		ErrorHandler:  HandleError,
		DetailedError: true,
	}

	r.UseHandlers(handlers)

	return r, nil
}

// HandleError returns default error handler.
func HandleError(w http.ResponseWriter, r *http.Request, e string, code int) {
	Error(w, r, AppError{
		Code:    code,
		Message: "Swagger router: OpenAPI request failed",
		Err:     errors.New(e),
	})
}

// Success returns successful response.
func Success(w http.ResponseWriter, data interface{}, code int) {
	dataMap := map[string]interface{}{"success": true}

	if data != nil {
		dataMap["data"] = data
	}
	byteMap, _ := json.Marshal(dataMap)

	w.WriteHeader(code)
	w.Write(byteMap)
}

// Error returns error response.
func Error(w http.ResponseWriter, r *http.Request, e AppError) {
	code := e.Code
	if code == 0 {
		code = http.StatusInternalServerError
	}

	var details string
	if e.Err != nil {
		details = e.Err.Error()
	}

	dataMap := map[string]interface{}{
		"success": false,
		"error": map[string]interface{}{
			"code":    code,
			"message": e.Message,
			"details": details,
		},
	}

	byteMap, _ := json.Marshal(dataMap)

	w.WriteHeader(code)
	w.Write(byteMap)
}
