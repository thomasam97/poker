package store

type Status string

const (
	StatusStart      Status = "Start"
	StatusInProgress Status = "InProgress"
	StatusRevealed   Status = "Revealed"
)
