package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/session/", CreateSessionHandler).Methods("POST")
	r.HandleFunc("/session/{id}", ConnectToSessionHandler).Methods("GET")

	handler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
	}).Handler(r)

	log.Info("Server running on :80")
	log.Fatal(http.ListenAndServe(":80", handler))
}
