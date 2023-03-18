package main

import (
	. "aoc/functional"
	"sort"
)

func ForFilesystemGet(f func(*Directory, map[*Directory]int64) (int64, error)) func([]Command) (int64, error) {
	return func(commands []Command) (int64, error) {
		root := createFileSystem(commands)
		return f(root, getFileSizes(root))
	}
}

// Sums up directories whose sizes are <= 100K
func SumOfDirectoriesLte100k(root *Directory, sizes map[*Directory]int64) (int64, error) {
	return Sum(Filter(func(i int64) bool { return i <= 100_000 }, GetValues(sizes))), nil
}

// Finds size of the smallest update-enabling directory (i.e. directory whose deletion would free up enough space for the update)
func SmallestUpdateEnablingDirectorySize(root *Directory, sizes map[*Directory]int64) (int64, error) {
	// The amount of memory that needs to be freed up
	var missing_space int64 = sizes[root] - 40_000_000

	// If there is enough space, we do not need to delete anything
	if missing_space <= 0 {
		return 0, nil
	}

	candidate_sizes := Filter(func(i int64) bool { return i >= missing_space }, GetValues(sizes))
	sort.Slice(candidate_sizes, func(i, j int) bool {
		return candidate_sizes[i] < candidate_sizes[j]
	})
	return append(candidate_sizes, 0)[0], nil
}

func getFileSizes(root *Directory) map[*Directory]int64 {
	stack := []*Directory{root}
	visited := map[*Directory]bool{}
	sizes := map[*Directory]int64{}

	for len(stack) != 0 {
		next := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if visited[next] {
			sizes[next] = Sum(append(
				Map(func(f File) int64 { return int64(f.size) }, next.files),                   // sum of file-sizes in the current directory
				Map(func(d *Directory) int64 { return sizes[d] }, GetValues(next.children))..., // sum of "children" directories sizes
			))

		} else {
			visited[next] = true
			stack = append(stack, next)
			stack = append(stack, GetValues(next.children)...)
		}
	}

	return sizes
}

func createFileSystem(commands []Command) *Directory {
	root := Directory{children: make(map[string]*Directory)}
	root.parent = &root

	current := &root
	for _, cmd := range commands {
		current = current.update(cmd)
	}

	return &root
}

func (current *Directory) update(cmd Command) *Directory {
	if cmd.command_type == ChangeDirectory {
		switch cmd.argument {
		case "/":
			return current.getRoot()
		case "..":
			return current.parent
		default:
			return current.children[cmd.argument]
		}
	}

	for _, item := range cmd.listed_items {
		if item.item_type == DirectoryType {
			current.children[item.name] = &Directory{parent: current, children: make(map[string]*Directory)}
		} else {
			current.files = append(current.files, File{item.name, item.size})
		}
	}
	return current
}

func (current *Directory) getRoot() *Directory {
	for current != current.parent {
		current = current.parent
	}
	return current
}
