package models

import (
	e "aoc/envelope"
	f "aoc/functional"
)

type CaloryListEnvelope struct {
	calory_list CaloryList
}

func CreateCaloryListEnvelope(calory_list CaloryList) e.Envelope[CaloryList] {
	return CaloryListEnvelope{calory_list}
}

func (cle CaloryListEnvelope) Get() CaloryList {
	return f.Map(
		func(bag []int) []int {
			new_bag := make([]int, len(bag))
			copy(new_bag, bag)
			return new_bag
		},
		cle.calory_list,
	)
}

func CaloryListEqualityFunc(lhs, rhs e.Envelope[CaloryList]) bool {
	return f.ArrayEqualWith(f.ArrayEqual[int])(lhs.Get(), rhs.Get())
}
