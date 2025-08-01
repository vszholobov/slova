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
	CheckOrigin: func(r *http.Request) bool {
		// ⚠️ Разрешает ВСЕ запросы (в проде лучше фильтровать)
		return true
	},
}

type SessionStatus string

const (
	WIN    SessionStatus = "WIN"
	SKIP   SessionStatus = "SKIP"
	ACTIVE SessionStatus = "ACTIVE"
	NEW    SessionStatus = "NEW"
)

type GameSession struct {
	sessionId string
	Status    SessionStatus    `json:"status"`
	Vote      *SessionVote     `json:"vote"`
	Guesses   *SessionGuesses  `json:"guesses"`
	Players   *SessionPlayers  `json:"players"`
	History   []SessionHistory `json:"history"`
}

type GameSessionState struct {
	GameSession
	You string `json:"you"`
}

type SessionVote struct {
	// TODO: поля с json-тегами
}

type SessionGuesses struct {
	Player1 []string `json:"player1"`
	Player2 []string `json:"player2"`
}

type SessionPlayers struct {
	Player1 *SessionPlayer `json:"player1"`
	Player2 *SessionPlayer `json:"player2"`
}

type SessionPlayer struct {
	Username string `json:"username"`
}

type SessionHistory struct {
	// TODO: поля с json-тегами
}

var Sessions = make(map[string]*GameSession)

func MakeGameSession(sessionId string) *GameSession {
	return &GameSession{
		sessionId: sessionId,
		Status:    NEW,
		Vote:      nil,
		Guesses: &SessionGuesses{
			Player1: make([]string, 0),
			Player2: make([]string, 0),
		},
		Players: &SessionPlayers{},
		History: nil,
	}
}

func CreateSessionHandler(w http.ResponseWriter, r *http.Request) {
	sessionId := generateSessionId()
	Sessions[sessionId] = MakeGameSession(sessionId)
	response := model.CreateSessionResponse{
		ID: sessionId,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func ConnectToSessionHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sessionId := vars["id"]
	log.Infof("SessionId connect: %v", sessionId)

	session, ok := Sessions[sessionId]
	if !ok {
		log.Warnf("Session %v not found", sessionId)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Warnf("Failed to upgrade: %v", err)
		return
	}

	var playerNum string
	if session.Players.Player1 == nil {
		playerNum = "1"
		session.Players.Player1 = &SessionPlayer{
			Username: "player1",
		}
	} else {
		playerNum = "2"
		session.Players.Player2 = &SessionPlayer{
			Username: "player2",
		}
	}

	go gameHandler(conn, session, playerNum)
}

func gameHandler(conn *websocket.Conn, session *GameSession, playerNum string) {
	defer conn.Close()

	sessionId := session.sessionId

	for {
		conn.WriteJSON(GameSessionState{
			GameSession: *session,
			You:         playerNum,
		})
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Warnf("Read error (%s): %v", sessionId, err)
			break
		}

		var rawCmd model.RawUserCommand
		if err := json.Unmarshal(message, &rawCmd); err != nil {
			log.Warnf("Invalid json (%s): %v", sessionId, err)
			continue
		}

		log.Infof("session %s received command: %s", sessionId, rawCmd.Command)

		switch rawCmd.Command {
		case model.CommandSetUsername:
			var body model.SetUsernameCommandBody
			if err := json.Unmarshal(rawCmd.Body, &body); err != nil {
				log.Warnf("invalid SET_USERNAME body: %v", err)
				continue
			}
			log.Infof("[%s] Set username: %s", sessionId, body.Username)

		case model.CommandGuess:
			var body model.GuessCommandBody
			if err := json.Unmarshal(rawCmd.Body, &body); err != nil {
				log.Warnf("invalid GUESS body: %v", err)
				continue
			}
			log.Infof("[%s] Guess: %s", sessionId, body.Guess)

		case model.CommandStopWin, model.CommandStopSkip:
			var body model.VoteCommandBody
			if err := json.Unmarshal(rawCmd.Body, &body); err != nil {
				log.Warnf("invalid vote body: %v", err)
				continue
			}
			log.Infof("[%s] Vote: %v", sessionId, body.Vote)

		case model.CommandStartNew:
			log.Infof("[%s] Start new round", sessionId)
			// Здесь можно вставить start logic

		default:
			log.Warnf("unknown command: %s", rawCmd.Command)
		}
	}
}

func generateSessionId() string {
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
