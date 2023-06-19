package reader

import (
	m "aoc/d15/models"
	"aoc/reading"
	ts "aoc/testers"
	"fmt"
	"testing"
)

func TestD15_ReaderTest(t *testing.T) {
	ts.ReaderTester(t, reading.ReadWith(SensorReportsReader)).
		ProvideEqualityFunction(m.SensorReportsEnvelopeEqualityFunction).
		AddTestCase("./tests/input-1.txt", ts.ExpectResult(m.SensorReportsEnvelope(
			m.SensorReport{Sensor: m.MakePoint(300, 300), Beacon: m.MakePoint(1_000, 1_000)},
			m.SensorReport{Sensor: m.MakePoint(69_000, 55_000), Beacon: m.MakePoint(70_000, 25_000)},
			m.SensorReport{Sensor: m.MakePoint(-5, -10), Beacon: m.MakePoint(-850, -650)},
			m.SensorReport{Sensor: m.MakePoint(1_234_567, 2_125_000), Beacon: m.MakePoint(999_800, 1_555_444)},
		))).
		AddTestCase("./tests/overlapping_sensors.txt", ts.ExpectError[m.SolverInput](
			m.FormatPoint(m.MakePoint(300, 300)), "sensor #1", "sensor #4", "final validation",
		)).
		AddTestCase("./tests/closer_beacon.txt", ts.ExpectError[m.SolverInput](
			"final validation",
			fmt.Sprintf("sensor at %s", m.FormatPoint(m.MakePoint(300, 300))),
			"closer beacon",
			"reported beacon",
			m.FormatPoint(m.MakePoint(500, 100)),
			m.FormatPoint(m.MakePoint(1_000, 1_000)),
		)).
		AddTestCase("./tests/equidistant_beacons.txt", ts.ExpectError[m.SolverInput](
			"final validation",
			fmt.Sprintf("sensor at %s", m.FormatPoint(m.MakePoint(300, 500))),
			"equidistant beacons",
			m.FormatPoint(m.MakePoint(300, 400)),
			m.FormatPoint(m.MakePoint(301, 599)),
		)).
		RunReaderTests()
}
