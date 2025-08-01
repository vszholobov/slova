package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

type CreateSessionResponse struct {
	ID string `json:"id"`
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	log.Info("Start...")
	r := mux.NewRouter()

	r.HandleFunc("/session", createSessionHandler).Methods("POST")
	r.HandleFunc("/session/{id}", connectToSessionHandler).Methods("GET")

	log.Info("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func createSessionHandler(w http.ResponseWriter, r *http.Request) {
	sessionID := generateSessionID()
	response := CreateSessionResponse{
		ID: sessionID,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func connectToSessionHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sessionID := vars["id"]
	log.Infof("SessionId connect: %v", sessionID)

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Errorf("failed to upgrade: %v", err)
		return
	}
	defer conn.Close()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Printf("read error: %v", err)
			break
		}

		log.Printf("received: %s", message)

		err = conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			log.Printf("write error: %v", err)
			break
		}
	}
}

func generateSessionID() string {
	rand.Seed(time.Now().UnixNano())
	return randomDigits(8)
}

func randomDigits(length int) string {
	digits := make([]byte, length)
	for i := 0; i < length; i++ {
		digits[i] = '0' + byte(rand.Intn(10))
	}
	return string(digits)
}
