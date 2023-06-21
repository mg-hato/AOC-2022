package common_test

import (
	c "aoc/common"
	ts "aoc/testers"
	"fmt"
	"testing"
)

func TestCommon_ArrayEqualities(t *testing.T) {
	ints1 := []int{1, 2, 3}
	ints2 := []int{1, 2, 3, 4, 5}
	ints3 := []int{4, 5, 1, 2, 3}

	ts.AssertEqual(t, c.ArrayEqual(ints1, ints2), false)
	ts.AssertEqual(t, c.ArrayEqual(ints3, ints2), false)
	ts.AssertEqual(t, c.ArrayEqual(ints1, ints2[:3]), true)

	ts.AssertEqual(t, c.ArrayEqualInAnyOrder(ints2, ints3), true)
	ts.AssertEqual(t, c.ArrayEqualInAnyOrder(ints1, ints3), false)
	ts.AssertEqual(t, c.ArrayEqualInAnyOrder(ints1, ints3[2:]), true)

	always_true := func(_, _ int) bool { return true }
	ts.AssertEqual(t, c.ArrayEqualWith(always_true)(ints2, ints3), true)
	ts.AssertEqual(t, c.ArrayEqualInAnyOrderWith(always_true)(ints2, ints3), true)
	ts.AssertEqual(t, c.ArrayEqualWith(always_true)(ints1, ints3), false)
	ts.AssertEqual(t, c.ArrayEqualInAnyOrderWith(always_true)(ints1, ints3), false)

	ignore_sign_equality := func(lhs, rhs int) bool { return lhs == rhs || lhs == -rhs }
	ints4 := []int{-1, -2, -3}
	ints5 := []int{-2, -1, -3}

	ts.AssertEqual(t, c.ArrayEqualWith(ignore_sign_equality)(ints4, ints1), true)
	ts.AssertEqual(t, c.ArrayEqualWith(ignore_sign_equality)(ints5, ints1), false)
	ts.AssertEqual(t, c.ArrayEqualInAnyOrderWith(ignore_sign_equality)(ints5, ints1), true)
	ts.AssertEqual(t, c.ArrayEqualInAnyOrderWith(ignore_sign_equality)(ints5, ints4), true)
}

func TestCommon_ArrayMap(t *testing.T) {
	integers := []int{1, 2, 3, 4}
	doubled := []int{2, 4, 6, 8}
	powers_of_2 := []int{4, 16, 64, 256}

	double_f := func(i int) int { return i * 2 }
	power_of_2_func := func(i int) int {
		p := 1
		for i > 0 {
			i--
			p *= 2
		}
		return p
	}

	ts.AssertEqualWithEqFunc(t, c.Map(double_f, integers), doubled, c.ArrayEqual[int])
	ts.AssertEqualWithEqFunc(t, c.Map(power_of_2_func, doubled), powers_of_2, c.ArrayEqual[int])

	strs := []string{"<4>", "<16>", "<64>", "<256>"}
	to_string_in_brackets := func(i int) string { return fmt.Sprintf("<%d>", i) }
	ts.AssertEqualWithEqFunc(t, c.Map(to_string_in_brackets, powers_of_2), strs, c.ArrayEqual[string])
}

func TestCommon_MatrixMap(t *testing.T) {
	mat := [][]int{
		{1, 2, 34},
		{7, 9},
		{10},
	}
	res := c.MatrixMap(func(i int) int { return 2 * i }, mat)
	exp := [][]int{
		{2, 4, 68},
		{14, 18},
		{20},
	}
	ts.AssertEqualWithEqFunc(t, res, exp, c.ArrayEqualWith(c.ArrayEqual[int]))
}

func TestCommon_Folds(t *testing.T) {
	integers := []int{1, 2, 3, 4}
	ts.AssertEqual(
		t,
		c.Foldl(func(res string, i int) string { return fmt.Sprintf("(%s <+> %d)", res, i) }, integers, "x"),
		"((((x <+> 1) <+> 2) <+> 3) <+> 4)",
	)
	ts.AssertEqual(
		t,
		c.Foldr(func(i int, res string) string { return fmt.Sprintf("(%s <+> %d)", res, i) }, integers, "x"),
		"((((x <+> 4) <+> 3) <+> 2) <+> 1)",
	)
}

func TestCommon_Flatten(t *testing.T) {
	mat := [][]int{
		{1, 2, 34},
		{7, 9},
		{10},
	}
	flat1 := []int{1, 2, 34, 7, 9, 10}
	ts.AssertEqualWithEqFunc(t, c.Flatten(mat), flat1, c.ArrayEqual[int])
}

func TestCommon_FlatMap(t *testing.T) {
	mat := [][]int{
		{1, 2, 34},
		{7, 9},
		{10},
	}
	flat1 := []int{2, 4, 68, 14, 18, 20}
	ts.AssertEqualWithEqFunc(
		t,
		c.FlatMap(
			func(row []int) []int { return c.Map(func(i int) int { return 2 * i }, row) },
			mat,
		),
		flat1,
		c.ArrayEqual[int],
	)
}
func TestCommon_Filter(t *testing.T) {

	nums := []int{2, 4, 68, 14, 18, 20}
	ts.AssertEqualWithEqFunc(
		t,
		c.Filter(c.InInclusiveRange(14, 50), nums),
		[]int{14, 18, 20},
		c.ArrayEqual[int],
	)
	ts.AssertEqualWithEqFunc(
		t,
		c.Filter(c.InInclusiveRange(1, 68), nums),
		nums,
		c.ArrayEqual[int],
	)
	ts.AssertEqualWithEqFunc(
		t,
		c.Filter(c.InInclusiveRange(68, 68), nums),
		[]int{68},
		c.ArrayEqual[int],
	)
	ts.AssertEqualWithEqFunc(
		t,
		c.Filter(c.InInclusiveRange(100, 200), nums),
		[]int{},
		c.ArrayEqual[int],
	)
}

func TestCommon_ForEach(t *testing.T) {

	nums := []int{2, 4, 68, 14, 18, 20}
	x := 0
	calls := 0
	c.ForEach(func(i int) {
		calls++
		x += 2 * i
	}, nums)

	ts.AssertEqual(t, calls, 6)
	ts.AssertEqual(t, x, 2*c.Sum(nums))
}

func TestCommon_TakeDrop(t *testing.T) {
	numbers := []int{1, 5, 7, 21, 35}
	ts.AssertEqualWithEqFunc(t, c.Take(2, numbers), []int{1, 5}, c.ArrayEqual[int])
	ts.AssertEqualWithEqFunc(t, c.Take(4, numbers), []int{1, 5, 7, 21}, c.ArrayEqual[int])
	ts.AssertEqualWithEqFunc(t, c.Take(10, numbers), numbers, c.ArrayEqual[int])
	ts.AssertEqualWithEqFunc(t, c.Take(-3, numbers), []int{}, c.ArrayEqual[int])

	ts.AssertEqualWithEqFunc(t, c.Drop(2, numbers), []int{7, 21, 35}, c.ArrayEqual[int])
	ts.AssertEqualWithEqFunc(t, c.Drop(4, numbers), []int{35}, c.ArrayEqual[int])
	ts.AssertEqualWithEqFunc(t, c.Drop(10, numbers), []int{}, c.ArrayEqual[int])
	ts.AssertEqualWithEqFunc(t, c.Drop(-3, numbers), numbers, c.ArrayEqual[int])
}

func TestCommon_AnyAll(t *testing.T) {
	numbers := []int{1, 5, 7, 21, 42}

	divby7 := func(i int) bool { return i%7 == 0 }
	even := func(i int) bool { return i%2 == 0 }
	positive := func(i int) bool { return i > 0 }
	negative := func(i int) bool { return i < 0 }
	single_digit := c.InRange(0, 10)

	ts.AssertEqual(t, c.All(divby7, numbers), false)
	ts.AssertEqual(t, c.Any(divby7, numbers), true)

	ts.AssertEqual(t, c.All(even, numbers), false)
	ts.AssertEqual(t, c.Any(even, numbers), true)

	ts.AssertEqual(t, c.All(positive, numbers), true)
	ts.AssertEqual(t, c.Any(positive, numbers), true)

	ts.AssertEqual(t, c.All(negative, numbers), false)
	ts.AssertEqual(t, c.Any(negative, numbers), false)

	ts.AssertEqual(t, c.All(single_digit, numbers), false)
	ts.AssertEqual(t, c.Any(single_digit, numbers), true)

	no_numbers := []int{}
	ts.AssertEqual(t, c.All(c.Const[int](true), no_numbers), true)
	ts.AssertEqual(t, c.Any(c.Const[int](true), no_numbers), false)

	ts.AssertEqual(t, c.All(c.Const[int](false), no_numbers), true)
	ts.AssertEqual(t, c.Any(c.Const[int](false), no_numbers), false)
}

func TestCommon_MaximumMinimumBy(t *testing.T) {
	numbers := []int{7, 21, 500, 42, 5, 77}

	ts.AssertEqual(t, c.MaximumBy(numbers, c.LessThan[int]), 500)
	ts.AssertEqual(t, c.MinimumBy(numbers, c.LessThan[int]), 5)

	ts.AssertEqual(t, c.MaximumBy(numbers, c.GreaterThan[int]), 5)
	ts.AssertEqual(t, c.MinimumBy(numbers, c.GreaterThan[int]), 500)

	nearer_to_100 := func(lhs, rhs int) bool {
		return c.Abs(100-lhs) <= c.Abs(100-rhs)
	}

	ts.AssertEqual(t, c.MaximumBy(numbers, nearer_to_100), 500)
	ts.AssertEqual(t, c.MinimumBy(numbers, nearer_to_100), 77)
}

func TestCommon_ArrayContains(t *testing.T) {
	numbers := []int{7, 21, 500, 42, 5, 77}

	ts.Assert(t, c.ArrayContains(numbers, 21))
	ts.Assert(t, c.ArrayContains(numbers, 77))
	ts.Assert(t, c.ArrayContains(numbers, 7))
	ts.Assert(t, !c.ArrayContains(numbers, 25))
}

func TestCommon_Zip(t *testing.T) {
	numbersA := []int{7, 21, 500, 42, 5, 77}
	numbersB := []int{14, 42, 1_000}

	ts.AssertEqualWithEqFunc(t, c.Zip(numbersA, numbersB), []c.Pair[int, int]{
		c.MakePair(7, 14),
		c.MakePair(21, 42),
		c.MakePair(500, 1_000),
	}, c.ArrayEqual[c.Pair[int, int]])

	ts.AssertEqualWithEqFunc(t, c.Zip(numbersB, numbersA[:2]), []c.Pair[int, int]{
		c.MakePair(14, 7),
		c.MakePair(42, 21),
	}, c.ArrayEqual[c.Pair[int, int]])
}

func TestCommon_Repeat(t *testing.T) {
	ts.AssertEqualWithEqFunc(t, c.Repeat("a", 0), []string{}, c.ArrayEqual[string])
	ts.AssertEqualWithEqFunc(t, c.Repeat("a", -10), []string{}, c.ArrayEqual[string])
	ts.AssertEqualWithEqFunc(t, c.Repeat("a", 3), []string{"a", "a", "a"}, c.ArrayEqual[string])
	ts.AssertEqualWithEqFunc(t, c.Repeat("a", 1), []string{"a"}, c.ArrayEqual[string])
}

func TestCommon_Ranges(t *testing.T) {
	ts.AssertEqualWithEqFunc(t, c.Range(5, 5), []int{}, c.ArrayEqual[int])
	ts.AssertEqualWithEqFunc(t, c.Range(5, 6), []int{5}, c.ArrayEqual[int])
	ts.AssertEqualWithEqFunc(t, c.Range(5, 10), []int{5, 6, 7, 8, 9}, c.ArrayEqual[int])
	ts.AssertEqualWithEqFunc(t, c.Range(10, 5), []int{}, c.ArrayEqual[int])

	ts.AssertEqualWithEqFunc(t, c.RangeInclusive(5, 5), []int{5}, c.ArrayEqual[int])
	ts.AssertEqualWithEqFunc(t, c.RangeInclusive(5, 6), []int{5, 6}, c.ArrayEqual[int])
	ts.AssertEqualWithEqFunc(t, c.RangeInclusive(5, 10), []int{5, 6, 7, 8, 9, 10}, c.ArrayEqual[int])
	ts.AssertEqualWithEqFunc(t, c.RangeInclusive(10, 5), []int{}, c.ArrayEqual[int])
	ts.AssertEqualWithEqFunc(t, c.RangeInclusive(3, 2), []int{}, c.ArrayEqual[int])
}

func TestCommon_Enumerate(t *testing.T) {
	strs := []string{"ab", "mnj", "jack", "tt"}

	ts.AssertEqualWithEqFunc(
		t,
		c.Enumerate(strs),
		[]c.Pair[int, string]{
			c.MakePair(0, "ab"),
			c.MakePair(1, "mnj"),
			c.MakePair(2, "jack"),
			c.MakePair(3, "tt"),
		},
		c.ArrayEqual[c.Pair[int, string]],
	)

	ts.AssertEqualWithEqFunc(
		t,
		c.EnumerateWithFirstIndex[string](7)(strs),
		[]c.Pair[int, string]{
			c.MakePair(7, "ab"),
			c.MakePair(8, "mnj"),
			c.MakePair(9, "jack"),
			c.MakePair(10, "tt"),
		},
		c.ArrayEqual[c.Pair[int, string]],
	)
}

func TestCommon_Count(t *testing.T) {
	divBy := func(d int) func(int) bool {
		return func(i int) bool { return i%d == 0 }
	}
	for _, divisor := range []int{7, 15, 31, 101} {
		ts.AssertEqual(
			t,
			c.Count(c.Range(0, 1000), divBy(divisor)),
			len(c.Filter(divBy(divisor), c.Range(0, 1000))),
		)
	}
}

func TestCommon_IndexOf(t *testing.T) {
	ts.AssertEqual(t, c.IndexOf([]int{10, 20, 100, 121}, c.InRange(100, 1000)), 2)
	ts.AssertEqual(t, c.IndexOf([]int{10, 20, 100, 121}, c.InRange(-5, -2)), -1)
	ts.AssertEqual(t, c.IndexOf([]int{10, 20, 100, 121}, c.InRange(120, 500)), 3)
}

func TestCommon_Reverse(t *testing.T) {

	ts.AssertEqualWithEqFunc(
		t,
		c.Reverse([]int{10, 20, 100}),
		[]int{100, 20, 10},
		c.ArrayEqual[int],
	)
	ts.AssertEqualWithEqFunc(
		t,
		c.Reverse([]int{10, 20, 100, 121}),
		[]int{121, 100, 20, 10},
		c.ArrayEqual[int],
	)
	ts.AssertEqualWithEqFunc(
		t,
		c.Reverse([]int{444}),
		[]int{444},
		c.ArrayEqual[int],
	)
	ts.AssertEqualWithEqFunc(
		t,
		c.Reverse([]int{}),
		[]int{},
		c.ArrayEqual[int],
	)
}
