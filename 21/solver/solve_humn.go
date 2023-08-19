package solver

import (
	c "aoc/common"
	m "aoc/d21/models"
	"errors"
	"fmt"
)

func SolveHumn(envelope m.SolverInput) (int64, error) {
	graph := make_graph(envelope.Get())
	humn_dependents := get_dependents_of("humn", graph)
	if !humn_dependents["root"] {
		return 0, errors.New(`solver error: "root" does not depend on "humn"`)
	}

	graph["root"] = graph["root"].(m.TwoOperandJob).ChangeOperation(m.Subtraction{})
	cached_solver := create_cached_solver(graph)

	var current_value int64 = 0
	current_id := "root"
	for current_id != "humn" {

		// Get count of "humn" dependents at the current ID's job
		humn_dependents_count := c.Count(
			graph[current_id].GetIdentifiers(),
			func(s string) bool { return humn_dependents[s] },
		)

		// Verify that there is exactly one operand that depends on "humn"
		if humn_dependents_count != 1 {
			return 0, fmt.Errorf(
				`solver error: at "%s" the number of humn dependents (%d) is not supported for solving`,
				current_id, humn_dependents_count,
			)
		}

		// The job at current ID must be a two operand job with two identifiers as operands
		lop, op, rop := graph[current_id].(m.TwoOperandJob).GetOperands()
		lid, rid := lop.(m.Identifier).GetId(), rop.(m.Identifier).GetId()

		var errs []error

		// Figure out left identifier's value
		if humn_dependents[lid] {
			right_value, e1 := cached_solver(rid)
			value, e2 := op.ResolveLeft(right_value, current_value)
			errs = []error{e1, e2}
			current_value, current_id = value, lid
		} else /* Figure out right identifier's value */ {
			left_value, e1 := cached_solver(lid)
			value, e2 := op.ResolveRight(left_value, current_value)
			errs = []error{e1, e2}
			current_value, current_id = value, rid
		}

		// Check if any of the steps produced an error
		if coalesced_err := coalesce_errors(errs); coalesced_err != nil {
			return 0, coalesced_err
		}
	}
	return current_value, nil
}

// returns first non-nil error
func coalesce_errors(errs []error) error {
	return c.Foldl(func(acc_err, e error) error {
		if acc_err == nil {
			return e
		} else {
			return acc_err
		}
	}, errs, nil)
}
