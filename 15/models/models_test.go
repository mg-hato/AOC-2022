package models

import (
	f "aoc/functional"
	ts "aoc/testers"
	"testing"
)

func TestD15_Distance(t *testing.T) {
	ts.AssertEqual(t, Distance(MakePoint(3, 4), MakePoint(5, 5)), 3)
	ts.AssertEqual(t, Distance(MakePoint(5, 5), MakePoint(3, 4)), 3)
	ts.AssertEqual(t, Distance(MakePoint(-10, 20), MakePoint(10, 20)), 20)
	ts.AssertEqual(t, Distance(MakePoint(-5, 0), MakePoint(0, 25)), 30)
}

func TestD15_Envelope(t *testing.T) {
	data := func() []SensorReport {
		return []SensorReport{
			{Sensor: MakePoint(10, 20), Beacon: MakePoint(9, 22)},
			{Sensor: MakePoint(100, 200), Beacon: MakePoint(94, 190)},
			{Sensor: MakePoint(-3, 0), Beacon: MakePoint(0, 0)},
		}
	}
	envelope := SensorReportsEnvelope(data()...)
	mutable_reports := envelope.Get()
	mutable_reports[0].Sensor.First = 12
	mutable_reports[2] = SensorReport{}

	ts.AssertEqualWithEqFunc(t, envelope.Get(), data(), f.ArrayEqual[SensorReport])
}
