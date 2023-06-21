package solver

import (
	c "aoc/common"
	m "aoc/d15/models"
	"sort"
)

// Represents a collection of integer ranges
type composite_range struct {
	ranges []c.Pair[int, int]
}

// Returns a list of x-coordinate exclusion ranges where no distress beacon can appear on the provided/fixed y coordinate
func compositeExclusionRangeForFixedY(y int, reports []m.SensorReport) composite_range {
	exclusion_ranges := make([]c.Pair[int, int], 0)

	// From the reports produce distress beacon x-axis-exclusion ranges for fixed y
	for i := 0; i < len(reports); i++ {

		axis_distance := m.Abs(y - reports[i].Sensor.Second)
		beacon_distance := m.Distance(reports[i].Sensor, reports[i].Beacon)

		if axis_distance <= beacon_distance {
			exclusion_ranges = append(exclusion_ranges, c.MakePair(
				reports[i].Sensor.First-(beacon_distance-axis_distance),
				reports[i].Sensor.First+(beacon_distance-axis_distance),
			))
		}
	}

	// Sort exclusion ranges by the left-component (i.e. First)
	sort.Slice(exclusion_ranges, func(i, j int) bool {
		return exclusion_ranges[i].First < exclusion_ranges[j].First
	})

	merged_ranges := make([]c.Pair[int, int], 0)
	if len(exclusion_ranges) > 0 {
		merged_ranges = append(merged_ranges, exclusion_ranges[0])
	}

	// Merge and simplify ranges
	// e.g. two ranges [1,5] and [3,7] can be represented by a single range [1,7]
	for i := 1; i < len(exclusion_ranges); i++ {
		last := len(merged_ranges) - 1
		if exclusion_ranges[i].First <= merged_ranges[last].Second {
			if exclusion_ranges[i].Second > merged_ranges[last].Second {
				merged_ranges[last].Second = exclusion_ranges[i].Second
			}
		} else {
			merged_ranges = append(merged_ranges, exclusion_ranges[i])
		}
	}
	return composite_range{merged_ranges}
}

// Get the number of different integers belonging to this composite range
func (cr composite_range) getCoverageCount() int {
	return c.Sum(c.Map(
		func(r c.Pair[int, int]) int { return r.Second - r.First + 1 },
		cr.ranges,
	))
}

// Returns true if and only if x belongs to this composite range
func (cr composite_range) includes(x int) bool {
	return c.Any(
		func(r c.Pair[int, int]) bool { return c.InInclusiveRange(r.First, r.Second)(x) },
		cr.ranges,
	)
}

// Get number of different integers belonging to this composite range that are in [left, right]
func (cr composite_range) getCoverageCountBetween(left, right int) int {
	count := 0
	for _, r := range cr.ranges {
		left_border := c.Max(left, r.First)
		right_border := c.Min(right, r.Second)
		if left_border <= right_border {
			count += right_border - left_border + 1
		}
	}
	return count
}

// Get first non-negative number that is not a member of this composite range
func (cr composite_range) getFirstNonNegativeNumberNotInCompositeRange() int {
	current := 0
	for _, r := range cr.ranges {
		if c.InInclusiveRange(r.First, r.Second)(current) {
			current = r.Second + 1
		}
	}
	return current
}
