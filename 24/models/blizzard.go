package models

type Blizzard struct {
	positions []Position
}

func MakeBlizzard(origin Position, dir Direction, positions map[Position]bool) Blizzard {

	// find loopback position
	loopback_position := origin
	looopback_found := false
	for !looopback_found {
		next := loopback_position.Move(dir.Opposite())
		if !positions[next] {
			looopback_found = true
		} else {
			loopback_position = next
		}
	}

	// create blizzard positions array
	blizzard_positions := []Position{origin}
	done := false
	current := origin
	for !done {
		next := current.Move(dir)
		if !positions[next] {
			next = loopback_position
		}

		if next == origin {
			done = true
		} else {
			blizzard_positions = append(blizzard_positions, next)
			current = next
		}
	}

	return Blizzard{blizzard_positions}
}

func (b Blizzard) GetBlizzardPositionAtTime(time int) Position {
	index := time % len(b.positions)
	return b.positions[index]
}
