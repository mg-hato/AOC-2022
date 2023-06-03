package models

type AnalyserReport interface {
	String() string
	equals(AnalyserReport) bool
}

func AnalyserReportEqualityFunction(lhs, rhs AnalyserReport) bool {
	return lhs.equals(rhs)
}
