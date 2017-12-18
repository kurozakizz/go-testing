package handler

import (
	"net/http"
	"testing"
)

type StubSuccessCustomHandler struct{}

func (ch StubSuccessCustomHandler) Get(url string) (resp *http.Response, err error) {
	r := &http.Response{
		StatusCode: http.StatusOK,
	}
	return r, nil
}

func Test_ExistURL_ShouldReturn_StatusOK(t *testing.T) {
	httpChecker := StubSuccessCustomHandler{}
	status := GetHTTPStatus(httpChecker, "http://www.google.co.th")
	if status != http.StatusOK {
		t.Fatal("Expected status", http.StatusOK, "but got", status)
	}
}
