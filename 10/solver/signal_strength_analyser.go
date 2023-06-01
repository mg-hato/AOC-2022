package solver

import (
	"aoc/day10/models"
	"sort"
)

func SignalStrengthAnalyser(cycles ...int) RegisterCapturingAnalyser {
	sort.Slice(cycles, func(i, j int) bool { return cycles[i] < cycles[j] })
	return &signal_strength_analyser{
		cycles: cycles,
	}
}

type signal_strength_analyser struct {
	cycles   []int
	current  int
	strength int64
}

func (ssa *signal_strength_analyser) Initialise() {
	ssa.current = 0
	ssa.strength = 0
}

func (ssa signal_strength_analyser) IsDone() bool {
	return ssa.current == len(ssa.cycles)
}

func (ssa *signal_strength_analyser) Capture(register int) {
	if !ssa.IsDone() {
		ssa.strength += int64(ssa.cycles[ssa.current]) * int64(register)
		ssa.current++
	}
}

func (ssa *signal_strength_analyser) NextCycle() int {
	if ssa.IsDone() {
		return -1
	} else {
		return ssa.cycles[ssa.current]
	}
}

func (ssa signal_strength_analyser) GenerateReport() models.AnalyserReport {
	if ssa.IsDone() {
		return models.SignalStrengthReport(ssa.strength)
	}
	return nil
}
