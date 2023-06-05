package solver

import (
	m "aoc/d13/models"
	e "aoc/envelope"
	f "aoc/functional"
)

func CountOrderedPacketPairs(envelope e.Envelope[[]m.PacketPair]) (int, error) {
	return f.Sum(f.Map(f.GetFirst[int, m.PacketPair], f.Filter(
		func(enumerated_packet_pair f.Pair[int, m.PacketPair]) bool {
			return m.ArePacketsInOrder(
				enumerated_packet_pair.Second.First,
				enumerated_packet_pair.Second.Second,
			)
		},
		f.EnumerateWithFirstIndex(envelope.Get(), 1),
	))), nil
}
