package models

import (
	c "aoc/common"
	"fmt"
	"strings"
)

func ImageReport(image [][]rune) AnalyserReport {
	return image_report{
		c.Map(func(image_row []rune) string { return string(image_row) }, image),
	}
}

type image_report struct {
	image []string
}

func (report image_report) String() string {
	return fmt.Sprintf(
		"\n\n%s",
		strings.Join(report.image, "\n"),
	)
}

func (this_report image_report) equals(other AnalyserReport) bool {
	other_report, ok := other.(image_report)
	return ok && c.ArrayEqual(this_report.image, other_report.image)
}
