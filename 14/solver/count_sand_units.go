package solver

import (
	m "aoc/d14/models"
	e "aoc/envelope"
)

func CountSandUnitsUntilStop(
	cave_system_simulator_provider func([]m.RockStructure) CaveSystemSimulator,
) func(e.Envelope[[]m.RockStructure]) (int, error) {
	return func(envelope e.Envelope[[]m.RockStructure]) (int, error) {
		cave_system_simulator := cave_system_simulator_provider(envelope.Get())
		sand_units_dropped := 0
		for {
			if cave_system_simulator.dropSandUnit() {
				sand_units_dropped++
			} else {
				return sand_units_dropped, nil
			}
		}
	}
}
