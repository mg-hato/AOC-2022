package main

import (
	"aoc/argshandle"
	"aoc/d20/reader"
	"aoc/d20/solver"
	"aoc/reading"
)

func main() {
	argshandle.Program(
		reading.ReadWith(reader.EncryptedFileReader),
		solver.Decrypt(1, int64(1)),
		solver.Decrypt(10, int64(811_589_153)),
	)
}
