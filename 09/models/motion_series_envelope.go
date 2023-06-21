package models

import (
	c "aoc/common"
	"fmt"
)

type motion_series_envelope struct {
	motion_series MotionSeries
}

func MotionSeriesEnvelope(motion_series MotionSeries) c.Envelope[MotionSeries] {
	return motion_series_envelope{motion_series}
}

func (envelope motion_series_envelope) Get() MotionSeries {
	return c.Map(c.Identity[Motion], envelope.motion_series)
}

func (envelope motion_series_envelope) String() string {
	return fmt.Sprintf("MotionSeriesEnvelope%s", envelope.motion_series)
}

func MotionSeriesEnvelopeEqualityFunc(lhs, rhs c.Envelope[MotionSeries]) bool {
	return c.ArrayEqual(lhs.Get(), rhs.Get())
}
