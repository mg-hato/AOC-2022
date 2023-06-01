package models

import "fmt"

func SignalStrengthReport(signal_strength int64) AnalyserReport {
	return signal_strength_report{signal_strength}
}

type signal_strength_report struct {
	signal_strength int64
}

func (report signal_strength_report) String() string {
	return fmt.Sprint(report.signal_strength)
}

func (this_report signal_strength_report) equals(other AnalyserReport) bool {
	other_report, ok := other.(signal_strength_report)
	return ok && this_report.signal_strength == other_report.signal_strength
}
