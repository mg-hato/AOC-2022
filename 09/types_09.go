package main

import (
	. "aoc/functional"
	"fmt"
)

// Direction of the motion: Left, Right, Up or Down
type Direction int

const (
	Right Direction = iota
	Left
	Up
	Down
)

func (d Direction) String() string {
	switch d {
	case Right:
		return "R"
	case Left:
		return "L"
	case Up:
		return "U"
	default:
		return "D"
	}
}

// Represents one motion of going right, left, up or down and how many steps
type Motion struct {
	direction Direction
	steps     int
}

func (m Motion) String() string {
	return fmt.Sprintf("%s(%d)", m.direction, m.steps)
}

type Vector Pair[int, int]

func motion(direction string, steps int) Motion {
	var d Direction
	switch direction {
	case "L":
		d = Left
	case "R":
		d = Right
	case "U":
		d = Up
	default:
		d = Down
	}
	return Motion{d, steps}
}

type HeadTail struct {
	head  Vector   // position of the head
	tails []Vector // position of the tails

	visited map[Vector]bool // Visited positions by the last tail
}

// Create Head-Tail data-type with n-tails
func CreateHeadTail(n int) HeadTail {
	return HeadTail{
		tails:   make([]Vector, n),
		visited: map[Vector]bool{{}: true},
	}
}

func (ht *HeadTail) update(d Direction) {
	ht.head = add(ht.head, d.getUpdateVector())
	target := ht.head
	for i := 0; i < len(ht.tails); i++ {
		ht.tails[i] = follow(ht.tails[i], target)
		target = ht.tails[i]
	}
	ht.visited[ht.tails[len(ht.tails)-1]] = true
}

// Returns the count of tiles visited by the last tail
func (ht *HeadTail) getVisitedCount() int {
	return len(ht.visited)
}

// Get a new position of the follower that follows target
func follow(follower, target Vector) Vector {

	// If the follower is adjacent to the target: do not move
	if areAdjacent(follower, target) {
		return follower
	}

	// If they are on the "same line" (same on x- or y-axis)
	// follower gets closer (in the middle between target and follower)
	if follower.First == target.First || follower.Second == target.Second {
		return Vector{
			First:  (follower.First + target.First) / 2,
			Second: (follower.Second + target.Second) / 2,
		}
	}

	// Otherwise, follower moves diagonally closer to the target

	// Diagonal directions
	diagonal_directions := []Vector{
		{First: 1, Second: 1},
		{First: -1, Second: -1},
		{First: 1, Second: -1},
		{First: -1, Second: 1},
	}

	// One of these positions must be the new position the follower takes
	candidate_updates := Map(func(v Vector) Vector { return add(follower, v) }, diagonal_directions)

	// Return new position
	return Filter(func(v Vector) bool {
		return areAdjacent(target, v)
	}, candidate_updates)[0]
}

// Create update-position vector
func (d Direction) getUpdateVector() Vector {
	switch d {
	case Left:
		return Vector{First: -1, Second: 0}
	case Right:
		return Vector{First: 1, Second: 0}
	case Up:
		return Vector{First: 0, Second: 1}
	default: // Down
		return Vector{First: 0, Second: -1}
	}
}

// Add vectors
func add(u, v Vector) Vector {
	return Vector{First: u.First + v.First, Second: u.Second + v.Second}
}

// Multiply scalar with the vector
func multiply(scalar int, v Vector) Vector {
	return Vector{
		First:  scalar * v.First,
		Second: scalar * v.Second,
	}
}

// Absolute value of an integer
func abs(i int) int {
	if i < 0 {
		return -i
	} else {
		return i
	}
}

// Length of vector v
func length(v Vector) int {
	return abs(v.First) + abs(v.Second)
}

// Negate pair
func negate(v Vector) Vector {
	return multiply(-1, v)
}

// Distance between two vectors
func distance(u, v Vector) int {
	return length(add(u, negate(v)))
}

// Returns true iff vectors u and v are adjacent:
// Two vectors are adjacent iff their difference-vector has both coordinates in range [-1,1]
func areAdjacent(u, v Vector) bool {
	diff := add(u, negate(v))
	return abs(diff.First) <= 1 && abs(diff.Second) <= 1
}
