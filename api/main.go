package main

import (
	"fmt"
	"github.com/gorilla/mux"

	"log"
	"net/http"
	"os"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", IndexHandler).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}

// IndexHandler response hello
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "hello database connection: "+os.Getenv("DB_HOST"))
}
