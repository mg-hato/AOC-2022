package models

import c "aoc/common"

type cube_side struct {
	id int

	neighbours map[Orientation]*cube_side

	roan Orientation // Relative orientation that points to the absolute north

	fields [][]Field

	pfields map[Position]Field // position to field mapping
}

func (cs *cube_side) relative_orientation_towards(other *cube_side) Orientation {
	for orientation, cs := range cs.neighbours {
		if cs == other {
			return orientation
		}
	}
	return Orientation(-1)
}

func (cs *cube_side) get_orientation_absolute_to_relative(absolute Orientation) Orientation {
	return get_orientation_rotator(North, cs.roan)(absolute)
}

func (cs *cube_side) get_orientation_relative_to_absolute(relative Orientation) Orientation {
	return get_orientation_rotator(cs.roan, North)(relative)
}

func (cs *cube_side) assign_fields(fields []Field) {
	cs.pfields = c.CreateKeyValueMap(fields, Field.ToPosition, c.Identity[Field])
	grouped_by_row := c.GroupBy(fields, Field.GetRow, c.Identity[Field])
	min_row := c.Minimum(c.GetKeys(grouped_by_row))
	max_row := c.Maximum(c.GetKeys(grouped_by_row))

	cs.fields = make([][]Field, max_row-min_row+1)
	for i := 0; i+min_row <= max_row; i++ {
		cs.fields[i] = sort_fields(grouped_by_row[i+min_row])
	}
}

func (cs *cube_side) get_next_pointer(ptr Pointer) Pointer {
	// If next position is on the same cube side, orientation does not change
	if field, exists := cs.pfields[ptr.orientation.move(ptr.position, 1)]; exists {
		return Pointer{
			position:    field.ToPosition(),
			orientation: ptr.orientation,
		}
	}

	// otherwise, we need to update orientation and determine which cube side the next position will be

	// relative orientation towards which we are going (from current cube side)
	relative_to := cs.get_orientation_absolute_to_relative(ptr.orientation)

	// neighbouring cube side on which we will appear next
	ncs := cs.neighbours[relative_to]

	// relative orientation from which we will come (from perspective of next cube side)
	relative_from := ncs.relative_orientation_towards(cs)

	origin_side := cs.get_ordered_side(relative_to)
	destination_side := ncs.get_ordered_side(relative_from)
	idx := c.IndexOf(origin_side, func(p Position) bool { return p == ptr.position })

	return Pointer{
		orientation: ncs.get_orientation_relative_to_absolute(relative_from.opposite()),
		position:    destination_side[idx],
	}
}

func (cs *cube_side) get_ordered_side(relative Orientation) []Position {
	// positions that are facing towards relative orientation (unordered)
	side_positions := cs.get_orientation_relative_to_absolute(relative).most(c.GetKeys(cs.pfields))

	// order them by adjacent neighbours
	return cs.get_orientation_relative_to_absolute(c.MaximumBy(
		relative.GetAdjacentOrientations(),
		func(lhs, rhs Orientation) bool {
			return cs.neighbours[lhs].id < cs.neighbours[rhs].id
		},
	)).order_towards(side_positions)
}

func create_cube_side_placeholders() []*cube_side {
	sides := make([]*cube_side, 6)
	for i, _ := range sides {
		sides[i] = &cube_side{id: i, neighbours: make(map[Orientation]*cube_side)}
	}

	// link cube's sides in zig-zag east-north fashion
	for i, _ := range sides {
		n := (i + 1) % 6
		if i%2 == 0 {
			sides[i].neighbours[East] = sides[n]
			sides[n].neighbours[West] = sides[i]
		} else {
			sides[i].neighbours[North] = sides[n]
			sides[n].neighbours[South] = sides[i]
		}
	}
	// Populate missing neighbour:
	// - North: do East and then North
	// - West: do South and then West
	// - South: do West and then South
	// - East: do North and then East
	for i, _ := range sides {
		if _, has_north := sides[i].neighbours[North]; !has_north {
			sides[i].neighbours[North] = sides[i].neighbours[East].neighbours[North]
		}
		if _, has_west := sides[i].neighbours[West]; !has_west {
			sides[i].neighbours[West] = sides[i].neighbours[South].neighbours[West]
		}
		if _, has_south := sides[i].neighbours[South]; !has_south {
			sides[i].neighbours[South] = sides[i].neighbours[West].neighbours[South]
		}
		if _, has_east := sides[i].neighbours[East]; !has_east {
			sides[i].neighbours[East] = sides[i].neighbours[North].neighbours[East]
		}
	}
	return sides
}
