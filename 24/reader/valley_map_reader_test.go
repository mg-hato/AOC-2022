package reader

import (
	"aoc/d24/models"
	"aoc/reading"
	"aoc/testers"
	"testing"
)

func TestD24_Reader(t *testing.T) {
	expected_result := models.CreateValleyMapEnvelope([][]rune{
		[]rune("#.#####"),
		[]rune("#..v..#"),
		[]rune("#..><.#"),
		[]rune("#####.#"),
	})

	testers.ReaderTester(t, reading.ReadWith(ValleyMapReader)).
		ProvideEqualityFunction(models.ValleyMapEnvelopeEqualityFunction).
		AddTestCase("./tests/sample_input.txt", testers.ExpectResult(expected_result)).
		RunReaderTests()
}
