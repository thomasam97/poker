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
	s.CleanUpRoomIfNeeded(room)
}

func (s *Store) CleanUpRoomIfNeeded(room *Room) {
	hasPlayers := len(room.players) > 0
	if !hasPlayers {
		delete(s.rooms, room.id)
	}
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

func (s *Store) searchRoom(roomID types.ID) *Room {
	room, ok := s.rooms[roomID]
	if !ok {
		logrus.WithFields(logrus.Fields{
			"fn":     "StartVoting",
			"roomID": roomID,
		}).Warn("room not found")
		return nil
	}

	return room
}

func (s *Store) StartVoting(roomID types.ID) {
	room := s.searchRoom(roomID)
	if room == nil {
		return
	}

	room.StartVoting()

}

func (s *Store) Reveal(roomID types.ID) {
	room := s.searchRoom(roomID)
	if room == nil {
		return
	}

	room.Reveal()
}

func (s *Store) Reset(roomID types.ID) {
	room := s.searchRoom(roomID)
	if room == nil {
		return
	}

	room.Reset()
}

func (s *Store) Choose(roomID types.ID, playerID types.ID, card string) {
	room := s.searchRoom(roomID)
	if room == nil {
		return
	}

	room.Choose(playerID, card)
}
