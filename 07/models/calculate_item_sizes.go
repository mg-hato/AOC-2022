package models

import c "aoc/common"

// Local type to denote a calculation step when calculating sizes of items in filesystem
type calculation_mode = int

const (
	EXPLORE   calculation_mode = iota // explore children of current directory
	CALCULATE                         // calculate size of current directory/file based on its children
)

func CalculateItemSizes(root *Directory) map[Item]int64 {
	size_mapping := make(map[Item]int64)

	calculation_queue := []c.Pair[Item, calculation_mode]{
		c.MakePair[Item](root, EXPLORE),
	}

	for len(calculation_queue) > 0 {
		current_node := calculation_queue[len(calculation_queue)-1]
		calculation_queue = calculation_queue[:len(calculation_queue)-1]
		calculation_queue = append(calculation_queue, handleCalculationNode(current_node, size_mapping)...)
	}
	return size_mapping
}

func handleCalculationNode(
	calculation_node c.Pair[Item, calculation_mode],
	size_mapping map[Item]int64,
) []c.Pair[Item, calculation_mode] {
	next_calculation_nodes := []c.Pair[Item, calculation_mode]{}
	switch item := calculation_node.First.(type) {
	case *File:
		size_mapping[item] = int64(item.size)
	case *Directory:
		if calculation_node.Second == EXPLORE {
			next_calculation_nodes = append(next_calculation_nodes, c.MakePair[Item](item, CALCULATE))
			next_calculation_nodes = append(next_calculation_nodes, c.Zip(
				item.GetItems(),
				c.Repeat(EXPLORE, len(item.items)),
			)...)
		} else {
			size_mapping[item] = c.Sum(c.Map(func(i Item) int64 { return size_mapping[i] }, item.GetItems()))
		}
	}
	return next_calculation_nodes
}
