package main

func SimulateRopeWithTailCount(n int) func([]Motion) (int, error) {
	return func(motions []Motion) (int, error) {
		ht := CreateHeadTail(n)
		for _, m := range motions {
			for s := 0; s < m.steps; s++ {
				ht.update(m.direction)
			}
		}
		return ht.getVisitedCount(), nil
	}
}
