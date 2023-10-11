package models

import "fmt"

type FieldType int

const (
	Blank FieldType = iota
	Dot
	Wall
)

func TryParseFieldType(field_type rune) (FieldType, error) {
	switch field_type {
	case ' ':
		return Blank, nil
	case '.':
		return Dot, nil
	case '#':
		return Wall, nil
	default:
		return FieldType(-1), fmt.Errorf(`could not parse unknown field type: "%c"`, field_type)
	}
}

func (ft FieldType) String() string {
	switch ft {
	case Blank:
		return "< >"
	case Dot:
		return "<.>"
	case Wall:
		return "<#>"
	default:
		return "[field-type:unknown]"
	}
}
