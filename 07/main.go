package main

import (
	"aoc/argshandle"
	"aoc/d07/reader"
	s "aoc/d07/solver"
	"aoc/reading"
)

func main() {
	spec := s.SimpleFilesystemSpec(70_000_000, 30_000_000)
	argshandle.Program(
		reading.ReadWith(reader.TerminalOutputReader),
		s.AnalyseFilesystem(spec, s.SumDirectoriesOfSizeAtMost(100_000)),
		s.AnalyseFilesystem(spec, s.FindSmallestDirectoryEnablingUpdate()),
	)
}
