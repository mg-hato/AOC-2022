package models

import (
	e "aoc/envelope"
	f "aoc/functional"
)

func PacketPairsEnvelope(packet_pairs ...PacketPair) e.Envelope[[]PacketPair] {
	return packet_pairs_envelope{packet_pairs}
}

type packet_pairs_envelope struct {
	packet_pairs []PacketPair
}

func (envelope packet_pairs_envelope) Get() []PacketPair {
	return f.Map(func(packet_pair PacketPair) PacketPair {
		return PacketPair{
			First:  packet_pair.First.copy(),
			Second: packet_pair.Second.copy(),
		}
	}, envelope.packet_pairs)
}

func PacketPairsEnvelopeEqualityFunction(lhs, rhs e.Envelope[[]PacketPair]) bool {
	return f.ArrayEqualWith(PacketPairEqualityFunction)(lhs.Get(), rhs.Get())
}
