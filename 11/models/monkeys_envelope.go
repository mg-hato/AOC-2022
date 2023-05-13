package models

import (
	e "aoc/envelope"
	f "aoc/functional"
	"fmt"
	"strings"
)

type MonkeysEnvelope struct {
	monkeys []Monkey
}

func (envelope MonkeysEnvelope) Get() []Monkey {
	return f.Map(func(monkey Monkey) Monkey {
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
			f.Map(func(m Monkey) string { return m.String() }, envelope.monkeys),
			", ",
		),
	)
}

func CreateMonkeysEnvelopeWith(monkeys []Monkey) e.Envelope[[]Monkey] {
	return MonkeysEnvelope{monkeys: monkeys}
}

func MonkeyEnvelopeEqFunc(lhs, rhs e.Envelope[[]Monkey]) bool {
	return f.ArrayEqualWith(MonkeyEqualityFunc)(lhs.Get(), rhs.Get())
}
