package models

type Packet interface {
	String() string

	copy() Packet
	equals(Packet) bool
	compare(Packet) comparison_outcome
}

func PacketEqualityFunction(lhs, rhs Packet) bool {
	return lhs.equals(rhs)
}

func ArePacketsInOrder(lhs, rhs Packet) bool {
	return lhs.compare(rhs) == in_order
}
