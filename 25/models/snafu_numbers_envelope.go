package models

import c "aoc/common"

type snafu_numbers_envelope struct {
	snafu_numbers []string
}

func CreateSnafuNumbersEnvelope(numbers []string) c.Envelope[[]string] {
	return snafu_numbers_envelope{numbers}
}

func (sne snafu_numbers_envelope) Get() []string {
	return c.ShallowCopy(sne.snafu_numbers)
}

func SnafuNumbersEnvelopeEqualityFunction(lhs, rhs c.Envelope[[]string]) bool {
	return c.ArrayEqual(lhs.Get(), rhs.Get())
}
