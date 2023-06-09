package models

import (
	f "aoc/functional"
	ts "aoc/testers"
	"testing"
)

func TestD13_EnvelopeTest(t *testing.T) {
	data := func() []PacketPair {
		return []PacketPair{
			{
				First:  PacketList(PacketList(), PacketNumber(12), PacketList(PacketNumber(1), PacketNumber(2))),
				Second: PacketList(PacketNumber(4), PacketList(PacketNumber(40), PacketList(PacketNumber(400)))),
			},
			{
				First:  PacketList(PacketList(PacketNumber(64), PacketNumber(36))),
				Second: PacketList(PacketNumber(365)),
			},
		}
	}
	envelope := PacketPairsEnvelope(data()...)
	packets := envelope.Get()
	packets[0].First = PacketList(PacketNumber(0))

	ts.AssertEqualWithEqFunc(t, envelope.Get(), data(), f.ArrayEqualWith(PacketPairEqualityFunction))
}

func TestD13_PacketComparison(t *testing.T) {
	// shortcut function names
	PL, PN := PacketList, PacketNumber
	packets := []Packet{
		PL(),                                   // []
		PL(PL()),                               // [[]]
		PL(PL(PL())),                           // [[[]]]
		PL(PN(1), PN(1), PN(3), PN(1), PN(1)),  // [1,1,3,1,1]
		PL(PN(1), PN(1), PN(5), PN(1), PN(1)),  // [1,1,5,1,1]
		PL(PL(PN(1)), PL(PN(2), PN(3), PN(4))), // [[1],[2,3,4]]
		PL(PN(1), PL(PN(2), PL(PN(3), PL(PN(4), PL(PN(5), PN(6), PN(0))))), PN(8), PN(9)), // [1,[2,[3,[4,[5,6,0]]]],8,9]
		PL(PN(1), PL(PN(2), PL(PN(3), PL(PN(4), PL(PN(5), PN(6), PN(7))))), PN(8), PN(9)), // [1,[2,[3,[4,[5,6,7]]]],8,9]
		PL(PL(PN(1)), PN(4)),                      // [[1], 4]
		PL(PN(3)),                                 // [3]
		PL(PL(PN(4), PN(4)), PN(4), PN(4)),        // [[4,4],4,4]
		PL(PL(PN(4), PN(4)), PN(4), PN(4), PN(4)), // [[4,4],4,4,4]
		PL(PN(7), PN(7), PN(7)),                   // [7,7,7]
		PL(PN(7), PN(7), PN(7), PN(7)),            // [7,7,7,7]
		PL(PL(PN(8), PN(7), PN(6))),               // [[8,7,6]]
		PL(PN(9)),                                 // [9]
	}

	for left := 0; left < len(packets); left++ {
		for right := 0; right < len(packets); right++ {
			expected_result := comparison_outcome(0)
			switch {
			case left < right:
				expected_result = in_order
			case left > right:
				expected_result = not_in_order
			default:
				expected_result = same
			}
			type to_assert struct {
				result      comparison_outcome
				left, right int
			}
			pack_it := func(res comparison_outcome) to_assert {
				return to_assert{
					result: res,
					left:   left,
					right:  right,
				}
			}
			ts.AssertEqual(t, pack_it(packets[left].compare(packets[right])), pack_it(expected_result))
		}
	}
}
