package models

import "fmt"

type Motion struct {
	Steps     int
	Direction Direction
}

func MakeMotion(steps int, direction Direction) Motion {
	return Motion{steps, direction}
}

func (m Motion) String() string {
	return fmt.Sprintf("%s-%d", m.Direction, m.Steps)
}
