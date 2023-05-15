package solver

import (
	m "aoc/day02/models"
	e "aoc/envelope"

	f "aoc/functional"
)

func CalculateScore(ri RoundInterpreter) func(e.Envelope[[]m.Round]) (int, error) {
	return func(envelope e.Envelope[[]m.Round]) (int, error) {
		return f.Sum(f.Map(ri.GetScore, envelope.Get())), nil
	}
}
