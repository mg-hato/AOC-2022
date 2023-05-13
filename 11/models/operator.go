package models

type Operator interface {
	String() string
	apply(int, int) int
}
