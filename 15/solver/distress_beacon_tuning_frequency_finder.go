package solver

import (
	c "aoc/common"
	m "aoc/d15/models"
)

func DistressBeaconTuningFrequencyFinder(limit int) func(m.SolverInput) (int64, error) {
	return func(input m.SolverInput) (int64, error) {

		reports := input.Get()
		distress_beacons_count := 0
		distress_beacon := c.Pair[int, int]{}
		for y := 0; y <= limit; y++ {

			composite_exclusion_range := compositeExclusionRangeForFixedY(y, reports)
			count := limit + 1 - composite_exclusion_range.getCoverageCountBetween(0, limit)
			distress_beacons_count += count

			if distress_beacons_count > 1 {
				return -1, too_many_distress_beacons_candidates_error()
			}

			if count == 1 {
				distress_beacon = m.MakePoint(
					composite_exclusion_range.getFirstNonNegativeNumberNotInCompositeRange(),
					y,
				)
			}
		}
		if distress_beacons_count == 0 {
			return -1, no_distress_beacon_found_error()
		}
		return int64(distress_beacon.First)*(int64(limit)) + int64(distress_beacon.Second), nil
	}
}
