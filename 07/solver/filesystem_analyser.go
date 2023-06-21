package solver

import m "aoc/d07/models"

type FilesystemAnalyser interface {
	AnalyseAndGetAnswer(FilesystemSpec, *m.Directory) (int64, error)
}
