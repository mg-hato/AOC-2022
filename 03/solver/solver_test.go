package solver

import (
	m "aoc/day03/models"
	"aoc/envelope"
	f "aoc/functional"
	ts "aoc/testers"
	"testing"
)

func TestD03_ItemPriorities(t *testing.T) {
	ts.AssertEqual(t, get_item_priority('a'), 1)
	ts.AssertEqual(t, get_item_priority('z'), 26)
	ts.AssertEqual(t, get_item_priority('A'), 27)
	ts.AssertEqual(t, get_item_priority('Z'), 52)
	ts.AssertEqual(t, get_item_priority('v'), 22)
	ts.AssertEqual(t, get_item_priority('L'), 38)
}

func TestD03_FindCommonItems(t *testing.T) {
	make_set := func(runes []rune) map[rune]bool {
		return f.AssociateWith(runes, func(rune) bool { return true })
	}
	type TestCase struct {
		input           []string
		expected_common string
	}

	ts.TestThat([]TestCase{
		{[]string{"aaAxv", "pvpqrx"}, "vx"},
		{[]string{"aaAXw", "pvpqrx"}, ""},
		{[]string{"aBcDeF", "aBcDeF", "ffffFfff"}, "F"},
		{[]string{"bbb", "BBB", "ffffFfff"}, ""},
		{[]string{"vJrwpWtwJgWr", "hcsFMMfFFhFp"}, "p"},
		{[]string{"weather", "feather", "neither"}, "ether"},
		{[]string{"plus", "minus", "thus", "bus", "just"}, "us"},
	}, func(tc TestCase) {
		ts.AssertEqualWithEqFunc(
			t,
			make_set(find_common_items(tc.input[0], tc.input[1:]...)),
			make_set([]rune(tc.expected_common)),
			f.MapEqual[rune, bool],
		)
	})
}

func TestD03_SolverTest(t *testing.T) {
	type Data = envelope.Envelope[[]m.Rucksack]
	// helper function
	data := func(rucksacks ...string) Data { return m.CreateRucksacksEnvelope(rucksacks) }

	ts.SolverTesterForComparableResults[Data, int](t).
		ProvideSolver(SumItemPriorities(CompartmentDuplicateItemLocator())).
		ProvideSolver(SumItemPriorities(BadgeItemLocator())).
		AddTestCase(
			data( // badges: B, W
				"Bapa", // a
				"xBMx", // x
				"WBMW", // W
				"txtW", // t
				"uWuL", // u
				"MoWo", // o
			),
			ts.ExpectResult(f.Sum(f.Map(get_item_priority, []rune("axWtuo")))),
			ts.ExpectResult(f.Sum(f.Map(get_item_priority, []rune("BW")))),
		).
		AddTestCase(
			data( // group #2 has no badge
				"Bapa",
				"xBMx",
				"WBMW",
				"txaZ", // no duplicate
				"uWuL",
				"MoWo",
			),
			ts.ExpectError[int]("rucksack #4", "no repeat"),
			ts.ExpectError[int]("group #2", "no badge"),
		).
		AddTestCase(
			data( // group #1 multiple badge candidates: B, x
				"Baxa",
				"xzBMxz", // multiple common items: x, z
				"oWBMWx",
				"txWpxZ",
				"uWuZ",
				"loWo",
			),
			ts.ExpectError[int]("rucksack #2", "multiple repeat"),
			ts.ExpectError[int]("group #1", "multiple badge"),
		).
		RunSolverTests()
}
