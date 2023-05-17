package solver

import m "aoc/day02/models"

type outcome_based_round_interpreter struct{}

// Treats X, Y and Z as a desired outcome: Lose, Draw and Win, respectively
func OutcomeBasedRoundInterpreter() RoundInterpreter {
	return &outcome_based_round_interpreter{}
}

func (outcome_based_round_interpreter) String() string {
	return "OutcomeBasedRoundInterpreter"
}

func (obri outcome_based_round_interpreter) GetScore(round m.Round) int {
	opponent_shape := m.Shape(int(round.Left))
	desired_outcome := m.GameOutcome(int(round.Right))

	player_shape := m.Shape((int(opponent_shape) + int(desired_outcome) + 2) % 3)

	outcome_score := int(desired_outcome) * 3
	shape_score := int(player_shape) + 1

	return outcome_score + shape_score
}
