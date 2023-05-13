package models

type Addition struct{}

func Add() Operator {
	return Addition{}
}

func (Addition) String() string {
	return "+"
}

func (Addition) apply(lhs, rhs int) int {
	return lhs + rhs
}
