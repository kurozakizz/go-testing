package handler

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStatus_OK_ShouldReturn_StatusOk(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "text/plain")
		io.WriteString(w, "Status OK")

	}))
	defer ts.Close()

	r, err := http.NewRequest("GET", "/StatusCheck", nil)
	if err != nil {
		t.Errorf("expected no error got %v", err)
	}

	w := httptest.NewRecorder()

	h := handler{
		EndPointUrl: ts.URL,
	}

	h.Status(w, r)

	if w.Body.String() != "Status OK. Endpoint is Alive" {
		t.Fatalf("Expected Status OK but got %v", w.Body.String())
	}

	if w.Code != http.StatusOK {
		t.Fatalf("Expected 200 but got %v", w.Code)
	}
}

func TestStatus_Failed_ShoulReturn_ServiceUnavailableStatus(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Failed", http.StatusServiceUnavailable)
	}))
	defer ts.Close()

	r, err := http.NewRequest("GET", "/StatusCheck", nil)
	if err != nil {
		t.Errorf("expected no error got %v", err)
	}

	w := httptest.NewRecorder()

	h := handler{
		EndPointUrl: ts.URL,
	}

	h.Status(w, r)
	if w.Code != http.StatusServiceUnavailable {
		t.Fatalf("Expected %v but got %v", http.StatusServiceUnavailable, w.Code)
	}
}
