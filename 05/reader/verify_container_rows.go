package reader

import (
	"aoc/day05/models"
	f "aoc/functional"
	"regexp"
	"strconv"
	"strings"
)

// Verifies container rows and stack ids.
// Returns a list of starting container stacks and any error as a result of failed verification.
// Expects as input:
//   - a line containing stack IDs
//   - all container row lines from bottom to top (i.e. containers on the ground first)
func verify_stack_ids_and_container_rows(
	stack_id_line string,
	container_row_lines []string,
) ([]models.Containers, error) {

	// Extract stack ids: make a mapping from position (on the line) of stack ID to the stack ID
	stack_ids := f.CreateKeyValueMap(
		regexp.MustCompile(`\d+`).FindAllStringIndex(stack_id_line, -1),
		func(index []int) int { return index[0] },
		func(index []int) int { i, _ := strconv.Atoi(stack_id_line[index[0]:index[1]]); return i },
	)

	// Extract container stacks: make a mapping from a container position to container label (letter)
	container_letter_re := regexp.MustCompile(`[A-Z]`)
	container_rows := f.Map(func(row string) map[int]string {
		return f.CreateKeyValueMap(
			container_letter_re.FindAllStringIndex(row, -1),
			func(index []int) int { return index[0] },
			func(index []int) string { return row[index[0]:index[1]] },
		)
	}, container_row_lines)

	// 1. Verify that stack ids are unique values from 1 to n (where n is the number of stack ids)
	if !f.ArrayEqualInAnyOrder(f.GetValues(stack_ids), f.RangeInclusive(1, len(stack_ids))) {
		return nil, invalid_stack_ids_error(len(stack_ids))
	}

	// 2. Verify that all the containers are aligned with stack IDs
	if allowed_positions := f.CreateSet(f.GetKeys(stack_ids), f.Identity[int]); !f.All(
		func(position int) bool { return allowed_positions[position] },
		f.FlatMap(f.GetKeys[int, string], container_rows),
	) {
		return nil, containers_not_aligned_with_stack_ids_error()
	}

	// 3. Verify that there are no floating containers
	if floating_stacks_positions := f.Filter(
		func(stack_id_position int) bool {
			floating_container_detected, empty_space_below := false, false
			for _, container_row := range container_rows {
				_, container_present := container_row[stack_id_position]
				if empty_space_below && container_present {
					floating_container_detected = true
				}
				empty_space_below = !container_present
			}
			return floating_container_detected
		},
		f.GetKeys(stack_ids),
	); len(floating_stacks_positions) > 0 {
		return nil, floating_containers_error(f.Map(
			func(floating_stack_position int) int { return stack_ids[floating_stack_position] },
			floating_stacks_positions,
		))
	}

	container_stacks := make([]models.Containers, len(stack_ids))
	for stack_position, stack_id := range stack_ids {
		container_stack := make([]string, 0)
		for _, row := range container_rows {
			if container, ok := row[stack_position]; ok {
				container_stack = append(container_stack, container)
			}
		}
		container_stacks[stack_id-1] = strings.Join(container_stack, "")
	}

	// Verification completed; all ok from here

	return container_stacks, nil
}
