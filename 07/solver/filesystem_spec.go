package solver

type FilesystemSpec interface {
	GetTotalMemory() int64
	GetMemoryRequiredForUpdate() int64
}
