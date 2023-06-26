package models

import (
	c "aoc/common"
	"strings"
)

type chamber struct {
	stopped_rocks         map[Point]bool
	width                 int
	height                int
	registered_closures   map[closure]closure_info
	optimised_height_skip int64
}

type closure struct {
	jetpattern_index int
	shape_index      int
	encoded_closure  string
}

type closure_info struct {
	height     int
	rock_index int64
}

func EmptyChamber(width int) *chamber {
	return &chamber{
		stopped_rocks:       make(map[Point]bool),
		registered_closures: map[closure]closure_info{},
		width:               width,
	}
}

func (ch *chamber) is_point_occuppied(point Point) bool {
	return HeightOf(point) <= 0 || !c.InRange(0, ch.width)(PositionOf(point)) || ch.stopped_rocks[point]
}

func (ch *chamber) ShapeCanBeMovedTo(shape Shape) bool {
	return !c.Any(ch.is_point_occuppied, shape)
}

func (ch *chamber) AddShape(
	shape Shape,
	jet_index int,
	shape_index int,
	rock_index int64,
) func(rock_index, goal int64) int64 {
	c.ForEach(func(point Point) {
		ch.stopped_rocks[point] = true
		ch.height = c.Max(ch.height, HeightOf(point))
	}, shape)

	return ch.attempt_to_build_optimisation_function(jet_index, shape_index, rock_index)
}

func encode_closure(closure [][]bool) string {
	return strings.Join(c.Map(
		func(row []bool) string {
			return strings.Join(
				c.Map(
					func(b bool) string {
						if b {
							return "1"
						} else {
							return "0"
						}
					},
					row,
				), "")
		},
		closure,
	), "")
}

func (ch *chamber) build_optimisation_function(previous_closure, new_closure closure_info) func(int64, int64) int64 {
	height_increment := new_closure.height - previous_closure.height
	rock_index_increment := new_closure.rock_index - previous_closure.rock_index

	return func(rock_index, goal int64) int64 {
		increment_count := (goal - 1 - rock_index) / rock_index_increment
		ch.optimised_height_skip += int64(height_increment) * int64(increment_count)
		return rock_index + rock_index_increment*increment_count
	}
}

func (ch *chamber) takeTop(n int) [][]bool {
	if ch.height < n {
		return nil
	}

	return c.Map(
		func(h int) []bool {
			return c.Map(
				func(p int) bool { return ch.stopped_rocks[MakePoint(h, p)] },
				c.Range(0, ch.width),
			)
		},
		c.RangeInclusive(ch.height+1-n, ch.height),
	)
}

func (ch *chamber) is_block_a_closure(block [][]bool) bool {
	if len(block) == 0 {
		return false
	}

	for p := 1; p < ch.width; p++ {
		for h := 0; h < len(block); h++ {
			b := block[h][p-1]
			if h-1 >= 0 {
				b = b || block[h-1][p-1]
			}
			if h+1 < len(block) {
				b = b || block[h+1][p-1]
			}
			block[h][p] = block[h][p] && b
		}
	}
	return c.Any(
		func(h int) bool { return block[h][ch.width-1] },
		c.Range(0, len(block)),
	)
}

func (ch *chamber) attempt_to_build_optimisation_function(
	jet_index int,
	shape_index int,
	rock_index int64,
) func(rock_index, goal int64) int64 {

	for i := 1; i <= ch.width; i++ {
		block := ch.takeTop(i)
		if ch.is_block_a_closure(block) {
			C := closure{
				jetpattern_index: jet_index,
				shape_index:      shape_index,
				encoded_closure:  encode_closure(ch.takeTop(i)),
			}
			CI := closure_info{
				height:     ch.height,
				rock_index: rock_index,
			}
			if _, already_registered := ch.registered_closures[C]; already_registered {
				return ch.build_optimisation_function(ch.registered_closures[C], CI)
			}
			ch.registered_closures[C] = CI
			return nil
		}
	}
	return nil
}

func (ch *chamber) GetHeight() int {
	return ch.height
}

func (ch *chamber) GetTotalHeight() int64 {
	return ch.optimised_height_skip + int64(ch.height)
}

func (ch *chamber) String() string {
	chstr := []rune{}
	for h := ch.GetHeight() + 2; h > 0; h-- {
		chstr = append(chstr, '|')
		for p := 0; p < ch.width; p++ {
			if ch.stopped_rocks[MakePoint(h, p)] {
				chstr = append(chstr, '#')
			} else {
				chstr = append(chstr, '.')
			}
		}
		chstr = append(chstr, '|', '\n')
	}
	chstr = append(chstr, '+')
	chstr = append(chstr, c.Repeat('-', ch.width)...)
	chstr = append(chstr, '+', '\n')
	return string(chstr)
}
