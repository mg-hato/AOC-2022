package solver

import (
	m "aoc/day07/models"
)

func SumDirectoriesOfSizeAtMost(size int64) FilesystemAnalyser {
	return sum_directories_sizes_analyser{
		func(_ *m.Directory, directory_size int64) bool {
			return directory_size <= size
		},
	}
}

type sum_directories_sizes_analyser struct {
	predicate func(*m.Directory, int64) bool
}

func (analyser sum_directories_sizes_analyser) AnalyseAndGetAnswer(_ FilesystemSpec, root *m.Directory) (int64, error) {
	size_mappings := m.CalculateItemSizes(root)
	size_sum := int64(0)
	for item, size := range size_mappings {
		if d, ok := item.(*m.Directory); ok && analyser.predicate(d, size) {
			size_sum += size
		}
	}
	return size_sum, nil
}
