package solver

type TreeLineAnalyser interface {
	initialise([][]tree)
	analyseTreeLine([]tree)
	finishAndGetResult() int
}
