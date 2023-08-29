package models

import c "aoc/common"

type monkey_jobs_envelope struct {
	jobs []c.Pair[string, MonkeyJob]
}

func CreateMonkeyJobsEnvelope(jobs ...c.Pair[string, MonkeyJob]) c.Envelope[[]c.Pair[string, MonkeyJob]] {
	return monkey_jobs_envelope{jobs}
}

func (jobs_envelope monkey_jobs_envelope) Get() []c.Pair[string, MonkeyJob] {
	return c.ShallowCopy(jobs_envelope.jobs)
}

func MokeyJobsEnvelopeEqualityFunc(lhs, rhs c.Envelope[[]c.Pair[string, MonkeyJob]]) bool {
	return c.ArrayEqualWith(func(left, right c.Pair[string, MonkeyJob]) bool {
		return left.First == right.First && AreMonkeyJobsEqual(left.Second, right.Second)
	})(lhs.Get(), rhs.Get())
}
