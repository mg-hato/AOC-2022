package solver

import (
	m "aoc/d13/models"
	e "aoc/envelope"
	f "aoc/functional"
	"sort"
)

func ExtractDecoderKey(envelope e.Envelope[[]m.PacketPair]) (int, error) {
	divider_packet_1 := m.PacketList(m.PacketList(m.PacketNumber(2)))
	divider_packet_2 := m.PacketList(m.PacketList(m.PacketNumber(6)))
	packets := append(
		f.FlatMap(
			func(packet_pair m.PacketPair) []m.Packet {
				return []m.Packet{
					packet_pair.First,
					packet_pair.Second,
				}
			},
			envelope.Get(),
		),
		divider_packet_1, divider_packet_2,
	)
	sort.Slice(packets, func(i, j int) bool {
		return m.ArePacketsInOrder(packets[i], packets[j])
	})
	get_index_of := func(packet m.Packet) int {
		return 1 + f.IndexOf(packets, func(p m.Packet) bool { return m.PacketEqualityFunction(p, packet) })
	}
	return get_index_of(divider_packet_1) * get_index_of(divider_packet_2), nil
}
