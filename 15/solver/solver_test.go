package solver

import (
	m "aoc/d15/models"
	ts "aoc/testers"
	"testing"
)

func TestD15_CompositeExclusionRange(t *testing.T) {
	reports := []m.SensorReport{
		{Sensor: m.MakePoint(2, 18), Beacon: m.MakePoint(-2, 15)},
		{Sensor: m.MakePoint(9, 16), Beacon: m.MakePoint(10, 16)},
		{Sensor: m.MakePoint(13, 2), Beacon: m.MakePoint(15, 3)},
		{Sensor: m.MakePoint(12, 14), Beacon: m.MakePoint(10, 16)},
		{Sensor: m.MakePoint(10, 20), Beacon: m.MakePoint(10, 16)},
		{Sensor: m.MakePoint(14, 17), Beacon: m.MakePoint(10, 16)},
		{Sensor: m.MakePoint(8, 7), Beacon: m.MakePoint(2, 10)},
		{Sensor: m.MakePoint(2, 0), Beacon: m.MakePoint(2, 10)},
		{Sensor: m.MakePoint(0, 11), Beacon: m.MakePoint(2, 10)},
		{Sensor: m.MakePoint(20, 14), Beacon: m.MakePoint(25, 17)},
		{Sensor: m.MakePoint(17, 20), Beacon: m.MakePoint(21, 22)},
		{Sensor: m.MakePoint(16, 7), Beacon: m.MakePoint(15, 3)},
		{Sensor: m.MakePoint(14, 3), Beacon: m.MakePoint(15, 3)},
		{Sensor: m.MakePoint(20, 1), Beacon: m.MakePoint(15, 3)},
	}

	c09 := compositeExclusionRangeForFixedY(9, reports)
	ts.AssertEqual(t, c09.includes(-2), false)
	ts.AssertEqual(t, c09.includes(-1), true)
	ts.AssertEqual(t, c09.getFirstNonNegativeNumberNotInCompositeRange(), 24)

	c10 := compositeExclusionRangeForFixedY(10, reports)
	ts.AssertEqual(t, c10.getCoverageCount(), 27)
	ts.AssertEqual(t, c10.getCoverageCountBetween(0, 20), 21)
	ts.AssertEqual(t, c10.includes(25), false)
	ts.AssertEqual(t, c10.getFirstNonNegativeNumberNotInCompositeRange(), 25)

	c11 := compositeExclusionRangeForFixedY(11, reports)
	ts.AssertEqual(t, c11.getFirstNonNegativeNumberNotInCompositeRange(), 14)
}
