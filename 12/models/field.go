package models

type Field struct {
	Position   Position
	HeightCode rune
}

// Get height
func (field Field) GetHeight() int {
	switch {
	case field.HeightCode == 'E':
		return int('z') - int('a')
	case field.HeightCode == 'S':
		return 0
	default:
		return int(field.HeightCode) - int('a')
	}
}
