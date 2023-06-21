package reader

import (
	m "aoc/d06/models"
	"aoc/reading"
	"fmt"
	"regexp"
)

type datastream_buffer_reader struct {
	line_number int

	done bool
	err  error

	datastream_re *regexp.Regexp
	empty_re      *regexp.Regexp

	datastream_buffer m.DatastreamBuffer
}

func DatastreamBufferReader() reading.ReaderAoC2022[m.DatastreamBuffer] {
	return &datastream_buffer_reader{
		line_number: 0,

		done: false,
		err:  nil,

		datastream_re: regexp.MustCompile("^[a-z]+$"),
		empty_re:      regexp.MustCompile("^ *$"),

		datastream_buffer: "",
	}
}

func (dbr datastream_buffer_reader) Error() error {
	return dbr.err
}

func (dbr datastream_buffer_reader) PerformFinalValidation() error {
	return nil
}

func (dbr datastream_buffer_reader) Done() bool {
	return dbr.done || dbr.err != nil
}

func (dbr *datastream_buffer_reader) ProvideLine(line string) {
	dbr.line_number++

	// Ignore empty lines
	if dbr.empty_re.MatchString(line) {
		return
	}

	// Otherwise, it must be the datastream buffer; produce an error if that is not the case
	if !dbr.datastream_re.MatchString(line) {
		dbr.err = fmt.Errorf(`Error: could not interpret line #%d: "%s"`, dbr.line_number, line)
		return
	}

	dbr.datastream_buffer = line
	dbr.done = true
}

func (dbr datastream_buffer_reader) FinishAndGetInputData() m.DatastreamBuffer {
	return dbr.datastream_buffer
}
