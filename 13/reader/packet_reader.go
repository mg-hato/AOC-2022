package reader

import (
	m "aoc/d13/models"
	e "aoc/envelope"
	f "aoc/functional"
	"aoc/reading"
	"regexp"
)

type packet_reader struct {
	err error

	line_number int

	current_status status
	line_readers   map[status]func(string)

	packet_re *regexp.Regexp
	empty_re  *regexp.Regexp

	packet_pairs []m.PacketPair
}

func PacketReader() reading.ReaderAoC2022[e.Envelope[[]m.PacketPair]] {
	packet_reader := &packet_reader{
		current_status: first_packet,

		empty_re:  regexp.MustCompile("^ *$"),
		packet_re: regexp.MustCompile(`^(?:\[|\]| |,|\d)*$`),
	}

	packet_reader.line_readers = map[int]func(string){
		empty_line:    packet_reader.readEmptyLine,
		first_packet:  packet_reader.readFirstPacket,
		second_packet: packet_reader.readSecondPacket,
	}

	return packet_reader
}

func (pr packet_reader) Error() error {
	return pr.err
}

func (pr packet_reader) PerformFinalValidation() error {
	if pr.current_status == second_packet {
		return second_packet_missing_final_validation_error()
	}

	flattened_packets := f.FlatMap(
		func(packet_pair m.PacketPair) []m.Packet {
			return []m.Packet{packet_pair.First, packet_pair.Second}
		},
		pr.packet_pairs,
	)

	for _, divider_packet := range []m.Packet{
		m.PacketList(m.PacketList(m.PacketNumber(2))),
		m.PacketList(m.PacketList(m.PacketNumber(6))),
	} {
		if index_match := f.IndexOf(
			flattened_packets,
			func(packet m.Packet) bool { return m.PacketEqualityFunction(divider_packet, packet) },
		); index_match != -1 {
			return divider_packet_detected_final_validation_error(divider_packet, index_match+1)
		}
	}
	return nil
}

func (pr packet_reader) Done() bool {
	return pr.Error() != nil
}

func (pr *packet_reader) ProvideLine(line string) {
	pr.line_number++
	pr.line_readers[pr.current_status](line)
}

func (pr packet_reader) FinishAndGetInputData() e.Envelope[[]m.PacketPair] {
	return m.PacketPairsEnvelope(pr.packet_pairs...)
}

// packet reader's line readers

func (pr *packet_reader) readEmptyLine(line string) {
	if !pr.empty_re.MatchString(line) {
		pr.err = expected_empty_line_error(pr.line_number, line)
		return
	}
	pr.current_status = first_packet
}

func (pr *packet_reader) readFirstPacket(line string) {
	// If it is an empty line, it's ok: ignore it
	if pr.empty_re.MatchString(line) {
		return
	}

	if !pr.packet_re.MatchString(line) {
		pr.err = expected_first_packet_error(pr.line_number, line)
		return
	}

	var packet m.Packet
	packet, pr.err = parse_packet(pr.line_number, line)
	pr.packet_pairs = append(pr.packet_pairs, m.PacketPair{First: packet, Second: nil})
	pr.current_status = second_packet
}

func (pr *packet_reader) readSecondPacket(line string) {
	if pr.empty_re.MatchString(line) || !pr.packet_re.MatchString(line) {
		pr.err = expected_second_packet_error(pr.line_number, line)
		return
	}

	var packet m.Packet
	packet, pr.err = parse_packet(pr.line_number, line)
	pr.packet_pairs[len(pr.packet_pairs)-1].Second = packet
	pr.current_status = empty_line
}
