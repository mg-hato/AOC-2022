package models

import (
	f "aoc/functional"
)

func CreateFilesystem(commands []Command) (*Directory, error) {
	root := &Directory{
		name: "/",
	}
	root.parent = root
	current := root
	var err error = nil
	for i, command := range commands {
		current, err = command.apply(current)
		if err != nil {
			return nil, create_filesystem_error(i+1, command, err)
		}
	}
	return root, verify_filesystem(root)
}

func verify_filesystem(root *Directory) error {
	queue := []*Directory{root}
	var i int = 0
	for i < len(queue) {
		current := queue[i]
		if current.items == nil {
			return create_filesystem_verification_error(directory_is_unexplored_error(current.name))
		}
		subdirectories := f.Map(
			func(item Item) *Directory { d, _ := item.(*Directory); return d },
			f.Filter(
				func(item Item) bool { _, ok := item.(*Directory); return ok },
				f.GetValues(current.items),
			),
		)
		queue = append(queue, subdirectories...)
		i++
	}
	return nil
}
