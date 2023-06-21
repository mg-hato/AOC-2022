package models

import c "aoc/common"

type PacketPair = c.Pair[Packet, Packet]

func PacketPairEqualityFunction(lhs, rhs PacketPair) bool {
	return PacketEqualityFunction(lhs.First, rhs.First) && PacketEqualityFunction(lhs.Second, rhs.Second)
}
