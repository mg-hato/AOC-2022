package main

import (
	. "aoc/functional"
	"testing"
)

func TestTranspose(t *testing.T) {
	// Transposed flattened heights
	tfh := Map(func(t Tree) byte { return t.height }, Flatten(transpose(createForest([][]byte{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}))))

	if !ArrayEqual(tfh, []byte{1, 4, 7, 2, 5, 8, 3, 6, 9}) {
		t.Error("Test for transpose function failed")
	}
}
