package solver

import (
	m "aoc/day10/models"
	e "aoc/envelope"
	f "aoc/functional"
)

func SimulateProgram(analyser RegisterCapturingAnalyser) func(e.Envelope[[]m.Instruction]) (m.AnalyserReport, error) {
	return func(envelope e.Envelope[[]m.Instruction]) (m.AnalyserReport, error) {

		register := 1
		current_instruction_idx := 0
		instructions := envelope.Get()

		// Calculate when each instruction ends (in cycle time)
		instructions_end_times := f.Foldl(
			func(end_times []int, instruction m.Instruction) []int {
				end_time := instruction.Length()
				if len(end_times) > 0 {
					end_time += end_times[len(end_times)-1]
				}
				return append(end_times, end_time)
			},
			instructions,
			make([]int, 0),
		)

		analyser.Initialise()
		for !analyser.IsDone() {
			if current_instruction_idx == len(instructions) ||
				analyser.NextCycle() <= instructions_end_times[current_instruction_idx] {
				analyser.Capture(register)
			} else {
				register = instructions[current_instruction_idx].Execute(register)
				current_instruction_idx++
			}
		}

		return analyser.GenerateReport(), nil
	}
}
