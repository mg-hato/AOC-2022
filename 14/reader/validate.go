package reader

import (
	c "aoc/common"
	m "aoc/d14/models"
)

func verify_that_rock_paths_are_horizontal_vertical(rock_structures []m.RockStructure) error {
	for rock_formation_number, rock_structure := range rock_structures {
		for i := 1; i < len(rock_structure); i++ {
			if rock_structure[i-1].First != rock_structure[i].First &&
				rock_structure[i-1].Second != rock_structure[i].Second {
				return rock_structures_are_not_horizontal_or_vertical_error(rock_formation_number+1, i)
			}
		}
	}
	return nil
}

func verify_that_no_rock_formation_overlaps_sand_source(sand_source c.Pair[int, int]) func([]m.RockStructure) error {
	return func(rock_structures []m.RockStructure) error {

		in_range := func(limit_1, limit_2 int) func(int) bool {
			switch {
			case limit_1 < limit_2:
				return c.InInclusiveRange(limit_1, limit_2)
			default:
				return c.InInclusiveRange(limit_2, limit_1)
			}
		}

		for rock_formation_number, rock_structure := range rock_structures {
			for i := 1; i < len(rock_structure); i++ {

				if in_range(rock_structure[i-1].First, rock_structure[i].First)(sand_source.First) &&
					in_range(rock_structure[i-1].Second, rock_structure[i].Second)(sand_source.Second) {
					return rock_formation_overlaps_sand_source_error(rock_formation_number+1, sand_source)
				}
			}
		}
		return nil
	}
}
