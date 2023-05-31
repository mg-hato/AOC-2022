package solver

import (
	m "aoc/day09/models"
	e "aoc/envelope"
	f "aoc/functional"
)

func CountPositionsVisitedByLastKnot(number_of_knots int) func(e.Envelope[m.MotionSeries]) (int, error) {
	return func(envelope e.Envelope[m.MotionSeries]) (int, error) {
		var knots []m.Position = f.Repeat(f.MakePair(0, 0), number_of_knots)

		visited := map[m.Position]bool{f.MakePair(0, 0): true}
		for _, motion := range envelope.Get() {
			for i := 0; i < motion.Steps; i++ {
				move_knots(motion.Direction, knots)
				visited[knots[number_of_knots-1]] = true
			}
		}

		return f.Count(f.GetValues(visited), f.Identity[bool]), nil
	}
}
