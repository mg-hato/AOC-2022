package reader

import (
	"aoc/d25/models"
	"aoc/reading"
	"aoc/testers"
	"testing"
)

func TestD25_Reader(t *testing.T) {
	sample_input_expected := models.CreateSnafuNumbersEnvelope([]string{
		"1",
		"2",
		"1=",
		"1-",
		"10",
		"11",
		"12",
		"2=",
		"2-",
		"20",
		"1=0",
		"1-0",
		"1=11-2",
		"1-0---0",
		"1121-1110-1=0",
	})
	testers.ReaderTester(t, reading.ReadWith(SnafuNumbersReader)).
		ProvideEqualityFunction(models.SnafuNumbersEnvelopeEqualityFunction).
		AddTestCase("./tests/sample_input.txt", testers.ExpectResult(sample_input_expected)).
		RunReaderTests()
}
