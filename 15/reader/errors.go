package reader

import (
	m "aoc/d15/models"
	"fmt"
)

func validation_error_prefix() string {
	return "error while performing reader final validation"
}

func bad_line_reader_error(line string, line_number int) error {
	return fmt.Errorf(`error while reading on line #%d: could not interpret the line "%s"`, line_number, line)
}

func two_sensors_overlap_validation_error(overlapping_point m.Point, sensor_number_1, sensor_number_2 int) error {
	return fmt.Errorf(
		`%s: there are two sensors at point %s, namely sensor #%d and sensor #%d`,
		validation_error_prefix(),
		m.FormatPoint(overlapping_point),
		sensor_number_1,
		sensor_number_2,
	)
}

func closer_beacon_found_validation_error(report m.SensorReport, closer_beacon m.Point) error {
	return fmt.Errorf(
		`%s: a closer beacon has been detected at %s for sensor at %s while its reported beacon is at %s`,
		validation_error_prefix(),
		m.FormatPoint(closer_beacon),
		m.FormatPoint(report.Sensor),
		m.FormatPoint(report.Beacon),
	)
}

func equidistant_beacons_validation_error(sensor, beacon_1, beacon_2 m.Point) error {
	return fmt.Errorf(
		`%s: sensor at %s has two equidistant beacons, namely beacons at %s and %s`,
		validation_error_prefix(),
		m.FormatPoint(sensor),
		m.FormatPoint(beacon_1),
		m.FormatPoint(beacon_2),
	)
}
