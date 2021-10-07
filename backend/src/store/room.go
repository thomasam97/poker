package store

import (
	"encoding/json"
	"sync"

	log "github.com/sirupsen/logrus"
	"github.com/sprinteins/poker/src/x/types"
)

type Room struct {
	id      types.ID
	players []*Player
	mtx     sync.Mutex
	status  Status
}

func NewRoom(id types.ID) Room {
	return Room{
		id:      id,
		players: make([]*Player, 0),
		status:  StatusStart,
	}
}

func (r *Room) AddPlayer(p Player) {
	r.mtx.Lock()
	r.players = append(r.players, &p)
	r.EmitState()
	r.mtx.Unlock()
}

func (r *Room) findPlayerIndex(p Player) int {
	wantedIndex := -1
	for pi, player := range r.players {
		if player.ID == p.ID {
			return pi
		}
	}
	return wantedIndex
}

func (r *Room) findPlayerByID(playerID types.ID) *Player {
	for _, player := range r.players {
		if player.ID == playerID {
			return player
		}
	}
	log.WithFields(log.Fields{
		"playerID": playerID,
		"roomID":   r.id,
	}).Error("could not find player")

	return nil
}

func (r *Room) RemovePlayer(p Player) {
	r.mtx.Lock()
	index := r.findPlayerIndex(p)
	r.players = append(r.players[:index], r.players[index+1:]...)
	r.EmitState()
	r.mtx.Unlock()
}

type State struct {
	Player  *Player   `json:"player"`
	Players []*Player `json:"players"`
	Status  Status    `json:"status"`
}

func (r *Room) EmitState() {

	if len(r.players) > 0 {
		r.players[0].IsAdmin = true
	}

	for _, player := range r.players {

		state := State{
			Player:  player,
			Players: r.players,
			Status:  r.status,
		}
		payload, err := json.Marshal(state)
		if err != nil {
			log.Error(err)
		}

		player.Conn.WriteMessage(1, payload)
	}

}

func (r *Room) StartVoting() {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	r.status = StatusInProgress
	r.EmitState()
}

func (r *Room) Reveal() {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	r.status = StatusRevealed
	r.EmitState()
}

func (r *Room) Reset() {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	r.status = StatusStart
	for _, player := range r.players {
		player.ChosenCard = ""
	}
	r.EmitState()
}

func (r *Room) Choose(playerID types.ID, card string) {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	player := r.findPlayerByID(playerID)
	if player == nil {
		return
	}
	log.WithFields(log.Fields{
		"card":     card,
		"playerID": playerID,
		"roomID":   r.id,
	}).Info("choosing card")
	player.ChosenCard = card
	r.EmitState()
}

func (r *Room) SetPlayerType(playerID types.ID, playerType PlayerType) {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	player := r.findPlayerByID(playerID)
	if player == nil {
		return
	}
	log.WithFields(log.Fields{
		"playerType": playerType,
		"playerID":   playerID,
		"roomID":     r.id,
	}).Info("set player type")
	player.Type = playerType
	r.EmitState()
}
