package reader

import m "aoc/d21/models"

func validate_existence_of_given_id(id string) func(monkey_jobs_reader) error {
	return func(mjr monkey_jobs_reader) error {
		if _, id_exists := mjr.monkey_ids[id]; !id_exists {
			return monkey_id_missing_validation_error(id)
		}
		return nil
	}
}

func validate_no_self_reference(mjr monkey_jobs_reader) error {
	for _, job := range mjr.monkey_jobs {
		monkey_id := job.First
		for _, id := range extract_id_references(job.Second) {
			if monkey_id == id {
				return self_reference_validation_error(monkey_id, mjr.monkey_ids[monkey_id])
			}
		}
	}
	return nil
}

func validate_no_unknown_dependents(mjr monkey_jobs_reader) error {
	for _, job := range mjr.monkey_jobs {
		monkey_id := job.First
		for _, id := range extract_id_references(job.Second) {
			if _, is_known := mjr.monkey_ids[id]; !is_known {
				return unknown_reference_validation_error(monkey_id, mjr.monkey_ids[monkey_id], id)
			}
		}
	}
	return nil
}

func extract_id_references(job m.MonkeyJob) []string {
	id_refs := make([]string, 0)
	switch j := job.(type) {
	case m.SingleOperandJob:
		operand := j.GetOperand()
		if id, is_id := operand.(m.Identifier); is_id {
			id_refs = append(id_refs, id.GetId())
		}
	case m.TwoOperandJob:
		lhs, _, rhs := j.GetOperands()
		if id, is_id := lhs.(m.Identifier); is_id {
			id_refs = append(id_refs, id.GetId())
		}
		if id, is_id := rhs.(m.Identifier); is_id {
			id_refs = append(id_refs, id.GetId())
		}
	}
	return id_refs
}
