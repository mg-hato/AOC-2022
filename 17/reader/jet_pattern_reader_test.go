package reader

import (
	"aoc/reading"
	ts "aoc/testers"
	"testing"
)

func TestD17_ReaderTest(t *testing.T) {
	ts.ReaderTesterForComparableData[string](t, reading.ReadWith(JetPatternReader)).
		AddTestCase("./tests/input-1.txt", ts.ExpectResult("<<<>><<<<><><><><<><><><")).
		AddTestCase("./tests/bad_line_input.txt", ts.ExpectError[string](
			"line #3",
			"<><><<><<<><|>>>>",
		)).
		AddTestCase(
			"./tests/empty_file.txt",
			ts.ExpectError[string](jet_pattern_not_read_validation_error().Error()),
		).
		RunReaderTests()
}
