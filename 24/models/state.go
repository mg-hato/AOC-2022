package models

type State struct {
	CurrentPosition Position
	PassedTime      int
}

func MakeState(position Position, time int) State {
	return State{position, time}
}
