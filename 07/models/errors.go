package models

import "fmt"

func directory_content_is_unknown_error(from, to string) error {
	return fmt.Errorf(`cannot change directory from %s to %s as there has been no ls command to list its contents`,
		from, to,
	)
}

func directory_does_not_exist_error(from, to string) error {
	return fmt.Errorf(`cannot change directory from %s to %s as it does not exist`, from, to)
}

func not_a_directory_error(from, to string) error {
	return fmt.Errorf(`cannot change directory from %s to %s as it is not a directory`, from, to)
}

func ls_items_do_not_match_error(directory_name string) error {
	return fmt.Errorf(
		`items in ls command do not match with previously given listing for directory %s`,
		directory_name,
	)
}

func directory_is_unexplored_error(name string) error {
	return fmt.Errorf(
		`no ls-listing was given for directory %s`, name,
	)
}

func create_filesystem_error(command_number int, command Command, underlying_error error) error {
	return fmt.Errorf(
		`error while creating filesystem on command #%d "%s": %s`,
		command_number, command, underlying_error.Error(),
	)
}

func create_filesystem_verification_error(underlying_error error) error {
	return fmt.Errorf(
		`verification error while creating filesystem: %s`, underlying_error.Error(),
	)
}
