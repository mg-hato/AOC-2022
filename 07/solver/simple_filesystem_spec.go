package solver

func SimpleFilesystemSpec(total_memory int64, memory_required_for_update int64) FilesystemSpec {
	return simple_filesystem_spec{
		total_memory:               total_memory,
		memory_required_for_update: memory_required_for_update,
	}
}

type simple_filesystem_spec struct {
	total_memory               int64
	memory_required_for_update int64
}

func (spec simple_filesystem_spec) GetTotalMemory() int64 {
	return spec.total_memory
}

func (spec simple_filesystem_spec) GetMemoryRequiredForUpdate() int64 {
	return spec.memory_required_for_update
}
