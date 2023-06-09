package reader

import (
	m "aoc/d13/models"
	"aoc/envelope"
	"aoc/reading"
	ts "aoc/testers"
	"testing"
)

func TestD13_ReaderTest(t *testing.T) {
	type Data = envelope.Envelope[[]m.PacketPair]
	ts.ReaderTester(t, reading.ReadWith(PacketReader)).
		ProvideEqualityFunction(m.PacketPairsEnvelopeEqualityFunction).
		AddTestCase("./tests/input-1.txt", ts.ExpectResult(m.PacketPairsEnvelope(
			m.PacketPair{
				First: m.PacketList(
					m.PacketNumber(1),
					m.PacketNumber(2),
					m.PacketNumber(3),
					m.PacketList(
						m.PacketNumber(4),
						m.PacketNumber(5),
						m.PacketNumber(6),
						m.PacketList(
							m.PacketNumber(7),
							m.PacketNumber(8),
							m.PacketList(
								m.PacketNumber(9),
								m.PacketList(m.PacketList(m.PacketList(m.PacketList(m.PacketNumber(10))))),
							),
						),
					),
				),
				Second: m.PacketList(),
			},
			m.PacketPair{
				First:  m.PacketList(m.PacketList(m.PacketList())),
				Second: m.PacketList(m.PacketNumber(0)),
			},
			m.PacketPair{
				First:  m.PacketList(m.PacketList(m.PacketList(m.PacketNumber(1)))),
				Second: m.PacketList(m.PacketNumber(1)),
			},
			m.PacketPair{
				First: m.PacketList(m.PacketList(
					m.PacketNumber(1),
					m.PacketList(m.PacketNumber(2)),
					m.PacketNumber(3),
				)),
				Second: m.PacketList(m.PacketNumber(1), m.PacketNumber(2), m.PacketNumber(3)),
			},
		))).
		AddTestCase("./tests/bad-1.txt", ts.ExpectError[Data]("reader final validation error", "second packet", "missing")).
		AddTestCase("./tests/bad-2.txt", ts.ExpectError[Data]("line #8", "empty line", "expected", `"[13]"`)).
		AddTestCase("./tests/bad-3.txt", ts.ExpectError[Data]("line #6", "second packet", "expected", `""`)).
		AddTestCase("./tests/bad-4.txt", ts.ExpectError[Data]("line #2", "first token", `"777"`)).
		AddTestCase("./tests/bad-5.txt", ts.ExpectError[Data]("line #2", "sequence", "tokens", `"[7,5,"`)).
		AddTestCase("./tests/bad-6.txt", ts.ExpectError[Data]("line #3", "there is 1 outstanding", "unclosed", "bracket")).
		AddTestCase("./tests/bad-7.txt", ts.ExpectError[Data]("line #6", "last token")).
		AddTestCase("./tests/bad-8.txt", ts.ExpectError[Data]("line #1", "exactly one", "packet", " 2 ", "detected")).
		AddTestCase("./tests/bad-9.txt", ts.ExpectError[Data]("line #11", "bad number token", "packet", `"0015"`)).
		AddTestCase("./tests/bad-10.txt", ts.ExpectError[Data]("line #4", "too many closed brackets")).
		AddTestCase(
			"./tests/bad-11.txt",
			ts.ExpectError[Data]("reader final validation error", "packet number #4", "divider packet [[2]]"),
		).
		AddTestCase(
			"./tests/bad-12.txt",
			ts.ExpectError[Data]("reader final validation error", "packet number #5", "divider packet [[6]]"),
		).
		RunReaderTests()
}
