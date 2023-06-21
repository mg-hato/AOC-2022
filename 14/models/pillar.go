package models

import c "aoc/common"

// Pillar represents the combination of rock-solid base that can have
// a number of sand blocks on top of it
type Pillar struct {
	sand_count int
	base       int
}

func MakePillar(base int) Pillar {
	return Pillar{base: base}
}

func (p *Pillar) AddSandBlock() {
	p.sand_count++
}

func (p Pillar) GetBase() int {
	return p.base
}

func (p Pillar) GetTop() int {
	return p.base - p.sand_count
}

func (p Pillar) ContainsDepth(depth int) bool {
	return c.InInclusiveRange(p.GetTop(), p.GetBase())(depth)
}
