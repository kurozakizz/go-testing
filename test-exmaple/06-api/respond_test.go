package respond

import (
	"net/http/httptest"
	"testing"
	"net/http"
)

func TestAll(t *testing.T){

	tests := []struct {
		Fn func (w http.ResponseWriter, r *http.Request)
		StatusCode int
		Body string
	}{
		{
			Fn: func (w http.ResponseWriter, r *http.Request) {
				With(w, r, http.StatusOK, "Hello world")
			},
			StatusCode: 200,
			Body: `"Hello world"`,
		},
	}

	for _, test := range tests {
		w := httptest.NewRecorder()
		test.Fn(w, nil)

		if test.StatusCode != w.Code {
			t.Fatalf("Expected %d but got %d", test.StatusCode, w.Code)
		}

		if test.Body != w.Body.String() {
			t.Fatalf("Expected %s but got %s", test.Body, w.Body.String())
		}
	}
}
