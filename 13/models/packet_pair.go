package models

import "aoc/functional"

type PacketPair = functional.Pair[Packet, Packet]

func PacketPairEqualityFunction(lhs, rhs PacketPair) bool {
	return PacketEqualityFunction(lhs.First, rhs.First) && PacketEqualityFunction(lhs.Second, rhs.Second)
}
