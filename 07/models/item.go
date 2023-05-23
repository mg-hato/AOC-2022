package models

type Item interface {
	String() string

	GetParent() *Directory
	setParent(*Directory)

	GetName() string

	shallow_copy() Item
	shallow_equal(Item) bool
}

func ItemShallowEqualityFunction(lhs, rhs Item) bool {
	return lhs.shallow_equal(rhs)
}
