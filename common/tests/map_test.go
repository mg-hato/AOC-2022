package common_test

import (
	c "aoc/common"
	ts "aoc/testers"
	"fmt"
	"testing"
)

func TestCommon_MapEqual(t *testing.T) {
	m1 := map[int]string{
		1:  "AA",
		2:  "222",
		10: "XX",
	}
	m2 := map[int]string{
		1:  "AA",
		2:  "222",
		10: "XX",
	}

	ts.AssertEqual(t, c.MapEqual(m1, m2), true)

	m2[10] = "Z"
	ts.AssertEqual(t, c.MapEqual(m1, m2), false)

	m1[10] = "Z"
	ts.AssertEqual(t, c.MapEqual(m1, m2), true)

	m1[7] = ""
	ts.AssertEqual(t, c.MapEqual(m1, m2), false)

	m2[7] = ""
	ts.AssertEqual(t, c.MapEqual(m1, m2), true)
}

func TestCommon_MapEqualWith(t *testing.T) {
	m1 := map[int][]int{
		1:  {1},
		2:  {2},
		10: {1, 0},
	}
	m2 := map[int][]int{
		1:  {1},
		2:  {2},
		10: {1, 0},
	}

	ts.AssertEqual(t, c.MapEqualWith[int](c.ArrayEqual[int])(m1, m2), true)

	delete(m2, 10)
	ts.AssertEqual(t, c.MapEqualWith[int](c.ArrayEqual[int])(m1, m2), false)

	delete(m1, 10)
	ts.AssertEqual(t, c.MapEqualWith[int](c.ArrayEqual[int])(m1, m2), true)

	m1[7] = []int{7}
	ts.AssertEqual(t, c.MapEqualWith[int](c.ArrayEqual[int])(m1, m2), false)

	m2[7] = []int{7}
	ts.AssertEqual(t, c.MapEqualWith[int](c.ArrayEqual[int])(m1, m2), true)

	m2[7] = []int{}
	ts.AssertEqual(t, c.MapEqualWith[int](c.ArrayEqual[int])(m1, m2), false)

}

func TestCommon_CreateKeyValueMap(t *testing.T) {
	m := c.CreateKeyValueMap(
		c.Range(0, 10),
		func(i int) int { return i * 2 },
		func(i int) string { return fmt.Sprint(i) },
	)

	exp := map[int]string{}
	for i := 0; i < 10; i++ {
		exp[2*i] = fmt.Sprint(i)
	}
	ts.AssertEqualWithEqFunc(t, m, exp, c.MapEqual[int, string])
}

func TestCommon_GroupBy(t *testing.T) {
	strs := []string{"aa", "BBB", "CD", "PV", "delta"}
	exp := map[int][]string{
		2: {"aa", "CD", "PV"},
		3: {"BBB"},
		5: {"delta"},
	}
	m := c.GroupBy(
		strs,
		func(s string) int { return len(s) },
		c.Identity[string],
	)

	ts.AssertEqualWithEqFunc(
		t,
		m,
		exp,
		c.MapEqualWith[int](c.ArrayEqual[string]),
	)
}
