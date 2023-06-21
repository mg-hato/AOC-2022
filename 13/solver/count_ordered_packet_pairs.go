package solver

import (
	c "aoc/common"
	m "aoc/d13/models"
)

func CountOrderedPacketPairs(input m.SolverInput) (int, error) {
	return c.Sum(c.Map(c.GetFirst[int, m.PacketPair], c.Filter(
		func(enumerated_packet_pair c.Pair[int, m.PacketPair]) bool {
			return m.ArePacketsInOrder(
				enumerated_packet_pair.Second.First,
				enumerated_packet_pair.Second.Second,
			)
		},
		c.EnumerateWithFirstIndex[m.PacketPair](1)(input.Get()),
	))), nil
}
