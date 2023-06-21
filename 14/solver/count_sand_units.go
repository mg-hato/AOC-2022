package solver

import (
	m "aoc/d14/models"
)

func CountSandUnitsUntilStop(
	cave_system_simulator_provider func([]m.RockStructure) CaveSystemSimulator,
) func(m.SolverInput) (int, error) {
	return func(input m.SolverInput) (int, error) {
		cave_system_simulator := cave_system_simulator_provider(input.Get())
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
