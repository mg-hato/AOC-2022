package models

import (
	e "aoc/envelope"
	f "aoc/functional"
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
		strings.Join(f.Map(func(r Round) string { return r.String() }, re.rounds), ", "),
	)
}

// Constructor function for RoundsEnvelope
func CreateRoundsEnvelope(rounds []Round) e.Envelope[[]Round] {
	return RoundsEnvelope{rounds}
}

func RoundsEnvelopeEqualityFunction(lhs, rhs e.Envelope[[]Round]) bool {
	return f.ArrayEqual(lhs.Get(), rhs.Get())
}
