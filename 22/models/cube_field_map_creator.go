package models

import (
	c "aoc/common"
	"fmt"
)

func CubeMapCreator(fields []Field) (FieldMap, error) {
	position_to_fields := c.CreateKeyValueMap(fields, Field.ToPosition, c.Identity[Field])
	var cube_side_length int
	var cube_side_fields [][]Field
	var cube_sides []*cube_side
	for _, processing_func := range []func() error{
		// First attempt to infer length
		func() error {
			temp_length, err := infer_length(len(fields))
			cube_side_length = temp_length
			return err
		},
		func() error {
			temp_cube_sides, err := split_fields_by_cube_sides(position_to_fields, cube_side_length)
			cube_side_fields = temp_cube_sides
			return err
		},
		func() error {
			temp_cube_sides, err := create_populated_cube_sides(cube_side_fields, cube_side_length)
			cube_sides = temp_cube_sides
			return err
		},
	} {
		if err := processing_func(); err != nil {
			return nil, err
		}
	}

	// create a mapping: position => it's encapsulating cube side
	position_to_cube_side := map[Position]*cube_side{}
	for _, cs := range cube_sides {
		c.ForEach(func(f Field) { position_to_cube_side[f.ToPosition()] = cs }, c.Flatten(cs.fields))
	}

	return CubeFieldMap{
		cube_sides:     cube_sides,
		fields:         position_to_fields,
		position_to_cs: position_to_cube_side,
	}, nil
}

const cfmc_error_prefix = "error while constructing the cube field map"

// Infer the length of a cube's side based on the number of fields
func infer_length(field_count int) (int, error) {
	length := 1
	for length*length*6 < field_count {
		length++
	}

	var err error = nil
	if length*length*6 != field_count {
		err = fmt.Errorf(
			`%s: could not find a suitable cube side length for %d field(s)`,
			cfmc_error_prefix, field_count,
		)
	}
	return length, err
}

// Split given fields by their membership to cube sides based on their positions and cube side length
func split_fields_by_cube_sides(positions_to_fields map[Position]Field, cube_side_length int) ([][]Field, error) {
	fields := sort_fields(c.GetValues(positions_to_fields))

	// An anchor field is the northmost westmost field of a cube side
	anchors := []Field{}
	side_membership_funcs := []func(Field) bool{}

	make_side_membership_func := func(anchor_field Field) func(Field) bool {
		row_check := c.InRange(anchor_field.Row, anchor_field.Row+cube_side_length)
		col_check := c.InRange(anchor_field.Column, anchor_field.Column+cube_side_length)
		return func(f Field) bool {
			return row_check(f.Row) && col_check(f.Column)
		}
	}

	for _, field := range fields {

		// If the current field belongs to one of the cube's sides, skip it
		if c.Any(func(smf func(Field) bool) bool { return smf(field) }, side_membership_funcs) {
			continue
		}

		// otherwise, it's an anchor and it defines another side of the cube
		anchors = append(anchors, field)
		side_membership_funcs = append(side_membership_funcs, make_side_membership_func(field))
	}

	// Partition fields by their assumed cube sides
	cube_side_fields := c.GetValues(c.GroupBy(
		fields,
		func(f Field) int {
			return c.IndexOf(side_membership_funcs, func(smf func(Field) bool) bool { return smf(f) })
		},
		c.Identity[Field],
	))

	// Check that there are exactly 6 sides and that each has expected number of fields on it
	if len(cube_side_fields) != 6 || c.Any(func(fs []Field) bool {
		return len(fs) != cube_side_length*cube_side_length
	}, cube_side_fields) {
		return nil, fmt.Errorf(
			`%s: failed at determining cube sides' fields`,
			cfmc_error_prefix,
		)
	}

	cube_side_fields = c.Map(sort_fields, cube_side_fields)
	for i, f := range cube_side_fields[0] {
		row, col := f.Row, f.Column
		if !c.All(func(fields []Field) bool {
			return (fields[i].Row-row)%cube_side_length == 0 && (fields[i].Column-col)%cube_side_length == 0
		}, cube_side_fields) {
			return nil, fmt.Errorf(`%s: cube side fields are not aligned`, cfmc_error_prefix)
		}
	}

	return cube_side_fields, nil
}

func create_populated_cube_sides(cube_side_fields [][]Field, cube_side_length int) ([]*cube_side, error) {
	cube_sides := create_cube_side_placeholders()
	anchors := c.Map(func(fields []Field) Field { return fields[0] }, cube_side_fields)

	index_to_id := map[int]int{0: 0}
	id_to_index := map[int]int{0: 0}

	queue := c.Queue(0)
	cube_sides[0].roan = North

	for !queue.IsEmpty() {
		anchor_idx, _ := queue.Dequeue()
		current_cs := cube_sides[index_to_id[anchor_idx]]
		anchor := anchors[anchor_idx]
		// fmt.Printf("CURRENT: %v\n", anchor)
		for _, abs_ori := range GetOrientations() {
			pos := abs_ori.move(anchor.ToPosition(), cube_side_length)
			nidx := c.IndexOf(anchors, func(f Field) bool { return f.ToPosition() == pos }) // neighbour's index
			if _, done := index_to_id[nidx]; nidx == -1 || done {
				// fmt.Printf("SKIP: %v\n", pos)
				continue
			}
			rel_ori := current_cs.get_orientation_absolute_to_relative(abs_ori)
			neighbour_cs := current_cs.neighbours[rel_ori]
			if _, done := id_to_index[neighbour_cs.id]; done {
				return nil, fmt.Errorf(
					`%s: cube side duplicated on the map`,
					cfmc_error_prefix,
				)
			}

			// register index-to-id and vice versa mappings with newly discovered neighbouring cube side
			id_to_index[neighbour_cs.id] = nidx
			index_to_id[nidx] = neighbour_cs.id

			queue.Enqueue(nidx)

			// Figure out the relative orientation that points to the absolute north for the neighbouring cube side
			neighbour_cs.relative_orientation_towards(current_cs)
			neighbour_cs.roan = get_orientation_rotator(
				abs_ori.opposite(),
				neighbour_cs.relative_orientation_towards(current_cs),
			)(North)
		}
	}

	// fmt.Printf("Cube sides identified: %d\n", len(id_to_index))
	// check that all cube sides have been discovered
	if len(id_to_index) != 6 {
		return nil, fmt.Errorf(`%s: failed to connect cube sides`, cfmc_error_prefix)
	}

	// assign each cube side its fields
	for i := range cube_side_fields {
		cube_sides[index_to_id[i]].assign_fields(cube_side_fields[i])
	}
	return cube_sides, nil
}

// Out of the fields provided get the northmost field out of the westmost fields
func get_westmost_northmost_point(fields []Field) Field {
	return c.MinimumBy(fields, func(lhs, rhs Field) bool {
		return lhs.Column < rhs.Column || (lhs.Column == rhs.Column && lhs.Row <= rhs.Row)
	})
}

// Given the fields, the cube's side westmost-northmost field and the cube's side length
// it returns a pair of fields (M, N) such that:
//   - Fields in M belong to the cube side
//   - Fields in N do not belong to the cube side
//   - M and N form a partition of the passed fields
func split_fields_by_cube_side_membership(
	fields []Field,
	westmost_northmost_position Field,
	cube_side_length int,
) ([]Field, []Field) {
	// Row & column number of the northmost-westmost field
	row := westmost_northmost_position.Row
	col := westmost_northmost_position.Column

	// function that checks whether a field belongs to the cube's side based on the northmost-westmost field
	membership_check := func(field Field) bool {
		return c.InRange(row, row+cube_side_length)(field.Row) && c.InRange(col, col+cube_side_length)(field.Column)
	}

	grouped := c.GroupBy(fields, membership_check, c.Identity[Field])
	return grouped[true], grouped[false]
}
