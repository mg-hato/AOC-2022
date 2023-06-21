package solver

import (
	m "aoc/d02/models"
	ts "aoc/testers"
	"testing"
)

func TestD02_SolverTest(t *testing.T) {

	// helper function
	r := func(left m.LeftSymbol, right m.RightSymbol) m.Round { return m.Round{Left: left, Right: right} }

	// shortcut for ABC & XYZ symbols
	x, y, z := m.X, m.Y, m.Z
	a, b, c := m.A, m.B, m.C

	ts.SolverTesterForComparableResults[m.SolverInput, int](t).
		ProvideSolver(CalculateScore(ShapeBasedRoundInterpreter())).
		ProvideSolver(CalculateScore(OutcomeBasedRoundInterpreter())).
		AddTestCase(
			m.CreateRoundsEnvelope([]m.Round{
				r(a, x), // RR: 3+1 | RL: 0+3
				r(b, z), // PS: 6+3 | PW: 6+3
				r(a, z), // RS: 0+3 | RW: 6+2
			}),
			ts.ExpectResult(16),
			ts.ExpectResult(20),
		).
		AddTestCase(
			m.CreateRoundsEnvelope([]m.Round{
				r(c, y), // SP: 0+2 | SD: 3+3
				r(c, z), // SS: 3+3 | SW: 6+1
				r(a, y), // RP: 6+2 | RD: 3+1
			}),
			ts.ExpectResult(16),
			ts.ExpectResult(17),
		).
		RunSolverTests()

}
