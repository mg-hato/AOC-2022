package solver

import (
	c "aoc/common"
	m "aoc/d11/models"
	ts "aoc/testers"
	"testing"
)

func TestD11_Solver(t *testing.T) {
	ts.SolverTesterForComparableResults[m.SolverInput, int64](t).
		ProvideSolver(CalculateMonkeyBusiness(20, DivBy3)).
		ProvideSolver(CalculateMonkeyBusiness(10_000, NoAdjustment)).
		AddTestCase(m.CreateMonkeysEnvelopeWith([]m.Monkey{
			// All monkeys will preserve divisibility outcome with 2 and 7 before and after their operations (and adjustment)
			// Monkey 0 is sink i.e. by the end of each round it will have all the items thrown to it
			// Monkey 1 will receive only even items
			// Monkey 2 will receive odd items & even items that are not divisible by 7
			{
				MonkeyId:     0, // 6 items per round: all
				Items:        []int{21, 37, 15, 49, 70, 16},
				InspectionOP: m.IOP(m.Old(), m.Mult(), m.Num(3)),
				DivTest:      2,
				OnTrue:       1,
				OnFalse:      2,
			},
			{
				MonkeyId:     1, // 2 items per round: 70, 16
				Items:        []int{},
				InspectionOP: m.IOP(m.Old(), m.Mult(), m.Num(3)),
				DivTest:      7,
				OnTrue:       0,
				OnFalse:      2,
			},
			{
				MonkeyId:     2, // 5 items per round: 21, 37, 15, 49, 16
				Items:        []int{},
				InspectionOP: m.IOP(m.Old(), m.Mult(), m.Num(3)),
				DivTest:      5,
				OnTrue:       0,
				OnFalse:      0,
			},
			// To visualise this Monkey Graph, see "monkey-graph-1.jpg"
		}),
			ts.ExpectResult[int64](6*5*20*20),
			ts.ExpectResult[int64](6*5*10_000*10_000),
		).
		AddTestCase(m.CreateMonkeysEnvelopeWith([]m.Monkey{
			{
				// Monkey 0 receives everything at the end of the round
				MonkeyId:     0, // 101 items per round: all
				Items:        c.RangeInclusive(0, 100),
				InspectionOP: m.IOP(m.Old(), m.Mult(), m.Num(3)),
				DivTest:      2,
				OnTrue:       1,
				OnFalse:      2,
			},
			{
				// Monkey 1 receives only items divisible with 2
				MonkeyId:     1, // 51 items per round: 0, 2, ... 98, 100
				Items:        []int{},
				InspectionOP: m.IOP(m.Old(), m.Mult(), m.Num(3)),
				DivTest:      5,
				OnTrue:       3,
				OnFalse:      4,
			},
			{
				// Monkey 2 receives only items that are not divisible with 2
				MonkeyId:     2, // 50 items per round: 1, 3, ... , 97, 99
				Items:        []int{},
				InspectionOP: m.IOP(m.Old(), m.Mult(), m.Num(3)),
				DivTest:      13,
				OnTrue:       0,
				OnFalse:      4,
			},
			{
				// Monkey 3 receives items that are divisible by 10
				MonkeyId:     3, // 11 items per round: 0, 10, 20, ... 100
				Items:        []int{},
				InspectionOP: m.IOP(m.Old(), m.Mult(), m.Num(3)),
				DivTest:      11,
				OnTrue:       5,
				OnFalse:      4,
			},
			{
				// Monkey 4 receives:
				// - even items that are not divisible with 5 (40 = 51 even items - 11 items div by 10)
				// - odd items that are not divisible with 13 (46 items = 50 odd items - 4 odd items not div by 13: [13, 39, 65, 91])
				// - items divisible with 10, but not with 11 (10 items: 10, 20, ..., 100)
				MonkeyId:     4, // 96 items per round: all
				Items:        []int{},
				InspectionOP: m.IOP(m.Old(), m.Mult(), m.Num(3)),
				DivTest:      7,
				OnTrue:       5,
				OnFalse:      0,
			},
			{
				// Monkey 5 receives:
				// - items divisible with 110 (only 1: it is item 0)
				// - items from monkey 4 that are divisible by 7 (a lot less than 96)
				// Whatever Monkey 4 has only the items divisible by 7 will be passed over to Monkey 5 + Item 0 (in total less than 96)
				// As such the two "most active" monkeys will be Monkey 0 & Monkey 4, formulating the "Monkey Business Value"
				// Hence we do not have to work out how many items exactly will Monkey 5 analyse per round
				MonkeyId:     5, // less than 97 items per round.
				Items:        []int{},
				InspectionOP: m.IOP(m.Old(), m.Mult(), m.Num(3)),
				DivTest:      1,
				OnTrue:       0,
				OnFalse:      0,
			},
			// To visualise this Monkey Graph, see "monkey-graph-2.jpg"
		}),
			ts.ExpectResult[int64](101*96*20*20),
			ts.ExpectResult[int64](101*96*10_000*10_000),
		).
		RunSolverTests()
}
