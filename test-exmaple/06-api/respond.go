package respond

import (
	"log"
	"encoding/json"
	"net/http"
)

func transform (data interface{}) interface{} {
	switch o := data.(type) {
	case error:
		data = map[string]interface{}{"error": o.Error}
	}
	return data
}

func With(w http.ResponseWriter, r *http.Request, status int, data interface{}) {

	data = transform(data)
	b, err := json.Marshal(data)
	if err != nil {
		With(w, r, http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}

	w.WriteHeader(status)
	if _, err := w.Write(b); err != nil {
		log.Fatalln("failed to write:", err)
	}
}