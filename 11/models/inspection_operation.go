package models

import "fmt"

type InspectionOperation struct {
	lhs, rhs Operand
	operator Operator
}

func IOP(lhs Operand, op Operator, rhs Operand) InspectionOperation {
	return InspectionOperation{
		lhs:      lhs,
		operator: op,
		rhs:      rhs,
	}
}

func (iop InspectionOperation) String() string {
	return fmt.Sprintf("%s %s %s", iop.lhs, iop.operator, iop.rhs)
}

func (iop InspectionOperation) Inspect(old_value int) int {
	return iop.operator.apply(iop.lhs.eval(old_value), iop.rhs.eval(old_value))
}
