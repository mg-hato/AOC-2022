package models

import "fmt"

func MakeFile(name string, size int) Item {
	return &File{
		name: name,
		size: size,
	}
}

type File struct {
	name   string
	parent *Directory
	size   int
}

func (f File) String() string {
	return fmt.Sprintf("File(%s: %d)", f.name, f.size)
}

func (f File) GetName() string {
	return f.name
}

func (f File) GetParent() *Directory {
	return f.parent
}

func (f *File) setParent(parent *Directory) {
	f.parent = parent
}

func (f File) GetSize() int {
	return f.size
}

func (f *File) shallow_equal(item Item) bool {
	if other, ok := item.(*File); ok {
		return f.name == other.name && f.size == other.size
	}
	return false
}

func (f *File) shallow_copy() Item {
	return MakeFile(f.name, f.size)
}
