package solver

import (
	c "aoc/common"
	m "aoc/d09/models"
	ts "aoc/testers"
	"testing"
)

func TestD09_MoveKnots(t *testing.T) {

	// shortcut for directions
	up, down, left, right := m.UP, m.DOWN, m.LEFT, m.RIGHT

	// shortcut to make a movement-vector corresponding to the given motion model
	make_movement := func(motion m.Motion) m.Movement {
		return c.MakePair(
			motion.Direction.AsMovement().First*motion.Steps,
			motion.Direction.AsMovement().Second*motion.Steps,
		)
	}

	// shortcut to make a position produced from following a sequence of motions starting from position (0,0)
	position := func(motions ...m.Motion) m.Position {
		p := c.MakePair(0, 0)
		c.ForEach(func(motion m.Motion) { p = m.Move(p, make_movement(motion)) }, motions)
		return p
	}

	knots := c.Repeat(c.MakePair(0, 0), 3)

	// move head 4 times up
	// we expect other knots to form a line one behind the other
	move_knots(up, knots)
	move_knots(up, knots)
	move_knots(up, knots)
	move_knots(up, knots)
	ts.AssertEqualWithEqFunc(t, knots, []m.Position{
		position(m.MakeMotion(4, up)),
		position(m.MakeMotion(3, up)),
		position(m.MakeMotion(2, up)),
	}, c.ArrayEqual[m.Position])

	// move head one time down
	// we expect that tailing knots have not moved
	move_knots(down, knots)
	ts.AssertEqualWithEqFunc(t, knots, []m.Position{
		position(m.MakeMotion(3, up)),
		position(m.MakeMotion(3, up)),
		position(m.MakeMotion(2, up)),
	}, c.ArrayEqual[m.Position])

	// move head one time left
	// again, no movement is expected of tailing knots
	move_knots(left, knots)
	ts.AssertEqualWithEqFunc(t, knots, []m.Position{
		position(m.MakeMotion(3, up), m.MakeMotion(1, left)),
		position(m.MakeMotion(3, up)),
		position(m.MakeMotion(2, up)),
	}, c.ArrayEqual[m.Position])

	// move head once more to the left
	// expect 2nd knot to follow with one movement to the left
	// last knot should still stay put
	move_knots(left, knots)
	ts.AssertEqualWithEqFunc(t, knots, []m.Position{
		position(m.MakeMotion(3, up), m.MakeMotion(2, left)),
		position(m.MakeMotion(3, up), m.MakeMotion(1, left)),
		position(m.MakeMotion(2, up)),
	}, c.ArrayEqual[m.Position])

	// move head once more to the left
	// 2nd knot should follow in-line
	// 3rd knot should make a diagonal movement up-left
	move_knots(left, knots)
	ts.AssertEqualWithEqFunc(t, knots, []m.Position{
		position(m.MakeMotion(3, up), m.MakeMotion(3, left)),
		position(m.MakeMotion(3, up), m.MakeMotion(2, left)),
		position(m.MakeMotion(3, up), m.MakeMotion(1, left)),
	}, c.ArrayEqual[m.Position])

	// move head two times to the right and once upwards
	// we expect knots 2 and 3 to stay put during these movements
	move_knots(right, knots)
	move_knots(right, knots)
	move_knots(up, knots)
	ts.AssertEqualWithEqFunc(t, knots, []m.Position{
		position(m.MakeMotion(4, up), m.MakeMotion(1, left)),
		position(m.MakeMotion(3, up), m.MakeMotion(2, left)),
		position(m.MakeMotion(3, up), m.MakeMotion(1, left)),
	}, c.ArrayEqual[m.Position])

	// lastly, we move head one more time to the right
	// we expect 2nd knot to move diagonally up-right
	// 3rd knot should stay put
	move_knots(right, knots)
	ts.AssertEqualWithEqFunc(t, knots, []m.Position{
		position(m.MakeMotion(4, up)),
		position(m.MakeMotion(4, up), m.MakeMotion(1, left)),
		position(m.MakeMotion(3, up), m.MakeMotion(1, left)),
	}, c.ArrayEqual[m.Position])
}
