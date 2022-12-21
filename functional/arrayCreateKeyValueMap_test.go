package functional

import (
	"fmt"
	"strings"
	"testing"
)

func TestCreateKeyValueMap(t *testing.T) {
	tests := []struct {
		arr         []int
		expectedMap map[int]string
	}{
		{[]int{1, 5, 3}, map[int]string{2: "11", 10: "55", 6: "33"}},
		{[]int{}, map[int]string{}},
		{[]int{7, 7, 7}, map[int]string{14: "77"}},
	}

	keyExtractor := func(i int) int { return i * 2 }
	valueExtractor := func(i int) string { return fmt.Sprintf("%d%d", i, i) }

	for test_number, test := range tests {
		if createdMap := CreateKeyValueMap(test.arr, keyExtractor, valueExtractor); !MapEqual(createdMap, test.expectedMap) {
			msg := strings.Join([]string{
				fmt.Sprintf("Test #%d failed on CreateKeyValueMap(%v, keyf, valf)", test_number+1, test.arr),
				fmt.Sprintf("\tCreated map: %v\n", createdMap),
				fmt.Sprintf("\tExpected map: %v\n", test.expectedMap),
			}, "")
			t.Error(msg)
		}
	}
}
