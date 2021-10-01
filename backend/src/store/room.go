package store

import (
	"encoding/json"
	"sync"

	log "github.com/sirupsen/logrus"
	"github.com/sprinteins/poker/src/x/types"
)

type Room struct {
	id      types.ID
	players []Player
	mtx     sync.Mutex
	status  Status
}

func NewRoom(id types.ID) Room {
	return Room{
		id:      id,
		players: make([]Player, 0),
		status:  StatusStart,
	}
}

func (r *Room) AddPlayer(p Player) {
	r.mtx.Lock()
	r.players = append(r.players, p)
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

func (r *Room) RemovePlayer(p Player) {
	r.mtx.Lock()
	index := r.findPlayerIndex(p)
	r.players = append(r.players[:index], r.players[index+1:]...)
	r.EmitState()
	r.mtx.Unlock()
}

type State struct {
	Player  Player   `json:"player"`
	Players []Player `json:"players"`
	Status  Status   `json:"status"`
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
	r.status = StatusInProgress
	r.EmitState()
	r.mtx.Unlock()
}

func (r *Room) Reveal() {
	r.mtx.Lock()
	r.status = StatusRevealed
	r.EmitState()
	r.mtx.Unlock()
}

func (r *Room) Reset() {
	r.mtx.Lock()
	r.status = StatusStart
	r.EmitState()
	r.mtx.Unlock()
}
