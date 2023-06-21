package models

import (
	c "aoc/common"
	"fmt"
	"strings"
)

type MonkeysEnvelope struct {
	monkeys []Monkey
}

func (envelope MonkeysEnvelope) Get() []Monkey {
	return c.Map(func(monkey Monkey) Monkey {
		new_monkey := monkey
		new_monkey.Items = make([]int, len(monkey.Items))
		copy(new_monkey.Items, monkey.Items)
		return new_monkey
	}, envelope.monkeys)
}

func (envelope MonkeysEnvelope) String() string {
	return fmt.Sprintf(
		"<<%s>>",
		strings.Join(
			c.Map(func(m Monkey) string { return m.String() }, envelope.monkeys),
			", ",
		),
	)
}

func CreateMonkeysEnvelopeWith(monkeys []Monkey) c.Envelope[[]Monkey] {
	return MonkeysEnvelope{monkeys: monkeys}
}

func MonkeyEnvelopeEqFunc(lhs, rhs c.Envelope[[]Monkey]) bool {
	return c.ArrayEqualWith(MonkeyEqualityFunc)(lhs.Get(), rhs.Get())
}
