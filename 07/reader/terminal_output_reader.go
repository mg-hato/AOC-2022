package reader

import (
	m "aoc/day07/models"
	e "aoc/envelope"
	f "aoc/functional"
	"aoc/reading"
	"regexp"
	"strconv"
)

type terminal_output_reader struct {
	err error

	line_number int

	current_reading_mode reading_mode
	line_processors      map[reading_mode]func(string)

	cd_re      *regexp.Regexp
	ls_re      *regexp.Regexp
	command_re *regexp.Regexp

	directory_re *regexp.Regexp
	file_re      *regexp.Regexp

	empty_re *regexp.Regexp

	commands []m.Command
	ls_items []m.Item
}

// Constructor function for section assignments list reader
func TerminalOutputReader() reading.ReaderAoC2022[e.Envelope[[]m.Command]] {
	tor := &terminal_output_reader{
		err: nil,

		line_number: 0,

		current_reading_mode: first_command,

		cd_re: regexp.MustCompile(`^ *\$ *cd +(/|\.\.|[a-z]+) *$`),
		ls_re: regexp.MustCompile(`^ *\$ *ls *$`),

		command_re: regexp.MustCompile(`^ *\$`),

		directory_re: regexp.MustCompile(`^ *dir +([a-z]+) *$`),
		file_re:      regexp.MustCompile(`^ *(0|[1-9][0-9]*) +([a-z]+(?:\.[a-z]+)*) *$`),

		empty_re: regexp.MustCompile(`^ *$`),

		commands: make([]m.Command, 0),
		ls_items: nil,
	}
	tor.line_processors = map[reading_mode]func(string){
		first_command: tor.processFirstCommand,
		new_command:   tor.processNewCommand,
		ls_command:    tor.processLsCommand,
	}
	return tor
}

func (tor terminal_output_reader) Error() error {
	return tor.err
}

func (tor terminal_output_reader) PerformFinalValidation() error {
	return nil
}

func (tor terminal_output_reader) Done() bool {
	return tor.Error() != nil
}

func (tor *terminal_output_reader) ProvideLine(line string) {
	tor.line_number++

	// If empty line, ignore it
	if !tor.empty_re.MatchString(line) {
		tor.line_processors[tor.current_reading_mode](line)
	}
}

func (tor terminal_output_reader) FinishAndGetInputData() e.Envelope[[]m.Command] {
	if tor.current_reading_mode == ls_command {
		return m.CreateCommandsEnvelope(append(tor.commands, m.MakeCommandLs(tor.ls_items...))...)
	} else {
		return m.CreateCommandsEnvelope(tor.commands...)
	}
}

func (tor *terminal_output_reader) processFirstCommand(line string) {
	matches := tor.cd_re.FindStringSubmatch(line)
	if matches == nil || matches[1] != "/" {
		tor.err = bad_first_command(tor.line_number, line)
		return
	}
	tor.commands = append(tor.commands, m.MakeCommandCd(matches[1]))
	tor.current_reading_mode = new_command
}

func (tor *terminal_output_reader) processNewCommand(line string) {
	cd_matches, ls_matches := tor.cd_re.FindStringSubmatch(line), tor.ls_re.FindStringSubmatch(line)
	if cd_matches == nil && ls_matches == nil {
		tor.err = bad_new_command(tor.line_number, line)
		return
	}

	if cd_matches != nil {
		tor.commands = append(tor.commands, m.MakeCommandCd(cd_matches[1]))
	} else {
		tor.ls_items = make([]m.Item, 0)
		tor.current_reading_mode = ls_command
	}
}

func (tor *terminal_output_reader) processLsCommand(line string) {
	// If new command started, save the ls listing read so far and switch reading mode
	if tor.command_re.MatchString(line) {
		tor.commands = append(tor.commands, m.MakeCommandLs(tor.ls_items...))
		tor.current_reading_mode = new_command
		tor.processNewCommand(line)
		return
	}

	dir_match, file_match := tor.directory_re.FindStringSubmatch(line), tor.file_re.FindStringSubmatch(line)
	if dir_match == nil && file_match == nil {
		tor.err = bad_item_ls_command(tor.line_number, line)
		return
	}

	var item m.Item
	if dir_match != nil {
		item = m.MakePartialDirectory(dir_match[1])
	} else {
		size, _ := strconv.Atoi(file_match[1])
		item = m.MakeFile(file_match[2], size)
	}

	if f.Any(func(name string) bool { return name == item.GetName() }, f.Map(m.Item.GetName, tor.ls_items)) {
		tor.err = duplicated_name_in_ls_items_listing(tor.line_number, item.GetName())
		return
	}

	tor.ls_items = append(tor.ls_items, item)
}
