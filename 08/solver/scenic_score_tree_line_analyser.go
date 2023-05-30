package solver

import f "aoc/functional"

type scenic_score_tree_line_analyser struct {
	scenic_scores map[position]int
}

func ScenicScoreTreeLineAnalyser() TreeLineAnalyser {
	return &scenic_score_tree_line_analyser{}
}

func (sstla *scenic_score_tree_line_analyser) analyseTreeLine(tree_line []tree) {
	scenic_scores := f.Repeat(0, 10)

	for i := 0; i < len(tree_line); i++ {
		sstla.scenic_scores[tree_line[i].pos] *= scenic_scores[tree_line[i].height]
		// trees of height less than current tree's height are blocked by it
		for height := byte(0); height <= tree_line[i].height; height++ {
			scenic_scores[height] = 1
		}
		// trees higher than the current tree's height see over it
		for height := tree_line[i].height + 1; height < 10; height++ {
			scenic_scores[height] += 1
		}
	}
}

func (sstla scenic_score_tree_line_analyser) finishAndGetResult() int {
	return f.Maximum(f.GetValues(sstla.scenic_scores), func(lhs, rhs int) bool {
		return lhs < rhs
	})
}

func (sstla *scenic_score_tree_line_analyser) initialise(forest [][]tree) {
	sstla.scenic_scores = make(map[position]int)
	f.ForEach(func(t tree) {
		sstla.scenic_scores[t.pos] = 1
	}, f.Flatten(forest))
}
