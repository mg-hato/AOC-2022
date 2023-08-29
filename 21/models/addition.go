package models

type Addition struct{}

func (Addition) Apply(lhs, rhs int64) (int64, error) {
	return lhs + rhs, nil
}

func (Addition) String() string {
	return "+"
}

func (Addition) ResolveRight(left, result int64) (int64, error) {
	return result - left, nil
}

func (op Addition) ResolveLeft(right, result int64) (int64, error) {
	return op.ResolveRight(right, result)
}
