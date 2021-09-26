package store

type ActionType string
type Action struct {
	Type ActionType `json:"type"`
}

const (
	TypeStartVoting = "StartVoting"
)
