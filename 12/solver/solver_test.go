package solver

import (
	c "aoc/common"
	m "aoc/d12/models"
	ts "aoc/testers"
	"testing"
)

func TestD12_FindAllDistances(t *testing.T) {
	terrain := m.EnumerateTerrain([]string{
		//            rows:
		"Sabcdef", // 0
		"iiyiiix", // 1
		"uxyzEyy", // 2
		"bwbbbbb", // 3
		"pzppppp", // 4
	})
	ts.AssertEqualWithEqFunc(
		t,
		findAllDistancesFromGoalPosition(terrain),
		map[m.Field]int{
			terrain[2][4]: 0, // E

			// To the right of E
			terrain[2][5]: 1, // y
			terrain[2][6]: 2, // y
			terrain[1][6]: 3, // x

			// To the left of E
			terrain[2][3]: 1, // z
			terrain[2][2]: 2, // y
			terrain[2][1]: 3, // x
			terrain[1][2]: 3, // y
			terrain[3][1]: 4, // w
			terrain[4][1]: 5, // z on row 4
		},
		c.MapEqual[m.Field, int],
	)
}

func TestD12_StartingPositionDistancePicker(t *testing.T) {

	var counter int = 0
	// make field helper
	make_field := func(r rune) m.Field {
		first := counter
		counter++
		return m.Field{
			HeightCode: r,
			Position:   m.Position{First: first},
		}
	}

	distance_mappings := map[m.Field]int{
		make_field('a'): 10,
		make_field('E'): 0,
		make_field('S'): 100,
		make_field('a'): 12,
		make_field('a'): 122,
		make_field('a'): 9,
		make_field('c'): 7,
		make_field('d'): 6,
	}

	type test_input struct {
		dp             DistancePicker
		expect_error   bool
		expected_value int
	}

	ts.TestThat([]test_input{
		{dp: StartingPositionDistancePicker('S'), expected_value: 100},
		{dp: StartingPositionDistancePicker('S', 'a'), expected_value: 9},
		{dp: StartingPositionDistancePicker('a'), expected_value: 9},
		{dp: StartingPositionDistancePicker('r'), expect_error: true},
		{dp: StartingPositionDistancePicker('z'), expect_error: true},
		{dp: StartingPositionDistancePicker('a', 'z', 'c'), expected_value: 7},
	}, func(ti test_input) {
		val, err := ti.dp.getDistance(distance_mappings)
		if ti.expect_error {
			ts.AssertError(t, err)
		} else {
			ts.AssertNoError(t, err)
			ts.AssertEqual(t, val, ti.expected_value)
		}
	})
}
