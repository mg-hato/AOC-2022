package main

func SignalStrengths(instructions []Instruction) Result {
	reg := 1
	cycle := 0
	cycles_of_interest := []int{20, 60, 100, 140, 180, 220}

	acc := 0

	var i int
	for len(cycles_of_interest) != 0 && i < len(instructions) {
		if cycle+instructions[i].CycleLength() >= cycles_of_interest[0] {
			acc += cycles_of_interest[0] * reg
			cycles_of_interest = cycles_of_interest[1:]
		} else {
			reg += instructions[i].DeltaX()
			cycle += instructions[i].CycleLength()
			i++
		}
	}
	return ResultInt{result: acc}
}

func DrawCRT(instructions []Instruction) Result {
	reg := 1
	cycle := 0
	i := 0

	finish_time := 240
	if len(instructions) > 0 {
		finish_time = instructions[i].CycleLength()
	}

	crt := make([]rune, 240)

	for cycle < 240 {
		if cycle == finish_time {
			reg += instructions[i].DeltaX()
			finish_time = 240
			i++
			if i < len(instructions) {
				finish_time = cycle + instructions[i].CycleLength()
			}
		} else {

			if c := cycle % 40; reg-1 <= c && c <= reg+1 {
				crt[cycle] = '#'
			} else {
				crt[cycle] = '.'
			}
			cycle++
		}
	}

	return ResultCRT{crt: string(crt)}
}
