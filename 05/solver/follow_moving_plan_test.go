package solver

import (
	m "aoc/day05/models"
	e "aoc/envelope"
	ts "aoc/testers"
	"testing"
)

func TestD05_SolverTest(t *testing.T) {
	type InputData = e.Envelope[m.MovingPlan]
	ts.SolverTesterForComparableResults[InputData, string](t).
		ProvideSolver(FollowMovingPlanWith(CrateMover9000())).
		ProvideSolver(FollowMovingPlanWith(CrateMover9001())).
		AddTestCase(m.CreateMovingPlanEnvelope(m.MovingPlan{
			StartingContainers: []string{
				"ABC",
				"",
				"DEF",
				"XYZPQRST",
			},
			Moves: []m.Move{
				m.MakeMove(5, 4, 2),
				m.MakeMove(2, 3, 1),
				m.MakeMove(1, 3, 1),
				m.MakeMove(3, 4, 3),
				m.MakeMove(1, 1, 4),
			},
		}),
			ts.ExpectResult("EPXD"),
			ts.ExpectResult("FTZD"),
		).
		AddTestCase(m.CreateMovingPlanEnvelope(m.MovingPlan{
			StartingContainers: []string{
				"ABC",
				"",
				"DEF",
			},
			Moves: []m.Move{
				m.MakeMove(1, 3, 1),
				m.MakeMove(3, 1, 3),
			},
		}),
			ts.ExpectError[string]("stacks have no containers", "[2]"),
			ts.ExpectError[string]("stacks have no containers", "[2]"),
		).
		AddTestCase(m.CreateMovingPlanEnvelope(m.MovingPlan{
			StartingContainers: []string{
				"ABC",
				"",
				"DEF",
			},
			Moves: []m.Move{
				m.MakeMove(1, 3, 1),
				m.MakeMove(3, 1, 3),
				m.MakeMove(6, 3, 2),
			},
		}),
			ts.ExpectError[string]("move #3", "6", "3", "2", "source stack", "not", "enough"),
			ts.ExpectError[string]("move #3", "6", "3", "2", "source stack", "not", "enough"),
		).
		RunSolverTests()
}
