package models

import "fmt"

type Move struct {
	Quantity    int
	Source      int
	Destination int
}

func (m Move) String() string {
	return fmt.Sprintf("M[%d: %d -> %d]", m.Quantity, m.Source, m.Destination)
}

// Helper constructor function
func MakeMove(quantity, source, destination int) Move {
	return Move{
		Quantity:    quantity,
		Source:      source,
		Destination: destination,
	}
}
