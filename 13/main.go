package main

import (
	"aoc/argshandle"
	"aoc/d13/reader"
	"aoc/d13/solver"
	"aoc/reading"
)

func main() {
	argshandle.AoC2022DefaultProgram(
		reading.ReadWith(reader.PacketReader),
		solver.CountOrderedPacketPairs,
		solver.ExtractDecoderKey,
	)
}
