package models

import "fmt"

type TurnInstruction struct {
	rotation Rotation
}

func CreateTurnInstruction(rot Rotation) Instruction {
	return TurnInstruction{rot}
}

func (ti TurnInstruction) String() string {
	return fmt.Sprintf("<INSTR: turn(%s)>", ti.rotation.String())
}

func (ti TurnInstruction) GetMovement() int {
	return 0
}

func (ti TurnInstruction) GetOrientation(current Orientation) Orientation {
	return Orientation((int(current) + int(ti.rotation) + 4) % 4)
}

func (ti TurnInstruction) Equals(instr Instruction) bool {
	other, is_turn_instr := instr.(TurnInstruction)
	return is_turn_instr && other.rotation == ti.rotation
}
