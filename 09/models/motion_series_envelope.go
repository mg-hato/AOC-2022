package models

import (
	e "aoc/envelope"
	f "aoc/functional"
	"fmt"
)

type motion_series_envelope struct {
	motion_series MotionSeries
}

func MotionSeriesEnvelope(motion_series MotionSeries) e.Envelope[MotionSeries] {
	return motion_series_envelope{motion_series}
}

func (envelope motion_series_envelope) Get() MotionSeries {
	return f.Map(f.Identity[Motion], envelope.motion_series)
}

func (envelope motion_series_envelope) String() string {
	return fmt.Sprintf("MotionSeriesEnvelope%s", envelope.motion_series)
}

func MotionSeriesEnvelopeEqualityFunc(lhs, rhs e.Envelope[MotionSeries]) bool {
	return f.ArrayEqual(lhs.Get(), rhs.Get())
}
