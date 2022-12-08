package store

import "github.com/gorilla/websocket"

type Player struct {
	ID         string          `json:"id"`
	Name       string          `json:"name"`
	IsAdmin    bool            `json:"isAdmin"`
	ChosenCard string          `json:"chosenCard"`
	Conn       *websocket.Conn `json:"-"`
	Type       PlayerType      `json:"type"`
	CardBack   CardBack        `json:"cardBack"`
}

type PlayerType = string

var (
	PlayerTypePlayer    = "Player"
	PlayerTypeSpectator = "Spectator"
)
