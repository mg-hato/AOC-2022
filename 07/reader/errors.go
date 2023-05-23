package reader

import "fmt"

func bad_first_command(line_number int, line string) error {
	return fmt.Errorf(
		`error reading first command on line #%d: first command should be "$ cd /" but the actual line is: "%s"`,
		line_number, line,
	)
}

func bad_new_command(line_number int, line string) error {
	return fmt.Errorf(
		`error reading a command on line #%d: "%s"`,
		line_number, line,
	)
}

func bad_item_ls_command(line_number int, line string) error {
	return fmt.Errorf(
		`errror while reading ls command listing on line #%d: "%s"`,
		line_number, line,
	)
}

func duplicated_name_in_ls_items_listing(line_number int, duplicated_name string) error {
	return fmt.Errorf(
		`error while reading ls command listing on line #%d: duplicate item name detected "%s"`,
		line_number, duplicated_name,
	)
}
