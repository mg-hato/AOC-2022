package main

import (
	"aoc/argshandle"
	"aoc/d22/models"
	"aoc/d22/reader"
	"aoc/d22/solver"
	"aoc/reading"
)

func main() {
	argshandle.Program(
		reading.ReadWith(reader.MonkeyMapReader),
		solver.GetFinalCoordinates(models.CreateSimpleFieldMap),
		solver.GetFinalCoordinates(models.CubeMapCreator),
	)
}
