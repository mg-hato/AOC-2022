package models

import "fmt"

type Rotation int

const (
	Left  Rotation = -1
	Right Rotation = 1
)

func (rot Rotation) Turn() int {
	switch rot {
	case Left:
		return -1
	case Right:
		return 1
	default:
		return 0
	}
}

func TryParseRotation(s string) (Rotation, error) {
	switch s {
	case "L", "l":
		return Left, nil
	case "R", "r":
		return Right, nil
	default:
		return Rotation(-1), fmt.Errorf(`could not parse rotation symbol: "%s"`, s)
	}
}

func (rot Rotation) String() string {
	switch rot {
	case Left:
		return "L"
	case Right:
		return "R"
	default:
		return "[Rotation:unknown]"
	}
}
