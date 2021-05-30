package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestFibHandlerValidEntry call's the FibHandler function with valid URL parameter,
// valid parameters: N -> Positive Integer
func TestFibHandlerValidEntry(t *testing.T) {
	testValue := "1"
	req, err := http.NewRequest("GET", "/fib?n="+testValue, nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(FibHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := fmt.Sprintf("{result: %s}\n", testValue)
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

// TestFibHandlerInvalidEntry call's the FibHandler function with invalid URL parameter,
// uses a Negative Integer and check both error status and message
func TestFibHandlerInvalidEntry(t *testing.T) {
	testValue := "-15"
	req, err := http.NewRequest("GET", "/fib?n="+testValue, nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(FibHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusUnprocessableEntity {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusUnprocessableEntity)
	}

	expected := "INVALID_ENTRY\n"

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
