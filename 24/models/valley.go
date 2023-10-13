package models

import c "aoc/common"

type Valley struct {
	positions map[Position]rune
	blizzards []Blizzard

	goal  Position
	start Position

	inverted bool

	allowed_positions map[int]map[Position]bool
}

func CreateValley(valley_map [][]rune) *Valley {

	// work out walkable positions
	walkable_tokens := c.CreateSet([]rune(".<>v^"), c.Identity[rune])
	positions := make(map[Position]rune)
	for row, valley_row := range valley_map {
		for col, token := range valley_row {
			if walkable_tokens[token] {
				positions[MakePosition(row, col)] = token
			}
		}
	}

	// work out blizzards
	position_set := c.CreateSet(c.GetKeys(positions), c.Identity[Position])
	blizzard_token_map := c.CreateKeyValueMap(Directions(), Direction.String, c.Identity[Direction])
	blizzards := make([]Blizzard, 0)
	for pos, token := range positions {
		if dir, is_blizzard := blizzard_token_map[string(token)]; is_blizzard {
			blizzards = append(blizzards, MakeBlizzard(pos, dir, position_set))
		}
	}
	valley := &Valley{
		positions: positions,
		blizzards: blizzards,
		inverted:  false,

		allowed_positions: map[int]map[Position]bool{},
	}
	valley.assign_start_and_goal_positions()
	return valley
}

func (v *Valley) assign_start_and_goal_positions() {
	row_comparison_func := func(lhs, rhs Position) bool { return lhs.Row() < rhs.Row() }
	v.start = c.MinimumBy(c.GetKeys(v.positions), row_comparison_func)
	v.goal = c.MaximumBy(c.GetKeys(v.positions), row_comparison_func)
}

func (v *Valley) Invert() {
	v.inverted = !v.inverted
}

func (v *Valley) GetGoal() Position {
	if v.inverted {
		return v.start
	} else {
		return v.goal
	}
}

func (v *Valley) GetStart() Position {
	if v.inverted {
		return v.goal
	} else {
		return v.start
	}
}

func (v Valley) GetStartingState(starting_time int) State {
	return MakeState(v.GetStart(), starting_time)
}

func (v *Valley) GetEstimatedTime(st State) int {
	row_diff := c.Abs(st.CurrentPosition.Row() - v.GetGoal().Row())
	col_diff := c.Abs(st.CurrentPosition.Column() - v.GetGoal().Column())
	return st.PassedTime + row_diff + col_diff
}

func (v *Valley) GetAllowedPositionsAtTime(time int) map[Position]bool {
	if _, is_cached := v.allowed_positions[time]; !is_cached {
		allowed := c.CreateSet(c.GetKeys(v.positions), c.Identity[Position])
		c.ForEach(func(b Blizzard) { allowed[b.GetBlizzardPositionAtTime(time)] = false }, v.blizzards)
		v.allowed_positions[time] = allowed
	}
	return v.allowed_positions[time]
}

func (v *Valley) GetNextStates(st State) []State {
	// If goal already reached, do not expand further states
	if st.CurrentPosition == v.GetGoal() {
		return []State{}
	}

	new_time := st.PassedTime + 1
	allowed_positions := v.GetAllowedPositionsAtTime(new_time)
	return c.Filter(
		func(new_state State) bool {
			return allowed_positions[new_state.CurrentPosition]
		},
		[]State{
			MakeState(st.CurrentPosition, new_time),
			MakeState(st.CurrentPosition.Move(North), new_time),
			MakeState(st.CurrentPosition.Move(East), new_time),
			MakeState(st.CurrentPosition.Move(South), new_time),
			MakeState(st.CurrentPosition.Move(West), new_time),
		},
	)
}
