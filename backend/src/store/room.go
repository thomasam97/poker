package store

import (
	"encoding/json"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/sprinteins/poker/src/x/types"
)

type Room struct {
	id               types.ID
	players          []*Player
	mtx              sync.Mutex
	status           Status
	cards            Cards
	autoReveal       bool
	timeboxInSeconds uint
	timer            *time.Timer
}

var defaultCards = cardSets[0]

func NewRoom(id types.ID) Room {
	return Room{
		id:               id,
		players:          make([]*Player, 0),
		status:           StatusStart,
		cards:            defaultCards.Cards,
		autoReveal:       false,
		timeboxInSeconds: 0,
		timer:            nil,
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

func (r *Room) findPlayerIndexByID(playerID types.ID) int {
	wantedIndex := -1
	for pi, player := range r.players {
		if player.ID == playerID {
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
	Player           *Player   `json:"player"`
	Players          []*Player `json:"players"`
	Status           Status    `json:"status"`
	Cards            Cards     `json:"cards"`
	Sets             []Set     `json:"sets"`
	AutoReveal       bool      `json:"autoReveal"`
	TimeboxInSeconds uint      `json:"timeboxInSeconds"`
}

func (r *Room) EmitState() {

	if len(r.players) > 0 {
		r.players[0].IsAdmin = true
	}

	for _, player := range r.players {

		state := State{
			Player:           player,
			Players:          r.players,
			Status:           r.status,
			Cards:            r.cards,
			Sets:             cardSets,
			AutoReveal:       r.autoReveal,
			TimeboxInSeconds: r.timeboxInSeconds,
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

	log.WithFields(log.Fields{
		"roomID": r.id,
	}).Info("start voting")

	r.status = StatusInProgress

	if r.timer != nil {
		log.Printf("stopping timer of previous voting round for roomid=%s", r.id)
		r.timer.Stop()
	}
	go r.RevealIfTimeboxIsUp()
	r.EmitState()
}

func (r *Room) Reveal() {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	if r.status == StatusRevealed {
		return
	}

	log.WithFields(log.Fields{
		"roomID": r.id,
	}).Info("reveal cards")

	r.status = StatusRevealed
	r.EmitState()
}

func (r *Room) Reset() {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	log.WithFields(log.Fields{
		"roomID": r.id,
	}).Info("reset room")

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

func (r *Room) ReVote(playerID types.ID) {
	r.Choose(playerID, "")
}

func (r *Room) AutoRevealIfCan() {
	if r.autoReveal && r.HasEverybodyChosen() {
		r.Reveal()
	}
}

func (r *Room) RevealIfTimeboxIsUp() {
	if r.timeboxInSeconds == 0 {
		return
	}

	r.timer = time.NewTimer(time.Duration(r.timeboxInSeconds) * time.Second)
	<-r.timer.C
	r.Reveal()

}

func (r *Room) HasEverybodyChosen() bool {
	for _, player := range r.players {
		if player.Type == PlayerTypePlayer && player.ChosenCard == "" {
			return false
		}
	}

	return true
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

func (r *Room) SwitchCardBack(playerID types.ID, cardback CardBack) {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	player := r.findPlayerByID(playerID)
	if player == nil {
		return
	}
	log.WithFields(log.Fields{
		"cardBack": cardback,
		"playerID": playerID,
		"roomID":   r.id,
	}).Info("switch card back")
	player.CardBack = cardback
	r.EmitState()
}

func (r *Room) SetCards(playerID types.ID, cards Cards) {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	player := r.findPlayerByID(playerID)
	if player == nil {
		return
	}
	log.WithFields(log.Fields{
		"cards":    cards,
		"playerID": playerID,
		"roomID":   r.id,
	}).Info("set player type")
	r.cards = cards
	r.EmitState()
}

func (r *Room) SetAutoReveal(autoReveal bool) {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	log.WithFields(log.Fields{
		"autoReveal": autoReveal,
		"roomID":     r.id,
	}).Info("set auto reveal")
	r.autoReveal = autoReveal
	r.EmitState()
}

func (r *Room) SetTimebox(timeboxInSeconds uint) {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	r.timeboxInSeconds = timeboxInSeconds
	log.WithFields(log.Fields{
		"timebox": r.timeboxInSeconds,
		"roomID":  r.id,
	}).Info("set timebox seconds")
	r.EmitState()
}

func (r *Room) SetAdmin(newAdminID types.ID) {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	log.WithFields(log.Fields{
		"newAdminID": newAdminID,
		"roomID":     r.id,
	}).Info("set admin")
	newAdminIndex := r.findPlayerIndexByID(newAdminID)

	newAdmin := r.players[newAdminIndex]
	oldAdmin := r.players[0]

	newAdmin.IsAdmin = true
	oldAdmin.IsAdmin = false

	r.players[0] = newAdmin
	r.players[newAdminIndex] = oldAdmin

	r.EmitState()
}
