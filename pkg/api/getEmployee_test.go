package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	"github.com/mnf-group/openapimux"

	"testing"
)

// HTTPTest defines struct for HTTP test.
type HTTPTest struct {
	Name string
	Req  *http.Request
	Res  HTTPResult
}

// HTTPResult defines standard HTTP Response.
type HTTPResult struct {
	Body APIResponse
	Code int
}

// APIResponse defines standard JSON API Response.
type APIResponse struct {
	Success bool                   `json:"success"`
	Data    map[string]interface{} `json:"data"`
	Error   map[string]interface{} `json:"error"`
}

func TestGetEmployee(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	tests := []HTTPTest{
		HTTPTest{
			Name: "Get employee successfully",
			Req: openapimux.WithPathParams(req, map[string]string{
				"id": "1",
			}),
			Res: HTTPResult{
				Code: http.StatusOK,
				Body: APIResponse{
					Success: true,
					Data: map[string]interface{}{
						"id": "1",
					},
				},
			},
		},
		HTTPTest{
			Name: "Get employee error",
			Req: openapimux.WithPathParams(req, map[string]string{
				"id": "100",
			}),
			Res: HTTPResult{
				Code: http.StatusInternalServerError,
				Body: APIResponse{
					Success: false,
					Error: map[string]interface{}{
						"code":    500.0,
						"details": "DB error",
						"message": "Something went wrong.",
					},
				},
			},
		},
	}

	handler := getEmployee{}

	for _, tt := range tests {
		t.Logf("Running %s", tt.Name)

		res := httptest.NewRecorder()
		handler.ServeHTTP(res, tt.Req)
		resultPayload, _ := ioutil.ReadAll(res.Body)

		resultObj := APIResponse{}
		json.Unmarshal(resultPayload, &resultObj)

		if resultObj.Success != tt.Res.Body.Success {
			t.Errorf("Wrong success: got %v want %v", resultObj.Success, tt.Res.Body.Success)
		}

		var compareError error
		if tt.Res.Code < 300 {
			compareError = compareMaps(tt.Res.Body.Data, resultObj.Data)
		} else {
			compareError = compareMaps(tt.Res.Body.Error, resultObj.Error)
		}

		if compareError != nil {
			t.Errorf("Response body is not expected: %s", compareError)
		}

		if res.Code != tt.Res.Code {
			t.Errorf("Wrong status code: got %v want %v", res.Code, tt.Res.Code)
		}
	}
}

func compareMaps(a map[string]interface{}, b map[string]interface{}) error {
	for k, v1 := range a {
		v2, ok := b[k]
		if v1 == nil {
			if ok {
				return fmt.Errorf("key %s should not be present", k)
			}

			return nil
		}

		if !ok {
			return fmt.Errorf("No key %s", k)
		}

		if v1 != v2 {
			return fmt.Errorf("not equal - %+v and %+v", v1, v2)
		}
	}

	return nil
}
