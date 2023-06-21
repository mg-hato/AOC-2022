package solver

import (
	c "aoc/common"
	m "aoc/d15/models"
)

func BeaconExclusionCount(y int) func(m.SolverInput) (int64, error) {
	return func(input m.SolverInput) (int64, error) {

		reports := input.Get()
		composite_exclusion_range := compositeExclusionRangeForFixedY(y, reports)

		// The number of beacons on the fixed Y coordinate that belong to the exclusion ranges
		beacons_in_exclusion_ranges := c.Count(
			c.GetKeys(c.CreateSet(
				c.Filter(func(report m.SensorReport) bool { return report.Beacon.Second == y }, reports),
				func(report m.SensorReport) m.Point { return report.Beacon },
			)),
			func(beacon m.Point) bool { return composite_exclusion_range.includes(beacon.First) },
		)

		return int64(composite_exclusion_range.getCoverageCount() - beacons_in_exclusion_ranges), nil
	}
}
