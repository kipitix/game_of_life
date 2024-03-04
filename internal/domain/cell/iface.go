package cell

//go:generate enumgen

type State int //enums:enum

const (
	Dead State = iota
	Life
)

type Cell interface {
	State() State
	SetState(State)
}
