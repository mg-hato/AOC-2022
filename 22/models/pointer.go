package models

import "fmt"

type Pointer struct {
	orientation Orientation
	position    Position
}

func MakePointer(row, col int, ori Orientation) Pointer {
	return Pointer{
		position:    MakePosition(row, col),
		orientation: ori,
	}
}

func (p Pointer) GetRow() int {
	return p.position.First
}

func (p Pointer) GetColumn() int {
	return p.position.Second
}

func (p Pointer) GetPosition() Position {
	return p.position
}

func (p Pointer) String() string {
	return fmt.Sprintf("(%d, %d, %s)", p.position.First, p.position.Second, p.orientation.String())
}

func (p Pointer) GetPassword() int {
	return 1_000*p.GetRow() + 4*p.GetColumn() + p.orientation.GetFacingValue()
}
