package reader

import (
	m "aoc/d13/models"
	f "aoc/functional"
	"regexp"
	"strconv"
)

func parse_packet(line_number int, packet_string string) (m.Packet, error) {
	tokens := tokenise(packet_string)

	if verification_error := verify_tokens(line_number, tokens); verification_error != nil {
		return nil, verification_error
	}
	tokens = f.Filter(func(token string) bool { return token != "," }, tokens)
	subpackets := [][]m.Packet{{}}
	subpacket_depth := 0
	for _, token := range tokens[1 : len(tokens)-1] {
		switch get_token_type(token) {
		case left_par:
			subpackets = append(subpackets, []m.Packet{})
			subpacket_depth++
		case right_par:
			subpacket := m.PacketList(subpackets[subpacket_depth]...)
			subpackets = subpackets[:subpacket_depth]
			subpacket_depth--
			subpackets[subpacket_depth] = append(subpackets[subpacket_depth], subpacket)
		case number:
			value, _ := strconv.Atoi(token)
			subpackets[subpacket_depth] = append(subpackets[subpacket_depth], m.PacketNumber(value))
		}
	}
	return m.PacketList(subpackets[0]...), nil
}

func tokenise(packet_string string) []string {
	return f.Map(
		func(submatches []string) string { return submatches[0] },
		regexp.MustCompile(`\[|\]|,|\d+`).FindAllStringSubmatch(packet_string, -1),
	)
}
