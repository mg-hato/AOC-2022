package reader

import (
	c "aoc/common"
	m "aoc/d23/models"
	"aoc/reading"
	"fmt"
	"regexp"
)

type grove_map_reader struct {
	err         error
	line_number int

	grove_line_re *regexp.Regexp

	empty_re  *regexp.Regexp
	grove_map [][]m.SpotType
}

func GroveMapReader() reading.ReaderAoC2022[c.Envelope[[][]m.SpotType]] {
	return &grove_map_reader{
		grove_line_re: regexp.MustCompile(`^((?:\.|#)+) *$`),

		empty_re:  regexp.MustCompile(`^ *$`),
		grove_map: make([][]m.SpotType, 0),
	}
}

func (gmr grove_map_reader) Error() error {
	return gmr.err
}

func (gmr grove_map_reader) PerformFinalValidation() error {
	todo_err := fmt.Errorf("error: todo")
	if len(gmr.grove_map) == 0 {
		return todo_err // empty grove error
	}
	if !c.All(func(row []m.SpotType) bool { return len(gmr.grove_map[0]) == len(row) }, gmr.grove_map) {
		return todo_err // grove not square shaped
	}
	if c.Count(c.Flatten(gmr.grove_map), func(st m.SpotType) bool { return st == m.Elf }) == 0 {
		return todo_err // no elves in the grove error
	}
	return nil
}

func (gmr grove_map_reader) Done() bool {
	return gmr.Error() != nil
}

func (gmr *grove_map_reader) ProvideLine(line string) {
	gmr.line_number++
	switch {
	case gmr.empty_re.MatchString(line):
	case gmr.grove_line_re.MatchString(line):
		gmr.process_grove_line(line)
	default:
		gmr.err = fmt.Errorf("something something error") // bad line error
	}
}

func (gmr *grove_map_reader) process_grove_line(line string) {
	data := gmr.grove_line_re.FindStringSubmatch(line)[1]
	grove_row := c.Map(func(r rune) m.SpotType {
		st, _ := m.TryParseSpotType(r)
		return st
	}, []rune(data))
	gmr.grove_map = append(gmr.grove_map, grove_row)
}

func (gmr grove_map_reader) FinishAndGetInputData() c.Envelope[[][]m.SpotType] {
	return m.CreateGroveMapEnvelope(gmr.grove_map)
}
