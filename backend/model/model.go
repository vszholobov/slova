package model

import "encoding/json"

type UserCommand string

const (
	CommandSetUsername UserCommand = "SET_USERNAME"
	CommandStartNew    UserCommand = "START_NEW"
	CommandGuess       UserCommand = "GUESS"
	CommandStopWin     UserCommand = "STOP_WIN"
	CommandStopSkip    UserCommand = "STOP_SKIP"
)

type RawUserCommand struct {
	Command UserCommand     `json:"command"`
	Body    json.RawMessage `json:"body"`
}

type VoteCommandBody struct {
	Vote bool `json:"vote"`
}

type SetUsernameCommandBody struct {
	Username string `json:"username"`
}

type GuessCommandBody struct {
	Guess string `json:"guess"`
}

type CreateSessionResponse struct {
	ID string `json:"id"`
}
