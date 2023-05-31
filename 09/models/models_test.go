package models

import (
	f "aoc/functional"
	ts "aoc/testers"
	"testing"
)

func TestD09_Envelope(t *testing.T) {
	data := func() MotionSeries {
		return MotionSeries{
			MakeMotion(10, RIGHT),
			MakeMotion(21, UP),
			MakeMotion(7, LEFT),
			MakeMotion(8, DOWN),
			MakeMotion(3, LEFT),
			MakeMotion(13, DOWN),
		}
	}

	envelope := MotionSeriesEnvelope(data())

	series := envelope.Get()
	series[0] = MakeMotion(7, LEFT)

	ts.AssertEqualWithEqFunc(t, envelope.Get(), data(), f.ArrayEqual[Motion])
}

func TestD09_Distance(t *testing.T) {
	p := f.MakePair[int, int]
	ts.AssertEqual(t, Distance(p(10, 20), p(20, 10)), 20)
	ts.AssertEqual(t, Distance(p(-100, 256), p(100, 255)), 201)
	ts.AssertEqual(t, Distance(p(0, 0), p(10, 0)), 10)
	ts.AssertEqual(t, Distance(p(151, 999), p(151, 999)), 0)
	ts.AssertEqual(t, Distance(p(4, 3), p(5, 4)), 2)
}

func TestD09_FollowLeader(t *testing.T) {
	p := f.MakePair[int, int]

	// No movement expected for the following
	ts.AssertEqual(t, FollowLeader(p(15, 10), p(15, 10)), p(15, 10))
	ts.AssertEqual(t, FollowLeader(p(101, 101), p(100, 100)), p(100, 100))
	ts.AssertEqual(t, FollowLeader(p(4, 4), p(3, 4)), p(3, 4))
	ts.AssertEqual(t, FollowLeader(p(0, 0), p(1, 0)), p(1, 0))

	// Axis parallel movement expected
	ts.AssertEqual(t, FollowLeader(p(10, 10), p(8, 10)), p(9, 10))
	ts.AssertEqual(t, FollowLeader(p(0, 0), p(0, 2)), p(0, 1))
	ts.AssertEqual(t, FollowLeader(p(-7, -7), p(-7, -9)), p(-7, -8))

	// Diagonal movement expected
	ts.AssertEqual(t, FollowLeader(p(0, 0), p(2, 2)), p(1, 1))
	ts.AssertEqual(t, FollowLeader(p(13, 10), p(11, 9)), p(12, 10))
	ts.AssertEqual(t, FollowLeader(p(-16, -79), p(-15, -81)), p(-16, -80))
}
