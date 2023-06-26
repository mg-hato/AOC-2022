package solver

import (
	m "aoc/d17/models"
	"os"
)

func GetHeightAfterNRocks(n int64) func(string) (int64, error) {
	return func(jet_pattern string) (int64, error) {
		repeating_jet_pattern := m.RepeatingJetPattern(jet_pattern)
		chamber := m.EmptyChamber(7)

		shape_sequence := []func() m.Shape{
			m.HorizontalLine,
			m.PlusShape,
			m.L_Shape,
			m.VerticalLine,
			m.Square,
		}

		var i int64 = 0

		for i < n {

			rock := m.GenerateShapeAt(
				chamber.GetHeight()+4,
				2,
				shape_sequence[i%int64(len(shape_sequence))],
			)

			rock_has_stopped := false
			for !rock_has_stopped {
				// rock is being pushed by the jet of hot gas
				update_func, jet_index := repeating_jet_pattern.Next()
				updated_rock_location := update_func(rock)
				if chamber.ShapeCanBeMovedTo(updated_rock_location) {
					rock = updated_rock_location
				}

				updated_rock_location = m.MoveDown(rock)
				if chamber.ShapeCanBeMovedTo(updated_rock_location) {
					rock = updated_rock_location
				} else {
					rock_has_stopped = true
					optimisation_function := chamber.AddShape(rock, jet_index, int(i%int64(len(shape_sequence))), i)
					if optimisation_function != nil {
						i = optimisation_function(i, n)
					}
				}
			}

			i++
		}
		// fmt.Println("Closure[3] hits count: ", chamber.GetClosureHits())
		os.WriteFile("chamber.txt", []byte(chamber.String()), 0644)
		return chamber.GetTotalHeight(), nil
	}
}
