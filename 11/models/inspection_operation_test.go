package models

import (
	"aoc/testers"
	"testing"
)

func TestD11_InspectionOperationTest(t *testing.T) {

	type old_and_new_val struct {
		old_val int
		new_val int
	}

	type test_input struct {
		op     InspectionOperation
		inputs []old_and_new_val
	}

	testers.TestThat([]test_input{
		{
			op:     IOP(Old(), Mult(), Old()),
			inputs: []old_and_new_val{{10, 100}, {7, 49}, {11, 121}},
		},
		{
			op:     IOP(Old(), Mult(), Num(21)),
			inputs: []old_and_new_val{{10, 210}, {7, 147}, {21, 441}},
		},
		{
			op:     IOP(Num(21), Add(), Old()),
			inputs: []old_and_new_val{{100, 121}, {7, 28}, {10_000, 10_021}},
		},
		{
			op:     IOP(Num(10), Add(), Num(21)),
			inputs: []old_and_new_val{{100, 31}, {1, 31}, {10_000, 31}},
		},
		{
			op:     IOP(Num(9), Mult(), Num(5)),
			inputs: []old_and_new_val{{1, 45}, {1_001, 45}},
		},
	}, func(ti test_input) {
		for _, input := range ti.inputs {
			testers.AssertEqual(t, ti.op.Inspect(input.old_val), input.new_val)
		}
	})

}
