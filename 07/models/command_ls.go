package models

import (
	f "aoc/functional"
	"fmt"
)

func MakeCommandLs(listed_items ...Item) Command {
	return Ls{listed_items}
}

type Ls struct {
	listed_items []Item
}

func (Ls) GetCommandType() CommandType {
	return LS
}

func (ls Ls) String() string {
	return fmt.Sprintf("ls %v", f.Map(Item.String, ls.listed_items))
}

func (ls Ls) Copy() Command {
	return MakeCommandLs(f.Map(Item.shallow_copy, ls.listed_items)...)
}

func (ls Ls) equal(command Command) bool {
	if other, ok := command.(Ls); ok {
		return f.ArrayEqualWith(ItemShallowEqualityFunction)(ls.listed_items, other.listed_items)
	}
	return false
}

func (ls Ls) apply(current *Directory) (*Directory, error) {
	// If the directory content has already been listed by any previous ls command:
	// Verify that the listings match (shallowly i.e. just name-wise & size-wise)
	if current.items != nil {
		if f.ArrayEqualInAnyOrderWith(Item.shallow_equal)(f.GetValues(current.items), ls.listed_items) {
			return current, nil
		} else {
			return nil, ls_items_do_not_match_error(current.name)
		}
	}

	// Otherwise, this is the first listing for the directory, so populate the items field accordingly
	current.items = f.CreateKeyValueMap(ls.listed_items, Item.GetName, Item.shallow_copy)
	f.ForEach(func(item Item) { item.setParent(current) }, f.GetValues(current.items))
	return current, nil
}
