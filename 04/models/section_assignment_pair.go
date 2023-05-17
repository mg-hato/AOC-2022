package models

import "fmt"

// An assigned pair of elves and their respective coverages
type SectionAssignmentPair struct {
	First, Second Coverage
}

func (sap SectionAssignmentPair) String() string {
	return fmt.Sprintf("<%s,%s>", sap.First, sap.Second)
}

// Helper function to simplify making of the section assignment pair
func MakeSAP(first_left, first_right, second_left, second_right int) SectionAssignmentPair {
	return SectionAssignmentPair{
		First: Coverage{
			Left:  first_left,
			Right: first_right,
		},
		Second: Coverage{
			Left:  second_left,
			Right: second_right,
		},
	}
}
