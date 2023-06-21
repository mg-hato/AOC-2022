package models

import (
	c "aoc/common"
)

func PacketPairsEnvelope(packet_pairs ...PacketPair) c.Envelope[[]PacketPair] {
	return packet_pairs_envelope{packet_pairs}
}

type packet_pairs_envelope struct {
	packet_pairs []PacketPair
}

func (envelope packet_pairs_envelope) Get() []PacketPair {
	return c.Map(func(packet_pair PacketPair) PacketPair {
		return PacketPair{
			First:  packet_pair.First.copy(),
			Second: packet_pair.Second.copy(),
		}
	}, envelope.packet_pairs)
}

func PacketPairsEnvelopeEqualityFunction(lhs, rhs c.Envelope[[]PacketPair]) bool {
	return c.ArrayEqualWith(PacketPairEqualityFunction)(lhs.Get(), rhs.Get())
}
