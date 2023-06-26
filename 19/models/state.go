package models

type State struct {
	OreRobotCount      int
	ClayRobotCount     int
	ObsidianRobotCount int
	GeodeRobotCount    int

	Ore  int
	Clay int

	Obsidian int
	Geode    int

	TimeLeft int

	Estimated int
}

func (state *State) GetGeodeCrackedUntilTimeout() int {
	return state.Geode + state.TimeLeft*state.GeodeRobotCount
}

func (state State) MoveTimeBy(amount int) State {
	state.TimeLeft -= amount

	state.Ore += amount * state.OreRobotCount
	state.Clay += amount * state.ClayRobotCount
	state.Obsidian += amount * state.ObsidianRobotCount
	state.Geode += amount * state.GeodeRobotCount

	return state
}
