package store

import "github.com/gorilla/websocket"

type Player struct {
	ID         string          `json:"id"`
	Name       string          `json:"name"`
	IsAdmin    bool            `json:"isAdmin"`
	ChosenCard string          `json:"chosenCard"`
	Conn       *websocket.Conn `json:"-"`
}
