package argshandle

import "fmt"

func argument_parsing_error_prefix() string {
	return "error while reading the command line arguments"
}

func invalid_argument_error(arg string) error {
	return fmt.Errorf(
		`%s: unknown argument was given "%s". Use --h flag to see help`,
		argument_parsing_error_prefix(),
		arg,
	)
}

func argument_option_used_multiple_times_error(usage string) error {
	return fmt.Errorf(
		`%s: argument option for %s was used multiple times`,
		argument_parsing_error_prefix(),
		usage,
	)
}

func argument_option_filename_not_provided_error() error {
	return fmt.Errorf(
		`%s: argument option for selecting the input file was used but no filename was provided`,
		argument_parsing_error_prefix(),
	)
}

func input_file_not_provided_error() error {
	return fmt.Errorf(
		`%s: no input file was given`,
		argument_parsing_error_prefix(),
	)
}
