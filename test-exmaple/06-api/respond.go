package respond

import (
	"log"
	"encoding/json"
	"net/http"
)

func With(w http.ResponseWriter, r *http.Request, status int, body interface{}) {
	b, err := json.Marshal(body)
	if err != nil {
		With(w, r, http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}

	w.WriteHeader(status)
	if _, err := w.Write(b); err != nil {
		log.Fatalln("failed to write:", err)
	}
}