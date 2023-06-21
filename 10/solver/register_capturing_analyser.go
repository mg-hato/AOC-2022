package solver

import "aoc/d10/models"

type RegisterCapturingAnalyser interface {
	Initialise()
	IsDone() bool
	Capture(int)
	NextCycle() int

	GenerateReport() models.AnalyserReport
}
