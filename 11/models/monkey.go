package models

import (
	f "aoc/functional"
	"fmt"
	"strings"
)

type Monkey struct {
	MonkeyId int

	Items []int

	InspectionOP InspectionOperation

	DivTest int
	OnTrue  int
	OnFalse int
}

func (m *Monkey) Receive(item int) {
	m.Items = append(m.Items, item)
}

func (m Monkey) PerformDivisionTest(item int) int {
	if item%m.DivTest == 0 {
		return m.OnTrue
	} else {
		return m.OnFalse
	}
}

func (m Monkey) String() string {
	return fmt.Sprintf(
		"{Monkey[id:%d <%s> OP(%s) divby(%d) ? %d : %d]",
		m.MonkeyId,
		strings.Join(f.Map(func(i int) string { return fmt.Sprint(i) }, m.Items), ","),
		m.InspectionOP.String(),
		m.DivTest, m.OnTrue, m.OnFalse,
	)
}

func MonkeyEqualityFunc(lhs, rhs Monkey) bool {
	return lhs.MonkeyId == rhs.MonkeyId &&
		lhs.DivTest == rhs.DivTest &&
		lhs.InspectionOP == rhs.InspectionOP &&
		lhs.OnFalse == rhs.OnFalse &&
		lhs.OnTrue == rhs.OnTrue &&
		f.ArrayEqual(lhs.Items, rhs.Items)
}
