package models

import (
	e "aoc/envelope"
	f "aoc/functional"
)

type sensor_reports_envelope struct {
	reports []SensorReport
}

func SensorReportsEnvelope(reports ...SensorReport) e.Envelope[[]SensorReport] {
	return sensor_reports_envelope{reports}
}

func (env sensor_reports_envelope) Get() []SensorReport {
	return f.Map(f.Identity[SensorReport], env.reports)
}

func SensorReportsEnvelopeEqualityFunction(lhs, rhs e.Envelope[[]SensorReport]) bool {
	return f.ArrayEqual[SensorReport](lhs.Get(), rhs.Get())
}
