package api

import (
	"testing"
)

func TestAPI(t *testing.T) {
	handlers := Handlers()
	if handlers == nil {
		t.Errorf("handlers must not be nil")
	}

	if len(handlers) == 0 {
		t.Errorf("handlers must not be empty")
	}
}
