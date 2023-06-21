package solver

import (
	m "aoc/d10/models"
	ts "aoc/testers"
	"testing"
)

func TestD10_SignalStrength(t *testing.T) {
	analyser := SignalStrengthAnalyser(10, 5, 100)
	analyser.Initialise()
	analyser.Capture(20)
	analyser.Capture(100)

	ts.AssertEqual(t, analyser.IsDone(), false)

	analyser.Capture(75)

	ts.AssertEqual(t, analyser.IsDone(), true)

	expected_report := m.SignalStrengthReport(5*20 + 10*100 + 100*75)
	ts.AssertEqualWithEqFunc(t, analyser.GenerateReport(), expected_report, m.AnalyserReportEqualityFunction)
}

func TestD10_ImageDrawer(t *testing.T) {
	analyser := ImageDrawerAnalyser(5, 2)
	analyser.Initialise()

	ts.AssertEqual(t, analyser.IsDone(), false)
	for i := 1; i <= 5; i++ {
		if i%2 == 1 {
			analyser.Capture(i)
		} else {
			analyser.Capture(-10)
		}
	}
	ts.AssertEqual(t, analyser.IsDone(), false)

	for i := 1; i <= 5; i++ {
		analyser.Capture(2)
	}

	ts.AssertEqual(t, analyser.IsDone(), true)

	expected_report := m.ImageReport([][]rune{
		[]rune("#.#.#"),
		[]rune(".###."),
	})
	ts.AssertEqualWithEqFunc(t, analyser.GenerateReport(), expected_report, m.AnalyserReportEqualityFunction)
}
