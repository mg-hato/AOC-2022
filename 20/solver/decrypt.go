package solver

import (
	c "aoc/common"
	m "aoc/d20/models"
)

func Decrypt(mixings_required int, multiplier int64) func(m.SolverInput) (int64, error) {
	return func(envelope m.SolverInput) (int64, error) {

		AVL := m.MakeAVL_Tree()
		nodes := make([]*m.Node, len(envelope.Get()))
		for i, value := range c.Map(func(value int) int64 { return int64(value) * multiplier }, envelope.Get()) {
			nodes[i] = m.MakeNode(value)
			AVL.InsertAtIndex(AVL.Size(), nodes[i])
		}

		for mixings_done := 0; mixings_done < mixings_required; mixings_done++ {
			for i := 0; i < len(nodes); i++ {
				var next_position int64 = 0
				current_position := nodes[i].GetMyIndex()
				if size := int64(AVL.Size()); size > 2 {
					next_position = int64(nodes[i].GetMyIndex()) + nodes[i].GetValue()
					next_position = ((next_position % (size - 1)) + size - 1) % (size - 1)
				}
				AVL.RemoveIndex(current_position)
				AVL.InsertAtIndex(int(next_position), nodes[i])
			}
		}

		array := AVL.GetAsArray()
		zero_idx := c.IndexOf(array, func(v int64) bool { return v == 0 })
		coordinates := c.Map(
			func(offset int) int64 {
				return array[(zero_idx+offset)%len(array)]
			},
			[]int{1000, 2000, 3000},
		)
		return c.Sum(coordinates), nil
	}
}
