package solver

import (
	m "aoc/d21/models"
)

func SolveRoot(envelope m.SolverInput) (int64, error) {
	return create_cached_solver(make_graph(envelope.Get()))("root")
}
