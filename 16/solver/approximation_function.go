package solver

// func approximation_function(S state) int {
// 	schedule := make_valve_opening_schedule(S, valves)
// 	if len(schedule) == 0 {
// 		return get_secured_pressure_release(S, deadline, valves)
// 	}

// 	distance := get_distance_to_closest_valve(S.current_valve, schedule, distances)

// 	// Calculate how much pressure will be released in the optimistic case
// 	index := 0
// 	time := S.current_minute + distance
// 	OAPR := 0 // optimistic additional pressure release
// 	for time <= deadline && index < len(schedule) {
// 		OAPR += schedule[index].Flow_rate * (deadline - time)
// 		index++
// 		time += 2 // +2 minutes = 1 minute to open the valve + 1 minute to move to the next scheduled valve
// 	}

// 	return OAPR + get_secured_pressure_release(S, deadline, valves)

// }

// func make_valve_opening_schedule(S state, valves map[string]m.Valve) []m.Valve {
// 	result := c.Filter(
// 		func(v m.Valve) bool { return v.Flow_rate > 0 && !S.opened_valves[v.ID] },
// 		c.GetValues(valves),
// 	)
// 	sort.Slice(result, func(i, j int) bool {
// 		return result[i].Flow_rate > result[j].Flow_rate
// 	})
// 	return result
// }

// func get_distance_to_closest_valve(
// 	current string,
// 	valves []m.Valve,
// 	distances map[c.Pair[string, string]]int,
// ) int {
// 	return c.Minimum(c.Map(func(valve m.Valve) int { return distances[c.MakePair(current, valve.ID)] }, valves))
// }
