package store

type ActionType string
type Action struct {
	Type ActionType `json:"type"`
}

const (
	TypeStartVoting   = "StartVoting"
	TypeReveal        = "Reveal"
	TypeReset         = "Reset"
	TypeChoose        = "Choose"
	TypeSetPlayerType = "SetPlayerType"
	TypeSetCards      = "SetCards"
)
