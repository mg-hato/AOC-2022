package functional

func InRange(left, right int) func(int) bool {
	return func(n int) bool { return left <= n && n < right }
}

func InInclusiveRange(left, right int) func(int) bool {
	return func(n int) bool { return left <= n && n <= right }
}
