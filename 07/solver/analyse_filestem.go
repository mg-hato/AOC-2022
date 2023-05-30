package solver

import (
	m "aoc/day07/models"
	e "aoc/envelope"
)

func AnalyseFilesystem(spec FilesystemSpec, analyser FilesystemAnalyser) func(e.Envelope[[]m.Command]) (int64, error) {
	return func(envelope e.Envelope[[]m.Command]) (int64, error) {
		root, err := m.CreateFilesystem(envelope.Get())
		if err != nil {
			return 0, err
		}
		return analyser.AnalyseAndGetAnswer(spec, root)
	}
}
