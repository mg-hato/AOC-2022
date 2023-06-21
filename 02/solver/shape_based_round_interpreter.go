package solver

import m "aoc/d02/models"

type shape_based_round_interpreter struct{}

// Treats X, Y and Z as a shape: Rock, Paper and Scissors, respectively
func ShapeBasedRoundInterpreter() RoundInterpreter {
	return &shape_based_round_interpreter{}
}

func (shape_based_round_interpreter) String() string {
	return "ShapeBasedRoundInterpreter"
}

func (sbri shape_based_round_interpreter) GetScore(round m.Round) int {
	opponent_shape := m.Shape(int(round.Left))
	player_shape := m.Shape(int(round.Right))
	outcome := m.GameOutcome(int(player_shape)-int(opponent_shape)+4) % 3

	outcome_score := int(outcome) * 3
	shape_score := int(player_shape) + 1

	return outcome_score + shape_score
}
