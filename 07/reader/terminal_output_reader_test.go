package reader

import (
	m "aoc/d07/models"
	"aoc/reading"
	ts "aoc/testers"
	"testing"
)

func TestD07_ReaderTest(t *testing.T) {
	ts.ReaderTester(t, reading.ReadWith(TerminalOutputReader)).
		ProvideEqualityFunction(m.CommandsEnvelopeEqualityFunction).
		AddTestCase("./tests/input-1.txt", ts.ExpectResult(m.CreateCommandsEnvelope(
			m.MakeCommandCd("/"),
			m.MakeCommandLs(
				m.MakePartialDirectory("pictures"),
				m.MakeFile("p.pdf", 10_134),
				m.MakePartialDirectory("music"),
			),
			m.MakeCommandCd("music"),
			m.MakeCommandLs(
				m.MakeFile("mamasc.wav", 245_689),
				m.MakeFile("strangersinthenight.wav", 301_011),
			),
			m.MakeCommandCd(".."),
			m.MakeCommandCd("pictures"),
			m.MakeCommandLs(
				m.MakeFile("dog.jpg", 15_987),
				m.MakeFile("cat.jpg", 15_101),
			),
		))).
		AddTestCase("./tests/bad-input-1.txt", ts.ExpectError[m.SolverInput]("line #1", "first", "command", "cd /", "cd xyz")).
		AddTestCase("./tests/bad-input-2.txt", ts.ExpectError[m.SolverInput]("line #3", ".gitignore")).
		AddTestCase("./tests/bad-input-3.txt", ts.ExpectError[m.SolverInput]("line #6", "duplicate")).
		RunReaderTests()
}
