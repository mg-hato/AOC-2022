package solver

import (
	c "aoc/common"
	m "aoc/d08/models"
	ts "aoc/testers"
	"testing"
)

func TestD08_Transpose(t *testing.T) {
	ts.AssertEqualWithEqFunc(
		t,
		transpose([][]int{
			{12, 24, 36, 48},
			{0, 0, 0, 0},
			{7, 77, 777, 7_777},
			{2, 4, 16, 256},
		}),
		[][]int{
			{12, 0, 7, 2},
			{24, 0, 77, 4},
			{36, 0, 777, 16},
			{48, 0, 7_777, 256},
		},
		c.ArrayEqualWith(c.ArrayEqual[int]),
	)
}

func TestD08_ForestEnumeration(t *testing.T) {
	ts.AssertEqualWithEqFunc(
		t,
		enumerate_forest(m.Forest{
			{3, 6, 9},
			{5, 0, 5},
		}),
		[][]tree{
			{{3, make_position(0, 0)}, {6, make_position(0, 1)}, {9, make_position(0, 2)}},
			{{5, make_position(1, 0)}, {0, make_position(1, 1)}, {5, make_position(1, 2)}},
		},
		c.ArrayEqualWith(c.ArrayEqual[tree]),
	)
}

func TestD08_VisibilityTreeLineAnalyser(t *testing.T) {
	forest := enumerate_forest(m.Forest{
		{7, 6, 7, 8}, // trees in the middle are not visible (6 and 7); others are
		{1, 5, 4, 3}, // first two are visible (1 and 5); trees behind are not visible (4 and 3)
		{7, 4, 2, 2}, // only first tree (7) is visible
		{9, 9, 2, 4}, // only first tree (9) is visible
	})
	analyser := visibility_tree_line_analyser{}
	analyser.initialise(forest)
	c.ForEach(analyser.analyseTreeLine, forest)

	visible := true
	not_visible := false

	expected_visibility_mapping := map[position]bool{
		make_position(0, 0): visible,
		make_position(0, 1): not_visible,
		make_position(0, 2): not_visible,
		make_position(0, 3): visible,

		make_position(1, 0): visible,
		make_position(1, 1): visible,
		make_position(1, 2): not_visible,
		make_position(1, 3): not_visible,

		make_position(2, 0): visible,
		make_position(2, 1): not_visible,
		make_position(2, 2): not_visible,
		make_position(2, 3): not_visible,

		make_position(3, 0): visible,
		make_position(3, 1): not_visible,
		make_position(3, 2): not_visible,
		make_position(3, 3): not_visible,
	}

	ts.AssertEqualWithEqFunc(
		t,
		analyser.visibility_mapping,
		expected_visibility_mapping,
		c.MapEqual[position, bool],
	)
}

func TestD08_ScenicScoreTreeLineAnalyser(t *testing.T) {
	forest := enumerate_forest(m.Forest{
		{7, 6, 7, 8}, // trees in the middle are not visible (6 and 7); others are
		{1, 5, 4, 3}, // first two are visible (1 and 5); trees behind are not visible (4 and 3)
		{7, 4, 2, 2}, // only first tree (7) is visible
		{9, 9, 2, 4}, // only first tree (9) is visible
	})
	analyser := scenic_score_tree_line_analyser{}
	analyser.initialise(forest)
	c.ForEach(analyser.analyseTreeLine, forest)

	expected_scenic_scores := map[position]int{
		// 7678
		make_position(0, 0): 0, // edge
		make_position(0, 1): 1, // blocked by left neighbour
		make_position(0, 2): 2, // sees until first 7
		make_position(0, 3): 3, // sees everything until the left edge

		// 1543
		make_position(1, 0): 0, // edge
		make_position(1, 1): 1, // sees everything until the left edge
		make_position(1, 2): 1, // blocked by left neighbour
		make_position(1, 3): 1, // blocked by left neighbour

		// 7422
		make_position(2, 0): 0, // edge
		make_position(2, 1): 1, // blocked by left neighbour
		make_position(2, 2): 1, // blocked by left neighbour
		make_position(2, 3): 1, // blocked by left neighbour

		// 9924
		make_position(3, 0): 0, // edge
		make_position(3, 1): 1, // blocked by left neighbour
		make_position(3, 2): 1, // blocked by left neighbour
		make_position(3, 3): 2, // sees until the second 9
	}

	ts.AssertEqualWithEqFunc(
		t,
		analyser.scenic_scores,
		expected_scenic_scores,
		c.MapEqual[position, int],
	)

}
