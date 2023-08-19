package common

type Pair[A, B any] struct {
	First  A
	Second B
}

func GetFirst[A, B any](pair Pair[A, B]) A {
	return pair.First
}

func GetSecond[A, B any](pair Pair[A, B]) B {
	return pair.Second
}

func (pair Pair[A, B]) Get() (A, B) {
	return pair.First, pair.Second
}

func MakePair[A, B any](first A, second B) Pair[A, B] {
	return Pair[A, B]{
		First:  first,
		Second: second,
	}
}
