package reader

import (
	m "aoc/d05/models"
	"aoc/reading"
	ts "aoc/testers"
	"testing"
)

func TestD05_ReaderTest(t *testing.T) {

	ts.ReaderTester(t, reading.ReadWith(MovingPlanReader)).
		ProvideEqualityFunction(m.MovingPlanEnvelopeEqualityFunction).
		AddTestCase("./tests/input-1.txt", ts.ExpectResult(m.CreateMovingPlanEnvelope(m.MovingPlan{
			StartingContainers: []string{"A", "XPI", ""},
			Moves: []m.Move{
				m.MakeMove(1, 1, 3),
				m.MakeMove(2, 2, 1),
			},
		}))).
		AddTestCase("./tests/input-2.txt", ts.ExpectResult(m.CreateMovingPlanEnvelope(m.MovingPlan{
			StartingContainers: []string{"J", "T", "ER", "APZ"},
			Moves: []m.Move{
				m.MakeMove(2, 4, 1),
				m.MakeMove(1, 3, 2),
				m.MakeMove(1, 1, 3),
			},
		}))).
		AddTestCase("./tests/bad-input-1.txt", ts.ExpectError[m.SolverInput]("float", "detected", "1", "2", "3", "4")).
		AddTestCase("./tests/bad-input-2.txt", ts.ExpectError[m.SolverInput]("float", "detected", "4")).
		AddTestCase("./tests/bad-input-3.txt", ts.ExpectError[m.SolverInput]("not aligned")).
		AddTestCase("./tests/bad-input-4.txt", ts.ExpectError[m.SolverInput]("stack ID", "unique integer", "[1, 4]")).
		AddTestCase("./tests/bad-input-5.txt", ts.ExpectError[m.SolverInput]("source", "destination", "[1, 4]", "line #8")).
		AddTestCase("./tests/bad-input-6.txt", ts.ExpectError[m.SolverInput]("source", "destination", "[1, 4]", "line #9")).
		AddTestCase("./tests/bad-input-7.txt", ts.ExpectError[m.SolverInput]("same", "source", "destination", "line #10")).
		AddTestCase("./tests/bad-input-8.txt", ts.ExpectError[m.SolverInput]("mov", "instruction", "interpret", "line #7", "bad-line")).
		AddTestCase("./tests/bad-input-9.txt", ts.ExpectError[m.SolverInput]("container", "interpret", "line #3", "bad-container-line")).
		RunReaderTests()
}
