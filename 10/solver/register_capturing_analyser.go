package solver

import "aoc/day10/models"

type RegisterCapturingAnalyser interface {
	Initialise()
	IsDone() bool
	Capture(int)
	NextCycle() int

	GenerateReport() models.AnalyserReport
}
