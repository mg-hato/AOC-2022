package solver

import (
	c "aoc/common"
	m "aoc/d08/models"
)

func AnalyseForest(analyser_supplier func() TreeLineAnalyser) func(m.SolverInput) (int, error) {
	return func(input m.SolverInput) (int, error) {

		forest := enumerate_forest(input.Get())

		analyser := analyser_supplier()
		analyser.initialise(forest)

		// analyse the forest from each possible viewpoint
		// this is achieved by transforming the original matrix
		for _, transformation := range []func([][]tree) [][]tree{
			func(x [][]tree) [][]tree { return x },
			func(x [][]tree) [][]tree { return c.Map(c.Reverse[tree], x) },
			func(x [][]tree) [][]tree { return transpose(x) },
			func(x [][]tree) [][]tree { return c.Map(c.Reverse[tree], transpose(x)) },
		} {
			c.ForEach(analyser.analyseTreeLine, transformation(forest))
		}

		return analyser.finishAndGetResult(), nil
	}
}
