package solver

import (
	m "aoc/d07/models"
)

func AnalyseFilesystem(spec FilesystemSpec, analyser FilesystemAnalyser) func(m.SolverInput) (int64, error) {
	return func(input m.SolverInput) (int64, error) {
		root, err := m.CreateFilesystem(input.Get())
		if err != nil {
			return 0, err
		}
		return analyser.AnalyseAndGetAnswer(spec, root)
	}
}
