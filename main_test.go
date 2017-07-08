package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type endpointT struct {
	*testing.T
	urlStr string
}

func (e *endpointT) Errorf(format string, args ...interface{}) {
	format = fmt.Sprintf("[%s] %s", e.urlStr, format)
	e.T.Errorf(format, args...)
}

func TestCountHandler(t *testing.T) {
	endpointT := &endpointT{t, "/count"}
	req, err := http.NewRequest(http.MethodGet, endpointT.urlStr, nil)
	if err != nil {
		endpointT.Fatalf("could not create request object: %v", err)
	}

	recorder := httptest.NewRecorder()
	handler := httpCountHandler()
	handler.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		endpointT.Errorf("returned wrong status code: got %v expected %v", status, http.StatusOK)
	}

	expectedBody := `{"size":0}`
	actualBody := strings.TrimSpace(recorder.Body.String())
	if actualBody != expectedBody {
		endpointT.Errorf("returned unexpected body: got %v expected %v", actualBody, expectedBody)
	}
}
