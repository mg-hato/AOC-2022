package models

import (
	c "aoc/common"
	"sort"
)

type CubeFieldMap struct {
	fields         map[Position]Field
	position_to_cs map[Position]*cube_side
	cube_sides     []*cube_side
}

func (cfm CubeFieldMap) GetInitialPointer() Pointer {
	dot_fields := c.Filter(func(field Field) bool { return field.FType == Dot }, c.GetValues(cfm.fields))
	sort.Slice(dot_fields, func(i, j int) bool {
		return dot_fields[i].Row < dot_fields[j].Row ||
			(dot_fields[i].Row == dot_fields[j].Row && dot_fields[i].Column < dot_fields[j].Column)
	})
	return Pointer{
		position:    dot_fields[0].ToPosition(),
		orientation: East,
	}
}

func (cfm CubeFieldMap) UpdatePointer(ptr Pointer, instr Instruction) Pointer {
	ptr.orientation = instr.GetOrientation(ptr.orientation)
	obstructed := false
	var i int = 0
	for i < instr.GetMovement() && !obstructed {
		next_ptr := cfm.position_to_cs[ptr.position].get_next_pointer(ptr)
		obstructed = cfm.fields[next_ptr.position].FType == Wall
		if !obstructed {
			ptr = next_ptr
		}
		i++
	}
	return ptr
}
