package reader

import (
	c "aoc/common"
	m "aoc/d22/models"
	"fmt"
)

func validate_at_least_one_dot(mmr monkey_map_reader) error {
	dot_count := c.Count(mmr.fields, func(field m.Field) bool { return field.FType == m.Dot })
	if dot_count == 0 {
		return map_does_not_contain_walkable_paths_error()
	}
	return nil
}

func validate_continuity(mmr monkey_map_reader) error {

	// continuity validation function with group-by-key & selector extractors
	validate_continuity := func(
		group_by_str string,
		group_by_func, select_func func(m.Field) int,
	) error {
		groups := c.GroupBy(mmr.fields, group_by_func, select_func)
		for group_key, group_values := range groups {
			values := c.CreateSet(group_values, c.Identity[int])
			min_val, max_val := c.Minimum(group_values), c.Maximum(group_values)
			for i := min_val; i <= max_val; i++ {
				if !values[i] {
					return discountinuity_in_map_detected(fmt.Sprintf("%s #%d", group_by_str, group_key))
				}
			}
		}
		return nil
	}

	// Verify whether there is any discontinuity in columns
	discontinuity_in_columns := validate_continuity("column", m.Field.GetColumn, m.Field.GetRow)
	if discontinuity_in_columns != nil {
		return discontinuity_in_columns
	}

	// Verify whether there is any discontinuity in rows
	discontinuity_in_rows := validate_continuity("row", m.Field.GetRow, m.Field.GetColumn)
	if discontinuity_in_rows != nil {
		return discontinuity_in_rows
	}

	return nil
}
