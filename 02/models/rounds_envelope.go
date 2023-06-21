package models

import (
	c "aoc/common"
	"fmt"
	"strings"
)

type RoundsEnvelope struct {
	rounds []Round
}

func (re RoundsEnvelope) Get() []Round {
	new_rounds := make([]Round, len(re.rounds))
	copy(new_rounds, re.rounds)
	return new_rounds
}

func (re RoundsEnvelope) String() string {
	return fmt.Sprintf("RoundsEnvelope[%s]",
		strings.Join(c.Map(func(r Round) string { return r.String() }, re.rounds), ", "),
	)
}

// Constructor function for RoundsEnvelope
func CreateRoundsEnvelope(rounds []Round) c.Envelope[[]Round] {
	return RoundsEnvelope{rounds}
}

func RoundsEnvelopeEqualityFunction(lhs, rhs c.Envelope[[]Round]) bool {
	return c.ArrayEqual(lhs.Get(), rhs.Get())
}
