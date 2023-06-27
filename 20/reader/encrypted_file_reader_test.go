package reader

import (
	m "aoc/d20/models"
	"aoc/reading"
	"aoc/testers"
	"testing"
)

func TestD20_ReaderTest(t *testing.T) {
	testers.ReaderTester(t, reading.ReadWith(EncryptedFileReader)).
		ProvideEqualityFunction(m.EncryptedFileEnvelopeEqualityFunction).
		AddTestCase("./tests/input-1.txt", testers.ExpectResult(m.EncryptedFileEnvelope([]int{
			1, 2, 3, 4, 5, 0,
		}))).
		RunReaderTests()
}
