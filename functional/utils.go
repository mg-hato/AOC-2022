package functional

func InRange(n, left, right int) bool {
	return left <= n && n < right
}

func InInclusiveRange(n, left, right int) bool {
	return left <= n && n <= right
}
