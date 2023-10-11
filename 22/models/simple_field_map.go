package models

import (
	c "aoc/common"
	"sort"
)

type SimpleFieldMap struct {
	fields             []Field
	map_by_orientation map[Orientation]map[int][]Field
}

func CreateSimpleFieldMap(fields []Field) (FieldMap, error) {
	create_map_by_orientation := func(
		group_key_func, sort_key_func func(Field) int,
		order func(int, int) bool,
	) map[int][]Field {
		group := c.GroupBy(fields, group_key_func, c.Identity[Field])
		for key := range group {
			sort.Slice(group[key], func(i, j int) bool {
				return order(sort_key_func(group[key][i]), sort_key_func(group[key][j]))
			})
		}
		return group
	}

	return SimpleFieldMap{
		fields: c.ShallowCopy(fields),
		map_by_orientation: map[Orientation]map[int][]Field{
			North: create_map_by_orientation(Field.GetColumn, Field.GetRow, c.GreaterThan[int]),
			East:  create_map_by_orientation(Field.GetRow, Field.GetColumn, c.LessThan[int]),
			South: create_map_by_orientation(Field.GetColumn, Field.GetRow, c.LessThan[int]),
			West:  create_map_by_orientation(Field.GetRow, Field.GetColumn, c.GreaterThan[int]),
		},
	}, nil
}

func (sfm SimpleFieldMap) GetInitialPointer() Pointer {
	dot_positions := c.Map(
		Field.ToPosition,
		c.Filter(func(field Field) bool { return field.FType == Dot }, sfm.fields),
	)
	return Pointer{
		position:    West.most(North.most(dot_positions))[0],
		orientation: East,
	}

	columns := c.GetKeys(sfm.map_by_orientation[East])
	min, max := c.Minimum(columns), c.Maximum(columns)
	for column := min; column <= max; column++ {
		for _, field := range sfm.map_by_orientation[South][column] {
			if field.FType == Dot {
				return Pointer{
					orientation: East,
					position:    MakePosition(field.Row, field.Column),
				}
			}
		}
	}
	return Pointer{}
}

func (sfm SimpleFieldMap) get_path(ptr Pointer) []Field {
	if ptr.orientation == North || ptr.orientation == South {
		return sfm.map_by_orientation[ptr.orientation][ptr.GetColumn()]
	}
	if ptr.orientation == East || ptr.orientation == West {
		return sfm.map_by_orientation[ptr.orientation][ptr.GetRow()]
	}
	return nil
}

func (sfm SimpleFieldMap) UpdatePointer(ptr Pointer, instr Instruction) Pointer {
	ptr.orientation = instr.GetOrientation(ptr.orientation)
	path := sfm.get_path(ptr)
	i := c.IndexOf(path, func(f Field) bool { return f.Row == ptr.GetRow() && f.Column == ptr.GetColumn() })
	steps := 0
	for steps < instr.GetMovement() {
		next := (i + 1) % len(path)
		if path[next].FType != Dot {
			steps = instr.GetMovement()
		} else {
			i = next
			steps++
		}
	}
	ptr.position = MakePosition(path[i].Row, path[i].Column)
	return ptr
}
