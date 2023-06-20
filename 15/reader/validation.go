package reader

import (
	m "aoc/d15/models"
	f "aoc/functional"
)

func verify_that_no_two_sensors_overlap(reports []m.SensorReport) error {
	sensors := make(map[m.Point]int)
	for i, report := range reports {
		if sensor_number, overlap := sensors[report.Sensor]; overlap {
			return two_sensors_overlap_validation_error(report.Sensor, sensor_number, i+1)
		} else {
			sensors[report.Sensor] = i + 1
		}
	}
	return nil
}

func verify_beacons_are_the_closest(reports []m.SensorReport) error {
	beacons := f.Map(func(report m.SensorReport) m.Point { return report.Beacon }, reports)
	for _, report := range reports {
		beacon_distance := m.Distance(report.Sensor, report.Beacon)
		closer_beacons := f.Filter(
			func(beacon m.Point) bool {
				return m.Distance(report.Sensor, beacon) < beacon_distance
			},
			beacons,
		)
		if len(closer_beacons) > 0 {
			return closer_beacon_found_validation_error(report, closer_beacons[0])
		}
	}
	return nil
}

func verify_no_sensor_has_multiple_equidistant_beacons(reports []m.SensorReport) error {
	beacons := f.Map(func(report m.SensorReport) m.Point { return report.Beacon }, reports)
	for _, report := range reports {
		beacon_distance := m.Distance(report.Sensor, report.Beacon)
		equidistant_beacons := f.CreateSet(f.Filter(
			func(beacon m.Point) bool { return m.Distance(report.Sensor, beacon) == beacon_distance },
			beacons,
		), f.Identity[m.Point])
		if len(equidistant_beacons) > 1 {
			sample_beacons := f.GetKeys(equidistant_beacons)
			return equidistant_beacons_validation_error(report.Sensor, sample_beacons[0], sample_beacons[1])
		}
	}
	return nil
}
