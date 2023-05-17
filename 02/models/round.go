package models

import "fmt"

// Round of rock-paper-scissors is described as two symbols
//  - Left symbol (A, B or C)
//  - and Right symbol (X, Y or Z)
type Round struct {
	Left  LeftSymbol
	Right RightSymbol
}

func (r Round) String() string {
	return fmt.Sprintf("<%s:%s>", r.Left, r.Right)
}
