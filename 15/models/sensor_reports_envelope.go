package models

import (
	c "aoc/common"
)

type sensor_reports_envelope struct {
	reports []SensorReport
}

func SensorReportsEnvelope(reports ...SensorReport) c.Envelope[[]SensorReport] {
	return sensor_reports_envelope{reports}
}

func (env sensor_reports_envelope) Get() []SensorReport {
	return c.Map(c.Identity[SensorReport], env.reports)
}

func SensorReportsEnvelopeEqualityFunction(lhs, rhs c.Envelope[[]SensorReport]) bool {
	return c.ArrayEqual[SensorReport](lhs.Get(), rhs.Get())
}
