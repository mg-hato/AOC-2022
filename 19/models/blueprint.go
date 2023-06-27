package models

import c "aoc/common"

type Blueprint struct {
	ID int

	OreRobot      ore_robot
	ClayRobot     clay_robot
	ObsidianRobot obsidian_robot
	GeodeRobot    geode_robot
}

func MakeBlueprint(id, ore4ore, ore4clay, ore4obs, clay4obs, ore4geo, obs4geo int) Blueprint {
	return Blueprint{
		ID: id,

		OreRobot:  ore_robot{ore4ore},
		ClayRobot: clay_robot{ore4clay},
		ObsidianRobot: obsidian_robot{
			Ore:  ore4obs,
			Clay: clay4obs,
		},
		GeodeRobot: geode_robot{
			Ore:      ore4geo,
			Obsidian: obs4geo,
		},
	}
}

func (bp *Blueprint) MakeInitialState(time int) State {
	state := State{
		OreRobotCount: 1,
		TimeLeft:      time,
	}
	state.Estimated = bp.approximate_geode(state)
	return state
}

func (blueprint *Blueprint) GetNextStateFunction() func(State) []State {
	return func(state State) []State {
		next_states := []State{}

		if state.OreRobotCount > 0 {
			next_states = append(
				next_states,
				blueprint.MakeOreRobot(state),
				blueprint.MakeClayRobot(state),
			)

			if state.ClayRobotCount > 0 {
				next_states = append(next_states, blueprint.MakeObsidianRobot(state))
			}
			if state.ObsidianRobotCount > 0 {
				next_states = append(next_states, blueprint.MakeGeodeRobot(state))
			}
		}

		next_states = c.Filter(func(s State) bool { return s.TimeLeft > 0 }, next_states)
		for i := range next_states {
			next_states[i].Estimated = blueprint.approximate_geode(next_states[i])
		}
		return next_states
	}
}

func (bp *Blueprint) approximate_geode(state State) int {
	ore := make([]int, state.TimeLeft+1)
	ore[0] = state.Ore
	ore_rc := state.OreRobotCount
	for i := 1; i < len(ore); i++ {
		ore_req := bp.OreRobot.Ore
		rc_inc := ore_rc + 1 - state.OreRobotCount
		if ore[i-1] >= ore_req*rc_inc {
			ore_rc++
		}
		ore[i] = ore[i-1] + ore_rc
	}

	clay := make([]int, state.TimeLeft+1)
	clay[0] = state.Clay
	clay_rc := state.ClayRobotCount
	for i := 1; i < len(clay); i++ {
		ore_req := bp.ClayRobot.Ore
		rc_inc := clay_rc + 1 - state.ClayRobotCount
		if ore[i-1] >= ore_req*rc_inc {
			clay_rc++
		}
		clay[i] = clay[i-1] + clay_rc
	}

	obsidian := make([]int, state.TimeLeft+1)
	obsidian[0] = state.Obsidian
	obsidian_rc := state.ObsidianRobotCount
	for i := 1; i < len(obsidian); i++ {
		ore_req, clay_req := bp.ObsidianRobot.Ore, bp.ObsidianRobot.Clay
		rc_inc := obsidian_rc + 1 - state.ObsidianRobotCount
		if ore[i-1] >= ore_req*rc_inc && clay[i-1] >= clay_req*rc_inc {
			obsidian_rc++
		}
		obsidian[i] = obsidian[i-1] + obsidian_rc
	}

	geode_approx, geode_rc := state.Geode, state.GeodeRobotCount
	for i := 1; i < state.TimeLeft+1; i++ {
		ore_req, obs_req := bp.GeodeRobot.Ore, bp.GeodeRobot.Obsidian
		rc_inc := geode_rc + 1 - state.GeodeRobotCount
		if ore[i-1] >= ore_req*rc_inc && obsidian[i-1] >= obs_req*rc_inc {
			geode_rc++
		}
		geode_approx += geode_rc
	}

	return geode_approx
}

func (blueprint *Blueprint) MakeOreRobot(state State) State {
	state = state.MoveTimeBy(1 + get_wait_time(blueprint.OreRobot.Ore, state.Ore, state.OreRobotCount))
	state.OreRobotCount++
	state.Ore -= blueprint.OreRobot.Ore
	return state
}

func (blueprint *Blueprint) MakeClayRobot(state State) State {
	state = state.MoveTimeBy(1 + get_wait_time(blueprint.ClayRobot.Ore, state.Ore, state.OreRobotCount))
	state.ClayRobotCount++
	state.Ore -= blueprint.ClayRobot.Ore
	return state
}

func (blueprint *Blueprint) MakeObsidianRobot(state State) State {
	state = state.MoveTimeBy(c.Max(
		1+get_wait_time(blueprint.ObsidianRobot.Clay, state.Clay, state.ClayRobotCount),
		1+get_wait_time(blueprint.ObsidianRobot.Ore, state.Ore, state.OreRobotCount),
	))
	state.ObsidianRobotCount++
	state.Clay -= blueprint.ObsidianRobot.Clay
	state.Ore -= blueprint.ObsidianRobot.Ore
	return state
}

func (blueprint *Blueprint) MakeGeodeRobot(state State) State {
	state = state.MoveTimeBy(c.Max(
		1+get_wait_time(blueprint.GeodeRobot.Ore, state.Ore, state.OreRobotCount),
		1+get_wait_time(blueprint.GeodeRobot.Obsidian, state.Obsidian, state.ObsidianRobotCount),
	))
	state.Ore -= blueprint.GeodeRobot.Ore
	state.Obsidian -= blueprint.GeodeRobot.Obsidian
	state.GeodeRobotCount++
	return state
}

func get_wait_time(required, current, production int) int {
	if current >= required {
		return 0
	}
	missing := required - current
	if missing%production == 0 {
		return missing / production
	} else {
		return (missing / production) + 1
	}
}
