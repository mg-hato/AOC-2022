package models

import (
	c "aoc/common"
)

type CaloryListEnvelope struct {
	calory_list CaloryList
}

func CreateCaloryListEnvelope(calory_list CaloryList) c.Envelope[CaloryList] {
	return CaloryListEnvelope{calory_list}
}

func (cle CaloryListEnvelope) Get() CaloryList {
	return c.Map(
		func(bag []int) []int {
			new_bag := make([]int, len(bag))
			copy(new_bag, bag)
			return new_bag
		},
		cle.calory_list,
	)
}

func CaloryListEqualityFunc(lhs, rhs c.Envelope[CaloryList]) bool {
	return c.ArrayEqualWith(c.ArrayEqual[int])(lhs.Get(), rhs.Get())
}
