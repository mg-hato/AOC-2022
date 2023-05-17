package main

import (
	"aoc/day04/models"
	"aoc/day04/reader"
	s "aoc/day04/solver"
	"aoc/envelope"
	"aoc/reading"
	ts "aoc/testers"
	"testing"
)

func TestD04_IntegrationTest(t *testing.T) {
	type Data = envelope.Envelope[[]models.SectionAssignmentPair]
	ts.IntegrationTesterForComparableResults[Data, int](t).
		ProvideReader(reading.ReadWith(reader.SectionAssignmentsListReader)).
		ProvideSolver(s.CountAssignmentPairs(s.IsFullOverlap)).
		ProvideSolver(s.CountAssignmentPairs(s.IsPartialOverlap)).
		AddTestCase("./tests/example.txt", ts.ExpectResult(2), ts.ExpectResult(4)).
		// input-1.txt contains all possible (valid) pair-ups where section numbers are taken from range [1,20]; see at the bottom of the file for case counting
		AddTestCase("./tests/input-1.txt",
			ts.ExpectResult(20+760+2_280+190+4_560+9_690),
			ts.ExpectResult(20+760+2_280+190+4_560+9_690+2_280+9_690)).
		RunIntegrationTests()
}

//
// For full containment we have these cases:
// 1. both elves are in charge of exactly same section (e.g. 4-4,4-4): 20
// 2. one elf has single section and the other fully contains it:
//   2.a the other elf's left-or-right (but not both) matches that single section (e.g. 1-5,5-5): 2*2*(20C2) = 760
//   2.b the other elf's left-nor-right matches that single section (e.g. 1-10,7-7): 2*(20C3) = 2,280
// 3. both elves have both ends matching (left and right), but their Left != Right (e.g. 1-5,1-5): 20*19/2 = 190
// 4. both elves have matching one end (left-or-right), but differing other end (e.g. 1-5,3-5 or 7-10,7-15): 2*2*(20C3) = 4,560
// 5. both elves do not have any of their ends matching and one fully contains the other (e.g. 1-10,3-9): 2*(20C4) = 9,690
//
// For partial containment we have:
// 1. Everything in the full contaimnet: 20+760+2_280+190+4_560+9_690
// 2. Both elves have multiple sections assigned to them, but they share exactly one bordering section together (e.g. 1-7,7-10): 2*(20C3) = 2,280
// 3. Both elves have multiple sections assigned to them, and they share multiple sections, but they do not fully overlap (e.g. 1-5,3-7): 2*(20C4) =9690
