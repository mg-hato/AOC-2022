package models

type SingleOperandJob struct {
	operand Operand
}

func CreateSingleJob(operand Operand) MonkeyJob {
	return SingleOperandJob{operand}
}

func (job SingleOperandJob) GetOperand() Operand {
	return job.operand
}

func (job SingleOperandJob) String() string {
	return job.operand.String()
}

func (job SingleOperandJob) GetIdentifiers() []string {
	if id, is_id := job.operand.(Identifier); is_id {
		return []string{id.id}
	} else {
		return []string{}
	}
}

func (job SingleOperandJob) Equal(other MonkeyJob) bool {
	otherJob, ok := other.(SingleOperandJob)
	return ok && job.operand.Equal(otherJob.operand)
}

func (job SingleOperandJob) Calculate(id_resolver func(string) (int64, error)) (int64, error) {
	return job.operand.Resolve(id_resolver)
}
