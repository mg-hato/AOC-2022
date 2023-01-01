package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func ReadTerminalOutput(filename string) ([]Command, error) {
	file, err := os.Open(filename)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	cmd_reader := NewCommandReader()

	for scanner.Scan() && cmd_reader.isOk() {
		cmd_reader.interpret(scanner.Text())
	}

	if !cmd_reader.isOk() {
		return nil, cmd_reader.err
	}

	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	return cmd_reader.commands, nil
}

type CommandReader struct {
	cd_regexp      *regexp.Regexp
	ls_regexp      *regexp.Regexp
	command_regexp *regexp.Regexp

	file_regexp *regexp.Regexp
	dir_regexp  *regexp.Regexp

	is_ls_output bool

	commands    []Command
	line_number int
	err         error
}

func NewCommandReader() CommandReader {
	return CommandReader{
		// Regular expression for "cd" command
		cd_regexp: regexp.MustCompile("^\\$ cd ((?:\\.\\.)|(?:/)|(?:[a-z\\.]+))$"),

		// Regular expression for "ls" command
		ls_regexp: regexp.MustCompile("^\\$ ls *$"),

		// Regular expression for a command: either "cd" or "ls" (starts with $)
		command_regexp: regexp.MustCompile("^\\$"),

		// Regular expression for "file" listed item
		file_regexp: regexp.MustCompile("^(\\d+) ([a-z\\.]+)$"),

		// Regular expression for "dir" listed item
		dir_regexp: regexp.MustCompile("^dir ([a-z\\.]+)$"),
	}
}

// Returns true iff no errors encountered
func (cmd_reader *CommandReader) isOk() bool {
	return cmd_reader.err == nil
}

// Interprets the line of input based on internal state of the reader & structure of the input-line
func (cmd_reader *CommandReader) interpret(line string) {

	cmd_reader.line_number++

	if cmd_reader.command_regexp.MatchString(line) {

		// If it is a command, interpret it

		cmd_reader.is_ls_output = false
		cmd_reader.interpretCommand(line)

	} else if cmd_reader.is_ls_output {

		// If it is not a command & we are still in "ls"-output sequence, interpret it as output
		cmd_reader.interpretListedItem(line)

	} else {

		// If none of the above, error
		cmd_reader.err = errors.New(fmt.Sprintf("Error: Command was expected on line %d. Line: \"%s\"", cmd_reader.line_number, line))
	}

}

// Interprets a line of input as a command
func (cmd_reader *CommandReader) interpretCommand(line string) {

	if args := cmd_reader.cd_regexp.FindStringSubmatch(line); len(args) == 2 {

		// 'cd' has only one input-string: 1 -> the name of the directory
		cmd_reader.addCommand(Command{ChangeDirectory, args[1], nil})

	} else if cmd_reader.ls_regexp.MatchString(line) {

		cmd_reader.addCommand(Command{List, "", nil})
		cmd_reader.is_ls_output = true

	} else {

		fmt.Println(len(args))
		// If neither "cd" nor "ls" command: error
		cmd_reader.err = errors.New(fmt.Sprintf("Error: Invalid command on line %d: \"%s\"", cmd_reader.line_number, line))
	}
}

// Add command to the list of processed commands
func (cmd_reader *CommandReader) addCommand(command Command) {
	cmd_reader.commands = append(cmd_reader.commands, command)
}

func (cmd_reader *CommandReader) interpretListedItem(line string) {

	if dir_name := cmd_reader.dir_regexp.FindStringSubmatch(line); len(dir_name) == 2 {

		// If it is directory, it only has a name
		cmd_reader.addListedItem(ListedItem{0, DirectoryType, dir_name[1]})

	} else if file_info := cmd_reader.file_regexp.FindStringSubmatch(line); len(file_info) == 3 {

		// File has two input-strings of value: 1st -> the size; 2nd -> the name
		filesize, _ := strconv.Atoi(file_info[1])
		cmd_reader.addListedItem(ListedItem{filesize, FileType, file_info[2]})

	} else {
		cmd_reader.err = errors.New(fmt.Sprintf("Error: Invalid listed item on line %d: \"%s\"", cmd_reader.line_number, line))
	}

}

func (cmd_reader *CommandReader) addListedItem(li ListedItem) {
	i := len(cmd_reader.commands) - 1
	cmd_reader.commands[i].listed_items = append(cmd_reader.commands[i].listed_items, li)
}
