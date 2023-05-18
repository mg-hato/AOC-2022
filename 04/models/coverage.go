package models

import "fmt"

// Elf's coverage: covers sections from Left to Right (inclusive)
type Coverage struct {
	Left  int
	Right int
}

func (c Coverage) String() string {
	return fmt.Sprintf("%d-%d", c.Left, c.Right)
}
