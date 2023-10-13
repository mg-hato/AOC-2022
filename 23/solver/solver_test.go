package solver

import (
	"aoc/common"
	c "aoc/common"
	"aoc/d23/models"
	m "aoc/d23/models"
	"aoc/d23/reader"
	"aoc/reading"
	ts "aoc/testers"
	"testing"
)

func TestD22_step_by_step_example(t *testing.T) {
	envelope, err := reading.ReadWith(reader.GroveMapReader)("../tests/example.txt")
	ts.AssertNoError(t, err)

	// p := m.MakePosition

	// ts.AssertEqualWithEqFunc(t, testtt(0)(envelope), map[m.Position]bool{
	// 	p(2, 7): true,
	// 	p(3, 5): true, p(3, 6): true, p(3, 7): true, p(3, 9): true,
	// 	p(4, 3): true, p(4, 7): true, p(4, 9): true,
	// 	p(5, 4): true, p(5, 8): true, p(5, 9): true,
	// 	p(6, 3): true, p(6, 5): true, p(6, 6): true, p(6, 7): true,
	// 	p(7, 3): true, p(7, 4): true, p(7, 6): true, p(7, 8): true, p(7, 9): true,
	// 	p(8, 4): true, p(8, 7): true,
	// }, c.SetEqual[m.Position])

	// pos := m.MakePosition(4, 7)
	ts.AssertEqualWithEqFunc(t, testtt(0)(envelope), to_nice_map(to_spottypes([]string{
		"..............",
		"..............",
		".......#......",
		".....###.#....",
		"...#...#.#....",
		"....#...##....",
		"...#.###......",
		"...##.#.##....",
		"....#..#......",
		"..............",
		"..............",
		"..............",
	})), c.SetEqual[m.Position])

	expected_elves_positions := [][]string{
		{ // No iterations:
			"..............",
			"..............",
			".......#......",
			".....###.#....",
			"...#...#.#....",
			"....#...##....",
			"...#.###......",
			"...##.#.##....",
			"....#..#......",
			"..............",
			"..............",
			"..............",
		},
		{ // 1 iteration:
			"..............",
			".......#......",
			".....#...#....",
			"...#..#.#.....",
			".......#..#...",
			"....#.#.##....",
			"..#..#.#......",
			"..#.#.#.##....",
			"..............",
			"....#..#......",
			"..............",
			"..............",
		},
		{ // 2 iterations:
			"..............",
			".......#......",
			"....#.....#...",
			"...#..#.#.....",
			".......#...#..",
			"...#..#.#.....",
			".#...#.#.#....",
			"..............",
			"..#.#.#.##....",
			"....#..#......",
			"..............",
			"..............",
		},
		{ // 3 iterations
			"..............",
			".......#......",
			".....#....#...",
			"..#..#...#....",
			".......#...#..",
			"...#..#.#.....",
			".#..#.....#...",
			".......##.....",
			"..##.#....#...",
			"...#..........",
			".......#......",
			"..............",
		},
		{ // 4 iterations
			"..............",
			".......#......",
			"......#....#..",
			"..#...##......",
			"...#.....#.#..",
			".........#....",
			".#...###..#...",
			"..#......#....",
			"....##....#...",
			"....#.........",
			".......#......",
			"..............",
		},
		{ // 5 iterations
			".......#......",
			"..............",
			"..#..#.....#..",
			".........#....",
			"......##...#..",
			".#.#.####.....",
			"...........#..",
			"....##..#.....",
			"..#...........",
			"..........#...",
			"....#..#......",
			"..............",
		},
		{ // 10 iterations
			".......#......",
			"...........#..",
			"..#.#..#......",
			"......#.......",
			"...#.....#..#.",
			".#......##....",
			".....##.......",
			"..#........#..",
			"....#.#..#....",
			"..............",
			"....#..#..#...",
			"..............",
		},
	}

	for i, iterations := range []int{0, 1, 2, 3, 4, 5, 10} {
		ts.AssertEqualWithEqFunc(t, testtt(iterations)(envelope), to_nice_map(to_spottypes(expected_elves_positions[i])), c.SetEqual[m.Position])

	}
	// positions := testtt(0)(envelope)
	// fmt.Printf("%v %v %v\n", pos.Adjacent(m.South), pos.Adjacent(m.SouthEast), pos.Adjacent(m.SouthWest))
	// fmt.Printf("%v %v %v\n", positions[pos.Adjacent(m.South)], positions[pos.Adjacent(m.SouthEast)], positions[pos.Adjacent(m.SouthWest)])
	// fmt.Printf("%v", get_iteration_movement_schedule(positions, 0)(m.MakePosition(4, 7)))
	// // fmt.Printf("%v\n", )

	// ts.AssertEqualWithEqFunc(t, testtt(1)(envelope), to_nice_map(to_spottypes([]string{
	// 	"..............",
	// 	".......#......",
	// 	".....#...#....",
	// 	"...#..#.#.....",
	// 	".......#..#...",
	// 	"....#.#.##....",
	// 	"..#..#.#......",
	// 	"..#.#.#.##....",
	// 	"..............",
	// 	"....#..#......",
	// 	"..............",
	// 	"..............",
	// })), c.SetEqual[m.Position])

	// ts.AssertEqualWithEqFunc(t, testtt())

}

func to_spottypes(ss []string) [][]m.SpotType {
	return c.Map(
		func(s string) []m.SpotType {
			return c.Map(
				func(r rune) m.SpotType {
					ret, _ := m.TryParseSpotType(r)
					return ret
				}, []rune(s),
			)
		}, ss,
	)
}

func testtt(iterations int) func(common.Envelope[[][]models.SpotType]) map[m.Position]bool {
	return func(e c.Envelope[[][]models.SpotType]) map[m.Position]bool {
		elves := to_nice_map(e.Get())
		for i := 0; i < iterations; i++ {
			elves = iteration(elves, i)
		}
		return elves
	}
}
