package reader

import (
	m "aoc/d22/models"
	"aoc/reading"
	ts "aoc/testers"
	"testing"
)

func TestD22_ReaderTest(t *testing.T) {
	ts.ReaderTester(t, reading.ReadWith(MonkeyMapReader)).
		ProvideEqualityFunction(m.FieldsAndInstructionsEnvelopeEqualityFunc).
		AddTestCase("./tests/example.txt", ts.ExpectResult(m.CreateFieldsAndInstructionsEnvelope(
			[]m.Field{
				{Row: 1, Column: 9, FType: m.Dot},
				{Row: 1, Column: 10, FType: m.Dot},
				{Row: 1, Column: 11, FType: m.Dot},
				{Row: 1, Column: 12, FType: m.Wall},

				{Row: 2, Column: 9, FType: m.Dot},
				{Row: 2, Column: 10, FType: m.Wall},
				{Row: 2, Column: 11, FType: m.Dot},
				{Row: 2, Column: 12, FType: m.Dot},

				{Row: 3, Column: 7, FType: m.Dot},
				{Row: 3, Column: 8, FType: m.Dot},
				{Row: 3, Column: 9, FType: m.Dot},
				{Row: 3, Column: 10, FType: m.Wall},
				{Row: 3, Column: 11, FType: m.Wall},
				{Row: 3, Column: 12, FType: m.Wall},
				{Row: 3, Column: 13, FType: m.Dot},
				{Row: 3, Column: 14, FType: m.Dot},
				{Row: 3, Column: 15, FType: m.Dot},
			},
			[]m.Instruction{
				m.CreateMoveInstruction(10),
				m.CreateTurnInstruction(m.Right),
				m.CreateMoveInstruction(5),
				m.CreateTurnInstruction(m.Left),
				m.CreateMoveInstruction(5),
				m.CreateTurnInstruction(m.Right),
				m.CreateMoveInstruction(10),
				m.CreateTurnInstruction(m.Left),
				m.CreateMoveInstruction(4),
				m.CreateTurnInstruction(m.Right),
				m.CreateMoveInstruction(5),
				m.CreateTurnInstruction(m.Left),
				m.CreateMoveInstruction(5),
			},
		))).
		AddTestCase("./tests/discontinuity_rows.txt", ts.ExpectError[m.SolverInput]("discountinuity", "row")).
		AddTestCase("./tests/discontinuity_columns.txt", ts.ExpectError[m.SolverInput]("discountinuity", "column")).
		RunReaderTests()
}
