package store

import (
	"log"

	"github.com/sirupsen/logrus"
	"github.com/sprinteins/poker/src/x/types"
)

type roomID = string
type rooms = map[roomID]*Room

type Store struct {
	rooms rooms
}

func New() Store {

	return Store{
		rooms: make(rooms, 0),
	}
}

func (s *Store) AddPlayer(roomID types.ID, newPlayer Player) {
	room := s.EnsureRoom(roomID)
	room.AddPlayer(newPlayer)
}

func (s *Store) RemovePlayer(roomID types.ID, player Player) {
	room := s.EnsureRoom(roomID)
	room.RemovePlayer(player)
}

func (s Store) EnsureRoom(roomID types.ID) *Room {
	room, ok := s.rooms[roomID]
	if !ok {
		log.Printf("room not found createing, roomid=%s", roomID)
		newRoom := NewRoom(roomID)
		s.rooms[roomID] = &newRoom
		room = s.rooms[roomID]
	} else {
		log.Printf("room found with roomid=%s", roomID)
	}

	return room
}

func (s *Store) StartVoting(roomID types.ID) {
	room, ok := s.rooms[roomID]
	if !ok {
		logrus.WithFields(logrus.Fields{
			"roomID": roomID,
		}).Warn("room not found")
		return
	}

	room.StartVoting()

}
