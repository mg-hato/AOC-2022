package models

type FieldMap interface {
	GetInitialPointer() Pointer
	UpdatePointer(Pointer, Instruction) Pointer
}
