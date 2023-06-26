package reader

import "fmt"

func bad_line(line_number int, line string) error {
	return fmt.Errorf(
		`error while reading line #%d: could not interpret line "%s"`,
		line_number, line,
	)
}
