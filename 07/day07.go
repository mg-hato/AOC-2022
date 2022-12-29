package main

import (
	"aoc/argshandle"
)

func main() {
	argshandle.AoC2022DefaultProgram(
		ReadTerminalOutput,
		ForFilesystemGet(SumOfDirectoriesLte100k),
		ForFilesystemGet(SmallestUpdateEnablingDirectorySize),
	)
}
