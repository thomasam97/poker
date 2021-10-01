package store

import "github.com/gorilla/websocket"

type Player struct {
	ID      string          `json:"id"`
	Name    string          `json:"name"`
	IsAdmin bool            `json:"isAdmin"`
	Conn    *websocket.Conn `json:"-"`
}
