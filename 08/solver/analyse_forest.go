package solver

import (
	m "aoc/day08/models"
	e "aoc/envelope"
	f "aoc/functional"
)

func AnalyseForest(analyser_supplier func() TreeLineAnalyser) func(e.Envelope[m.Forest]) (int, error) {
	return func(envelope e.Envelope[m.Forest]) (int, error) {

		forest := enumerate_forest(envelope.Get())

		analyser := analyser_supplier()
		analyser.initialise(forest)

		// analyse the forest from each possible viewpoint
		// this is achieved by transforming the original matrix
		for _, transformation := range []func([][]tree) [][]tree{
			func(x [][]tree) [][]tree { return x },
			func(x [][]tree) [][]tree { return f.Map(f.Reverse[tree], x) },
			func(x [][]tree) [][]tree { return transpose(x) },
			func(x [][]tree) [][]tree { return f.Map(f.Reverse[tree], transpose(x)) },
		} {
			f.ForEach(analyser.analyseTreeLine, transformation(forest))
		}

		return analyser.finishAndGetResult(), nil
	}
}
