package models

type MonkeyJob interface {
	String() string
	Equal(MonkeyJob) bool
	GetIdentifiers() []string

	Calculate(func(string) (int64, error)) (int64, error)
}

func AreMonkeyJobsEqual(lhs, rhs MonkeyJob) bool {
	return lhs.Equal(rhs)
}
