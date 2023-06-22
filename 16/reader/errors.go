package reader

import "fmt"

func reader_error_prefix(line_number int) string {
	return fmt.Sprintf("error while reading line #%d", line_number)
}

func validation_error_prefix() string {
	return "error while performing reader's final validaiton"
}

func bad_line_error(line_number int, line string) error {
	return fmt.Errorf(
		`%s: could not interpret line "%s"`,
		reader_error_prefix(line_number),
		line,
	)
}

func starting_valve_missing_error() error {
	return fmt.Errorf(
		`%s: there is no valve with ID "AA" which is supposed to be the starting valve`,
		validation_error_prefix(),
	)
}

func valve_defined_twice_validation_error(
	valve_id string,
	first_definition_number,
	second_definition_number int,
) error {
	return fmt.Errorf(
		`%s: valve with name %s defined twice both as valve #%d and valve #%d`,
		validation_error_prefix(),
		valve_id,
		first_definition_number,
		second_definition_number,
	)
}

func valve_has_self_loop_validation_error(valve_id string) error {
	return fmt.Errorf("%s: valve %s leads to itself", validation_error_prefix(), valve_id)
}

func valve_has_two_tunnels_leading_to_the_same_valve(valve_id, tunnel string) error {
	return fmt.Errorf(
		"%s: valve %s has more than one tunnel leading to valve %s",
		validation_error_prefix(),
		valve_id, tunnel,
	)
}

func valve_has_a_tunnel_leading_to_undefined_valve_validation_error(valve_id, tunnel string) error {
	return fmt.Errorf(
		"%s: valve %s has a tunnel leading to undefined valve %s",
		validation_error_prefix(),
		valve_id,
		tunnel,
	)
}
