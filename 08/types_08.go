package main

import (
	. "aoc/functional"
	"sort"
)

type Tree struct {
	id     int // id
	height byte
}

type ForestAnalyser interface {
	AnalyseForestRow([]Tree)
	GetResult() int
}

// VisibilityAnalyser
// Each forest row gets analysed for the visibility of a tree

type VisibilityAnalyster struct {
	visibility map[int]bool
}

func (analyser *VisibilityAnalyster) AnalyseForestRow(trees []Tree) {
	var max_height_set bool
	var max_height byte
	for _, tree := range trees {
		if !max_height_set || tree.height > max_height {
			max_height_set = true
			max_height = tree.height
			analyser.visibility[tree.id] = true
		}
	}
}

func (analyser *VisibilityAnalyster) GetResult() int {
	return len(Filter(Identity[bool], GetValues(analyser.visibility)))
}

func NewVisibilityAnalyser() ForestAnalyser {
	return &VisibilityAnalyster{
		visibility: make(map[int]bool),
	}
}

// ScenicScoreAnalyser

type ScenicScoreAnalyser struct {
	scenic_scores map[int]int
}

// For each tree, it finds out the visible distance to the "left" and updates the scenic score of that tree's id
func (analyser *ScenicScoreAnalyser) AnalyseForestRow(trees []Tree) {

	// Maximum height: M
	M := byte(10)

	distance_for_height := make([]int, M)

	for _, tree := range trees {
		if _, present := analyser.scenic_scores[tree.id]; !present {
			analyser.scenic_scores[tree.id] = 1
		}
		analyser.scenic_scores[tree.id] *= distance_for_height[tree.height]

		for i := byte(0); i < M; i++ {
			if i <= tree.height {
				distance_for_height[i] = 1
			} else {
				distance_for_height[i]++
			}
		}
	}

}

func (analyser *ScenicScoreAnalyser) GetResult() int {
	scores := GetValues(analyser.scenic_scores)
	sort.Slice(scores, func(i, j int) bool {
		return scores[i] > scores[j]
	})
	return scores[0]
}

func NewScenicScoreAnalyser() ForestAnalyser {
	return &ScenicScoreAnalyser{
		scenic_scores: map[int]int{},
	}
}
