package models

type Multiplication struct{}

func (Multiplication) Apply(lhs, rhs int64) (int64, error) {
	return lhs * rhs, nil
}

func (Multiplication) String() string {
	return "*"
}

func (m Multiplication) ResolveRight(left, result int64) (int64, error) {
	if result == 0 || left == 0 || result%left != 0 {
		return 0, could_not_resolve_right(m, left, result)
	} else {
		return result / left, nil
	}
}

func (m Multiplication) ResolveLeft(right, result int64) (int64, error) {
	if result == 0 || right == 0 || result%right != 0 {
		return 0, could_not_resolve_left(m, right, result)
	} else {
		return result / right, nil
	}
}
