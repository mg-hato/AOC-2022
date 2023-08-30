package models

import "fmt"

func unfit_field_count_cube_field_map_creation_error(field_count int) error {
	return fmt.Errorf("cannot create a cube field map out of %d field(s)", field_count)
}
