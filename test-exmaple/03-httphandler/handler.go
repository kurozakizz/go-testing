package handler

import (
	"io"
	"net/http"
)

type handler struct {
	EndPointUrl string
}

func (h handler) Status(w http.ResponseWriter, r *http.Request) {
	response, err := http.Get(h.EndPointUrl)
	defer response.Body.Close()

	if err != nil {
		http.Error(w, "Failed to ping endpoint", http.StatusServiceUnavailable)
		return
	}

	if response.StatusCode != http.StatusOK {
		http.Error(w, "Failed to ping endpoint", response.StatusCode)
		return
	}

	io.WriteString(w, "Status OK. Endpoint is Alive")
}
