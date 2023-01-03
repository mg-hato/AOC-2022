package main

func SimulateRopeWithTailCount(n int) func([]Motion) int {
	return func(motions []Motion) int {
		ht := CreateHeadTail(n)
		for _, m := range motions {
			for s := 0; s < m.steps; s++ {
				ht.update(m.direction)
			}
		}
		return ht.getVisitedCount()
	}
}
