package reader

import (
	c "aoc/common"
	m "aoc/d16/models"
)

func verify_starting_valve_exists(valves []m.Valve) error {
	valve_ids := c.Map(func(v m.Valve) string { return v.ID }, valves)
	if !c.ArrayContains(valve_ids, "AA") {
		return starting_valve_missing_error()
	}
	return nil
}

func verify_that_valves_have_unique_names(valves []m.Valve) error {
	valve_ids := map[string]int{}

	for i, valve_id := range c.Map(func(v m.Valve) string { return v.ID }, valves) {
		if valve_number, already_defined := valve_ids[valve_id]; already_defined {
			return valve_defined_twice_validation_error(valve_id, valve_number, i+1)
		} else {
			valve_ids[valve_id] = i + 1
		}
	}

	return nil
}

func verify_that_no_valve_tunnel_leads_to_itself(valves []m.Valve) error {
	for _, valve := range valves {
		if c.ArrayContains(valve.Tunnels, valve.ID) {
			return valve_has_self_loop_validation_error(valve.ID)
		}
	}
	return nil
}

func verify_that_all_valve_tunnels_lead_to_distinct_valve(valves []m.Valve) error {
	for _, valve := range valves {
		tunnels := map[string]bool{}
		for _, tunnel := range valve.Tunnels {
			if tunnels[tunnel] {
				return valve_has_two_tunnels_leading_to_the_same_valve(valve.ID, tunnel)
			} else {
				tunnels[tunnel] = true
			}
		}
	}
	return nil
}

func verify_that_valve_tunnels_lead_to_defined_valves(valves []m.Valve) error {
	valve_ids := c.CreateSet(valves, func(v m.Valve) string { return v.ID })

	for _, valve := range valves {
		for _, tunnel := range valve.Tunnels {
			if !valve_ids[tunnel] {
				return valve_has_a_tunnel_leading_to_undefined_valve_validation_error(valve.ID, tunnel)
			}
		}
	}
	return nil
}
