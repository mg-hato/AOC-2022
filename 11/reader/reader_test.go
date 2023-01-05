package reader

import (
	m "aoc/day11/models"
	"aoc/reading"
	"aoc/testers"
	"testing"
)

func TestD11_ReaderTest(t *testing.T) {

	items := func(i ...int) []int { return i }

	testers.DefaultReaderTester(reading.ReadWith(MonkeyGraphReader), "MonkeyGraphReader").
		AddGoodInputTest(m.CreateMonkeysEnvelopeWith([]m.Monkey{
			{
				MonkeyId:     0,
				Items:        items(79, 98),
				InspectionOP: m.IOP(m.Old(), m.Mult(), m.Num(19)),
				DivTest:      23,
				OnTrue:       2,
				OnFalse:      3,
			},
			{
				MonkeyId:     1,
				Items:        items(54, 65, 75, 74),
				InspectionOP: m.IOP(m.Old(), m.Add(), m.Num(6)),
				DivTest:      19,
				OnTrue:       2,
				OnFalse:      0,
			},
			{
				MonkeyId:     2,
				Items:        items(79, 60, 97),
				InspectionOP: m.IOP(m.Old(), m.Mult(), m.Old()),
				DivTest:      13,
				OnTrue:       1,
				OnFalse:      3,
			},
			{
				MonkeyId:     3,
				Items:        items(74),
				InspectionOP: m.IOP(m.Old(), m.Add(), m.Num(3)),
				DivTest:      17,
				OnTrue:       0,
				OnFalse:      1,
			},
		})).
		AddGoodInputTest(m.CreateMonkeysEnvelopeWith([]m.Monkey{
			{
				MonkeyId:     0,
				Items:        items(),
				InspectionOP: m.IOP(m.Old(), m.Add(), m.Num(100)),
				DivTest:      7,
				OnTrue:       1,
				OnFalse:      1,
			},
			{
				MonkeyId:     1,
				Items:        items(1),
				InspectionOP: m.IOP(m.Old(), m.Mult(), m.Num(2)),
				DivTest:      5,
				OnTrue:       0,
				OnFalse:      0,
			},
		})).
		ProvideEqualityFunctionForTypeT(m.MonkeyEnvelopeEqFunc).
		AddErrorInputTest("Divisibility test with 0").
		AddErrorInputTest("Self-loop on monkey 0").
		AddErrorInputTest("Monkey 1 throws out of bounds in false-case").
		AddErrorInputTest("Monkey 1 throws out of bounds in true-case").
		AddErrorInputTest("Bad ID order").
		RunBothGoodAndErrorInputTests(t)
}
