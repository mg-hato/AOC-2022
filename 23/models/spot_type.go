package models

import "fmt"

type SpotType int

const (
	Dot SpotType = iota
	Elf
)

func (st SpotType) String() string {
	switch st {
	case Dot:
		return "<.>"
	case Elf:
		return "<#>"
	default:
		return "<UNKNOWN_SPOT>"
	}
}

func TryParseSpotType(r rune) (SpotType, error) {
	switch r {
	case '.':
		return Dot, nil
	case '#':
		return Elf, nil
	default:
		return SpotType(-1), fmt.Errorf(`failed to parse spot type: unknown symbol "%c"`, r)
	}
}
