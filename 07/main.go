package main

import (
	"aoc/argshandle"
	"aoc/day07/reader"
	s "aoc/day07/solver"
	"aoc/reading"
)

func main() {
	spec := s.SimpleFilesystemSpec(70_000_000, 30_000_000)
	argshandle.AoC2022DefaultProgram(
		reading.ReadWith(reader.TerminalOutputReader),
		s.AnalyseFilesystem(spec, s.SumDirectoriesOfSizeAtMost(100_000)),
		s.AnalyseFilesystem(spec, s.FindSmallestDirectoryEnablingUpdate()),
	)
}
