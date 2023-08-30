package models

import (
	"aoc/common"
	"fmt"
	"sort"
)

type Field struct {
	FType       FieldType
	Row, Column int
}

func (f Field) String() string {
	return fmt.Sprintf("(%s, %d, %d)", f.FType.String(), f.Row, f.Column)
}

func (f Field) GetColumn() int {
	return f.Column
}

func (f Field) GetRow() int {
	return f.Row
}

func (f Field) ToPosition() Position {
	return MakePosition(f.Row, f.Column)
}

// Sorts fields from North to South, West to East
func sort_fields(fields []Field) []Field {
	fields = common.ShallowCopy(fields)
	sort.Slice(fields, func(i, j int) bool {
		return fields[i].Row < fields[j].Row || (fields[i].Row == fields[j].Row && fields[i].Column < fields[j].Column)
	})
	return fields
}
