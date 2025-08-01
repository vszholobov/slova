package main

import (
	"encoding/json"
	"net/http"
	"slova/model"
	"time"

	"math/rand"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func CreateSessionHandler(w http.ResponseWriter, r *http.Request) {
	sessionID := generateSessionID()
	response := model.CreateSessionResponse{
		ID: sessionID,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func ConnectToSessionHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sessionID := vars["id"]
	log.Infof("SessionId connect: %v", sessionID)

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Warnf("Failed to upgrade: %v", err)
		return
	}

	go gameHandler(conn, sessionID)
}

func gameHandler(conn *websocket.Conn, sessionID string) {
	defer conn.Close()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Warnf("Read error (%s): %v", sessionID, err)
			break
		}

		var rawCmd model.RawUserCommand
		if err := json.Unmarshal(message, &rawCmd); err != nil {
			log.Warnf("Invalid json (%s): %v", sessionID, err)
			continue
		}

		log.Infof("session %s received command: %s", sessionID, rawCmd.Command)

		switch rawCmd.Command {
		case model.CommandSetUsername:
			var body model.SetUsernameCommandBody
			if err := json.Unmarshal(rawCmd.Body, &body); err != nil {
				log.Warnf("invalid SET_USERNAME body: %v", err)
				continue
			}
			log.Infof("[%s] Set username: %s", sessionID, body.Username)

		case model.CommandGuess:
			var body model.GuessCommandBody
			if err := json.Unmarshal(rawCmd.Body, &body); err != nil {
				log.Warnf("invalid GUESS body: %v", err)
				continue
			}
			log.Infof("[%s] Guess: %s", sessionID, body.Guess)

		case model.CommandStopWin, model.CommandStopSkip:
			var body model.VoteCommandBody
			if err := json.Unmarshal(rawCmd.Body, &body); err != nil {
				log.Warnf("invalid vote body: %v", err)
				continue
			}
			log.Infof("[%s] Vote: %v", sessionID, body.Vote)

		case model.CommandStartNew:
			log.Infof("[%s] Start new round", sessionID)
			// Здесь можно вставить start logic

		default:
			log.Warnf("unknown command: %s", rawCmd.Command)
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
