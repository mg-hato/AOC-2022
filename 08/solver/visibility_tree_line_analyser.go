package solver

import (
	c "aoc/common"
)

type visibility_tree_line_analyser struct {
	visibility_mapping map[position]bool
}

func VisibilityTreeLineAnalyser() TreeLineAnalyser {
	return &visibility_tree_line_analyser{}
}

func (vtla *visibility_tree_line_analyser) analyseTreeLine(tree_line []tree) {

	// tree on the edge is by default visible
	vtla.visibility_mapping[tree_line[0].pos] = true
	max_height := tree_line[0].height

	// analyse for others
	for i := 1; i < len(tree_line); i++ {
		if tree_line[i].height > max_height {
			max_height = tree_line[i].height
			vtla.visibility_mapping[tree_line[i].pos] = true
		}
	}
}

func (vtla visibility_tree_line_analyser) finishAndGetResult() int {
	return c.Count(c.GetValues(vtla.visibility_mapping), c.Identity[bool])
}

func (vtla *visibility_tree_line_analyser) initialise(forest [][]tree) {
	vtla.visibility_mapping = make(map[position]bool)

	// Assume by default that the tree is not visible
	// If during the course of analysis we observe it is visible, we mark it appropriately
	c.ForEach(func(t tree) {
		vtla.visibility_mapping[t.pos] = false
	}, c.Flatten(forest))
}
