package models

import (
	f "aoc/functional"
	"fmt"
	"strings"
)

func PacketList(packets ...Packet) Packet {
	return &packet_list{packets}
}

type packet_list struct {
	Packets []Packet
}

func (packet packet_list) String() string {
	return fmt.Sprintf("[%s]", strings.Join(f.Map(Packet.String, packet.Packets), ","))
}

func (packet packet_list) copy() Packet {
	return &packet_list{f.Map(Packet.copy, packet.Packets)}
}

func (packet packet_list) equals(rhs Packet) bool {
	rhs_packet_list, ok := rhs.(*packet_list)
	return ok && f.ArrayEqualWith(PacketEqualityFunction)(packet.Packets, rhs_packet_list.Packets)
}

func (lhs packet_list) compare(rhs Packet) comparison_outcome {
	switch rhs := rhs.(type) {
	case *packet_number:
		return lhs.compare(PacketList(rhs))
	case *packet_list:
		i := 0
		for i < len(lhs.Packets) && i < len(rhs.Packets) {
			if outcome := lhs.Packets[i].compare(rhs.Packets[i]); outcome != same {
				return outcome
			}
			i++
		}
		switch {
		case len(lhs.Packets) < len(rhs.Packets):
			return in_order
		case len(lhs.Packets) > len(rhs.Packets):
			return not_in_order
		default:
			return same
		}
	}
	return -1
}
