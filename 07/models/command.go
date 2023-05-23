package models

type CommandType int

const (
	LS CommandType = iota
	CD
)

type Command interface {
	GetCommandType() CommandType
	String() string
	apply(*Directory) (*Directory, error)
	equal(Command) bool
	Copy() Command
}

func CommandEqualityFunc(lhs, rhs Command) bool {
	return lhs.equal(rhs)
}
