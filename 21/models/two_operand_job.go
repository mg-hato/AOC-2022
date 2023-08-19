package models

type TwoOperandJob struct {
	lhs, rhs Operand
	op       Operation
}

func CreateTwoOperandJob(lhs Operand, op Operation, rhs Operand) MonkeyJob {
	return TwoOperandJob{
		lhs: lhs,
		op:  op,
		rhs: rhs,
	}
}

func (job TwoOperandJob) GetOperands() (Operand, Operation, Operand) {
	return job.lhs, job.op, job.rhs
}

func (job TwoOperandJob) String() string {
	return PrintOperation(job.lhs.String(), job.op, job.rhs.String())
}

func (job TwoOperandJob) GetIdentifiers() []string {
	ids := []string{}
	if id, is_id := job.lhs.(Identifier); is_id {
		ids = append(ids, id.id)
	}
	if id, is_id := job.rhs.(Identifier); is_id {
		ids = append(ids, id.id)
	}
	return ids
}

func (job TwoOperandJob) Equal(other MonkeyJob) bool {
	otherJob, ok := other.(TwoOperandJob)
	return ok &&
		job.lhs.Equal(otherJob.lhs) &&
		job.rhs.Equal(otherJob.rhs) &&
		job.op == otherJob.op
}

func (job TwoOperandJob) Calculate(id_resolver func(string) (int64, error)) (int64, error) {
	lhs, lhs_err := job.lhs.Resolve(id_resolver)
	rhs, rhs_err := job.rhs.Resolve(id_resolver)
	result, result_err := job.op.Apply(lhs, rhs)
	for _, err := range []error{lhs_err, rhs_err, result_err} {
		if err != nil {
			return 0, err
		}
	}
	return result, nil
}

func (job TwoOperandJob) ChangeOperation(op Operation) MonkeyJob {
	return CreateTwoOperandJob(job.lhs, op, job.rhs)
}
