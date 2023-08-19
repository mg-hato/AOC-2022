package models

type Subtraction struct{}

func (Subtraction) Apply(lhs, rhs int64) (int64, error) {
	return lhs - rhs, nil
}

func (Subtraction) String() string {
	return "-"
}

func (Subtraction) ResolveRight(left, result int64) (int64, error) {
	return left - result, nil
}

func (Subtraction) ResolveLeft(right, result int64) (int64, error) {
	return right + result, nil
}
