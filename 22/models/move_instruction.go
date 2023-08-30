package models

import "fmt"

type MoveInstruction struct {
	distance int
}

func CreateMoveInstruction(d int) Instruction {
	return MoveInstruction{d}
}

func (mi MoveInstruction) String() string {
	return fmt.Sprintf("<INSTR: move(%d)>", mi.distance)
}

func (mi MoveInstruction) GetMovement() int {
	return mi.distance
}

func (mi MoveInstruction) GetOrientation(current Orientation) Orientation {
	return current
}

func (mi MoveInstruction) Equals(instr Instruction) bool {
	other, is_move_instr := instr.(MoveInstruction)
	return is_move_instr && other.distance == mi.distance
}
