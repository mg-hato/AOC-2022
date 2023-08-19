package reader

import (
	c "aoc/common"
	m "aoc/d21/models"
	"aoc/reading"
	ts "aoc/testers"
	"testing"
)

func TestD21_ReaderTest(t *testing.T) {
	num, id := m.CreateNumber, m.CreateIdentifier
	sjob := m.CreateSingleJob
	opjob := func(lhs m.Operand, op string, rhs m.Operand) m.MonkeyJob {
		op_id, _ := m.TryParseOperation(op)
		return m.CreateTwoOperandJob(lhs, op_id, rhs)
	}

	ts.ReaderTester(t, reading.ReadWith(MonkeyJobsReader)).
		ProvideEqualityFunction(m.MokeyJobsEnvelopeEqualityFunc).
		AddTestCase("./tests/valid_input_1.txt", ts.ExpectResult(m.CreateMonkeyJobsEnvelope(
			c.MakePair("root", sjob(num(189))),
			c.MakePair("humn", sjob(num(99))),
		))).
		AddTestCase("./tests/valid_input_2.txt", ts.ExpectResult(m.CreateMonkeyJobsEnvelope(
			c.MakePair("root", opjob(id("abc"), "-", id("humn"))),
			c.MakePair("abc", opjob(id("ghijk"), "*", id("def"))),
			c.MakePair("def", sjob(num(255))),
			c.MakePair("ghijk", opjob(id("humn"), "+", id("def"))),
			c.MakePair("humn", sjob(num(101))),
		))).
		RunReaderTests()
}
