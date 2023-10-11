package solver

import m "aoc/d22/models"

func GetFinalCoordinates(field_map_constructor func([]m.Field) (m.FieldMap, error)) func(input m.SolverInput) (int, error) {
	return func(input m.SolverInput) (int, error) {

		fields, instructions := input.Get().Get()
		field_map, err := field_map_constructor(fields)
		if err != nil {
			return 0, err
		}

		pointer := field_map.GetInitialPointer()
		for _, instruction := range instructions {
			pointer = field_map.UpdatePointer(pointer, instruction)
		}

		return pointer.GetPassword(), nil
	}
}
