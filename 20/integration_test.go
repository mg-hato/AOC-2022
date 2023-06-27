package main

import (
	m "aoc/d20/models"
	"aoc/d20/reader"
	"aoc/d20/solver"
	"aoc/reading"
	ts "aoc/testers"
	"testing"
)

func TestD20_IntegrationTest(t *testing.T) {
	ts.IntegrationTesterForComparableResults[m.SolverInput, int64](t).
		ProvideReader(reading.ReadWith(reader.EncryptedFileReader)).
		ProvideSolver(solver.Decrypt(1, 1)).
		ProvideSolver(solver.Decrypt(10, 811_589_153)).
		AddTestCase("./tests/example.txt", ts.ExpectResult[int64](3), ts.ExpectResult[int64](1_623_178_306)).
		RunIntegrationTests()
}
