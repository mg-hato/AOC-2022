package main

import (
	. "aoc/functional"
	"fmt"
	"strings"
)

// A type of the command (either "ls" or "cd")
type CommandType int

const (
	ChangeDirectory CommandType = iota
	List
)

func (ct CommandType) String() string {
	switch ct {
	case ChangeDirectory:
		return "cd"
	default:
		return "ls"
	}
}

// A type of the item listed by "ls" command
type ListedItemType int

const (
	DirectoryType ListedItemType = iota
	FileType
)

func (it ListedItemType) String() string {
	switch it {
	case DirectoryType:
		return "dir"
	default:
		return "file"
	}
}

// Listed item represents one line of output of "ls" (a file with size or a directory)
type ListedItem struct {
	size      int
	item_type ListedItemType
	name      string
}

func (li ListedItem) String() string {
	str := fmt.Sprintf("%s(%s)", li.item_type, li.name)
	if li.item_type == FileType {
		str = fmt.Sprintf("%s[%d]", str, li.size)
	}
	return str
}

// Represents the command from the terminal in day 07
type Command struct {
	command_type CommandType
	argument     string
	listed_items []ListedItem
}

func (cmd Command) String() string {
	if cmd.command_type == ChangeDirectory {
		return fmt.Sprintf("%s %s;", cmd.command_type, cmd.argument)
	} else {
		listed_str := strings.Join(Map(func(li ListedItem) string { return fmt.Sprint(li) }, cmd.listed_items), " ")
		return fmt.Sprintf("%s <%s>;", cmd.command_type, listed_str)
	}
}

// File
type File struct {
	name string
	size int
}

// Directory
type Directory struct {
	parent   *Directory
	files    []File
	children map[string]*Directory
}
