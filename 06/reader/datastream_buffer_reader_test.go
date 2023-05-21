package reader

import (
	"aoc/reading"
	ts "aoc/testers"
	"testing"
)

func TestD06_ReaderTest(t *testing.T) {
	ts.ReaderTesterForComparableData(t, reading.ReadWith(DatastreamBufferReader)).
		AddTestCase("./tests/input-1.txt", ts.ExpectResult("sagydjbhnlaskdmjiashdiashdasjkncxjkbvujlsfdhguosihdnh")).
		AddTestCase("./tests/bad-input-1.txt", ts.ExpectError[string]("not", "interpret", "line #1", "sagydjbhnlaskdmjiasdsadasXashdiashdasjkncxjkbvujlsfdhguosihdnh")).
		RunReaderTests()

}
