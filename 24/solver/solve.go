package solver

import (
	c "aoc/common"
	m "aoc/d24/models"
)

func solve_valley_problem(valley *m.Valley, starting_time int) (int, error) {
	starting_state := valley.GetStartingState(starting_time)
	discovered_states := map[m.State]bool{starting_state: true}

	pq := c.PriorityQueue(func(lhs, rhs c.Pair[m.State, int]) bool {
		return lhs.Second < rhs.Second
	})
	pq.Add(c.MakePair(starting_state, valley.GetEstimatedTime(starting_state)))

	is_record_set := false
	record := 0
	for !pq.IsEmpty() {
		element, _ := pq.Pop()
		state, estimate := element.Get()
		if is_record_set && record <= estimate {
			break
		}
		if state.CurrentPosition == valley.GetGoal() {
			if !is_record_set {
				record = state.PassedTime
			}
			is_record_set = true
			record = c.Min(record, state.PassedTime)
		}
		for _, next_state := range valley.GetNextStates(state) {
			// If we have already considering an equivalent state, do not add it to the queue
			if discovered_states[next_state] {
				continue
			}
			// If the estimate is worse than actually recorded time, do not consider this state
			next_estimate := valley.GetEstimatedTime(next_state)
			if is_record_set && record <= next_estimate {
				continue
			}
			pq.Add(c.MakePair(next_state, next_estimate))
			discovered_states[next_state] = true
		}
	}
	return record, nil
}

func CalculateSingleTripBestTime(envelope c.Envelope[[][]rune]) (int, error) {
	valley := m.CreateValley(envelope.Get())
	return solve_valley_problem(valley, 0)
}

func CalculateTripleTripBestTime(envelope c.Envelope[[][]rune]) (int, error) {
	valley := m.CreateValley(envelope.Get())
	first_trip_time, first_err := solve_valley_problem(valley, 0)
	if first_err != nil {
		return 0, first_err
	}

	valley.Invert()
	second_trip_time, second_err := solve_valley_problem(valley, first_trip_time)
	if second_err != nil {
		return 0, second_err
	}

	valley.Invert()
	return solve_valley_problem(valley, second_trip_time)
}

// func SolveValleyMapProblem(envelope c.Envelope[[][]rune]) (int, error) {
// 	valley := m.CreateValley(envelope.Get())
// 	starting_state := valley.GetStartingState()
// 	discovered_states := map[m.State]bool{starting_state: true}

// 	pq := c.PriorityQueue(func(lhs, rhs c.Pair[m.State, int]) bool {
// 		return lhs.Second < rhs.Second
// 	})
// 	pq.Add(c.MakePair(starting_state, valley.GetEstimatedTime(starting_state)))

// 	is_record_set := false
// 	record := 0
// 	for !pq.IsEmpty() {
// 		element, _ := pq.Pop()
// 		state, estimate := element.Get()
// 		if is_record_set && record <= estimate {
// 			break
// 		}
// 		if state.CurrentPosition == valley.GetGoal() {
// 			if !is_record_set {
// 				record = state.PassedTime
// 			}
// 			is_record_set = true
// 			record = c.Min(record, state.PassedTime)
// 		}
// 		for _, next_state := range valley.GetNextStates(state) {
// 			// If we have already considering an equivalent state, do not add it to the queue
// 			if discovered_states[next_state] {
// 				continue
// 			}
// 			// If the estimate is worse than actually recorded time, do not consider this state
// 			next_estimate := valley.GetEstimatedTime(next_state)
// 			if is_record_set && record <= next_estimate {
// 				continue
// 			}
// 			pq.Add(c.MakePair(next_state, next_estimate))
// 			discovered_states[next_state] = true
// 		}
// 	}

// 	return record, nil
// }
