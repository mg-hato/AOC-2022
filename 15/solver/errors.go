package solver

import "fmt"

func solver_error_prefix() string {
	return "solver error occurred"
}

func too_many_distress_beacons_candidates_error() error {
	return fmt.Errorf("%s: too many candidates for distress beacons", solver_error_prefix())
}

func no_distress_beacon_found_error() error {
	return fmt.Errorf("%s: no distress beacon has been found", solver_error_prefix())
}
