package main

import (
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("Start...")
	r := mux.NewRouter()

	r.HandleFunc("/session", CreateSessionHandler).Methods("POST")
	r.HandleFunc("/session/{id}", ConnectToSessionHandler).Methods("GET")

	log.Info("Server running on :80")
	log.Fatal(http.ListenAndServe(":80", r))
}
