package solver

import (
	m "aoc/d15/models"
	e "aoc/envelope"
	f "aoc/functional"
)

func BeaconExclusionCount(y int) func(e.Envelope[[]m.SensorReport]) (int64, error) {
	return func(envelope e.Envelope[[]m.SensorReport]) (int64, error) {

		reports := envelope.Get()
		composite_exclusion_range := compositeExclusionRangeForFixedY(y, reports)

		// The number of beacons on the fixed Y coordinate that belong to the exclusion ranges
		beacons_in_exclusion_ranges := f.Count(
			f.GetKeys(f.CreateSet(
				f.Filter(func(report m.SensorReport) bool { return report.Beacon.Second == y }, reports),
				func(report m.SensorReport) m.Point { return report.Beacon },
			)),
			func(beacon m.Point) bool { return composite_exclusion_range.includes(beacon.First) },
		)

		return int64(composite_exclusion_range.getCoverageCount() - beacons_in_exclusion_ranges), nil
	}
}
