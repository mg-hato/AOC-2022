package solver

import (
	c "aoc/common"
	m "aoc/d16/models"
)

func findAllDistances(valves []m.Valve) map[c.Pair[string, string]]int {
	distances := map[string]map[string]int{}
	for _, valve := range valves {
		distances[valve.ID] = map[string]int{valve.ID: 0}
		for _, tunnel := range valve.Tunnels {
			distances[valve.ID][tunnel] = 1
		}
	}

	done := false
	for !done {
		done = true
		updated_distances := make_replica_of_double_map(distances)

		for i := 0; i < len(valves); i++ {
			U := valves[i].ID

			for j := 0; j < len(valves); j++ {
				V := valves[j].ID

				if _, ok := distances[U][V]; i == j || !ok {
					continue
				}

				for k := 0; k < len(valves); k++ {
					W := valves[k].ID

					if _, ok := distances[V][W]; i == k || j == k || !ok {
						continue
					}

					d1 := distances[U][V]
					d2 := distances[V][W]

					d3, path3_ok := updated_distances[U][W]
					if !path3_ok || d1+d2 < d3 {
						updated_distances[U][W] = d1 + d2
						done = false
					}
				}
			}
		}

		distances = updated_distances
	}

	paired_distances := make(map[c.Pair[string, string]]int)
	for i := 0; i < len(valves); i++ {
		for j := 0; j < len(valves); j++ {
			paired_distances[c.MakePair(valves[i].ID, valves[j].ID)] = distances[valves[i].ID][valves[j].ID]
		}
	}

	return paired_distances
}

func make_replica_of_double_map(original map[string]map[string]int) map[string]map[string]int {
	replica := make(map[string]map[string]int)

	for k1, sub_map := range original {
		replica[k1] = map[string]int{}
		for k2, value := range sub_map {
			replica[k1][k2] = value
		}
	}

	return replica
}
