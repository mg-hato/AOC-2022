package models

import "fmt"

func PacketNumber(value int) Packet {
	return &packet_number{value}
}

type packet_number struct {
	Value int
}

func (packet packet_number) String() string {
	return fmt.Sprint(packet.Value)
}

func (packet packet_number) copy() Packet {
	return &packet_number{packet.Value}
}

func (packet packet_number) equals(rhs Packet) bool {
	rhs_packet_number, ok := rhs.(*packet_number)
	return ok && packet.Value == rhs_packet_number.Value
}

func (lhs packet_number) compare(rhs Packet) comparison_outcome {
	switch rhs := rhs.(type) {
	case *packet_number:
		switch {
		case lhs.Value < rhs.Value:
			return in_order
		case lhs.Value > rhs.Value:
			return not_in_order
		default:
			return same
		}
	case *packet_list:
		return PacketList(&lhs).compare(rhs)
	}
	return -1
}
