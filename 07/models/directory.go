package models

import (
	f "aoc/functional"
	"fmt"
	"strings"
)

func MakePartialDirectory(name string) Item {
	return &Directory{
		name: name,
	}
}

type Directory struct {
	name   string
	parent *Directory
	items  map[string]Item
}

func (d Directory) String() string {
	return fmt.Sprintf(
		"[Dir %s: %s]",
		d.name,
		func() string {
			if len(d.items) > 0 {
				return strings.Join(f.GetKeys(d.items), ", ")
			}
			return "<empty>"
		}(),
	)
}

func (d Directory) GetName() string {
	return d.name
}

func (d Directory) GetParent() *Directory {
	return d.parent
}

func (d *Directory) setParent(parent *Directory) {
	d.parent = parent
}

func (d Directory) GetItems() []Item {
	return f.GetValues(d.items)
}

func (d *Directory) shallow_equal(item Item) bool {
	if other, ok := item.(*Directory); ok {
		return d.name == other.name
	}
	return false
}

func (d *Directory) shallow_copy() Item {
	return &Directory{name: d.name}
}
