package solver

import (
	c "aoc/common"
	m "aoc/d07/models"
	"fmt"
	"sort"
)

func FindSmallestDirectoryEnablingUpdate() FilesystemAnalyser {
	return find_smallest_directory_enabling_update_analyser{}
}

type find_smallest_directory_enabling_update_analyser struct{}

func (analyser find_smallest_directory_enabling_update_analyser) AnalyseAndGetAnswer(spec FilesystemSpec, root *m.Directory) (int64, error) {
	size_mapping := m.CalculateItemSizes(root)
	directories_sizes := make([]int64, 0)
	for item, size := range size_mapping {
		if _, ok := item.(*m.Directory); ok {
			directories_sizes = append(directories_sizes, size)
		}
	}
	free_memory := spec.GetTotalMemory() - size_mapping[root]
	if free_memory < 0 {
		return 0, fmt.Errorf("error: the free memory is a negative number: %d", free_memory)
	}

	if free_memory >= spec.GetMemoryRequiredForUpdate() {
		return 0, fmt.Errorf("error: there is enough free memory for update")
	}

	missing_memory := spec.GetMemoryRequiredForUpdate() - free_memory
	directories_sizes = c.Filter(func(size int64) bool { return size >= missing_memory }, directories_sizes)
	if len(directories_sizes) == 0 {
		return 0, fmt.Errorf("error: could not find candidate directories for deletion")
	}

	sort.Slice(directories_sizes, func(i, j int) bool {
		return directories_sizes[i] < directories_sizes[j]
	})
	return directories_sizes[0], nil
}
