package solver

import (
	c "aoc/common"
	"aoc/d10/models"
)

func ImageDrawerAnalyser(width, height int) RegisterCapturingAnalyser {
	return &image_drawer_analyser{
		width:  width,
		height: height,
	}
}

type image_drawer_analyser struct {
	width, height int
	current       int
	image         [][]rune
}

func (ida *image_drawer_analyser) Initialise() {
	ida.current = 1
	ida.image = c.Map(
		func([]rune) []rune { return c.Repeat('.', ida.width) },
		make([][]rune, ida.height),
	)
}

func (ida image_drawer_analyser) IsDone() bool {
	return ida.current > ida.height*ida.width
}

func (ida *image_drawer_analyser) Capture(register int) {
	current_position := (ida.current - 1) % ida.width
	if c.InInclusiveRange(register-1, register+1)(current_position) {
		current_height := (ida.current - 1) / ida.width
		ida.image[current_height][current_position] = '#'
	}
	ida.current++
}

func (ida image_drawer_analyser) NextCycle() int {
	return ida.current
}

func (ida image_drawer_analyser) GenerateReport() models.AnalyserReport {
	if ida.IsDone() {
		return models.ImageReport(ida.image)
	}
	return nil
}
