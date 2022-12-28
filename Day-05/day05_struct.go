package main

import (
	. "aoc/functional"
	"fmt"
	"strings"
)

// Structure representing move/rearrangement instruction
type Move struct {
	qty         int // how many crates to move
	source      int // stack id from which crates are moved
	destination int // stack id to which crates are moved
}

type Stack struct {
	stack_id int    // id of the stack
	boxes    []rune // stacked boxes
}

func (stack Stack) String() string {
	return fmt.Sprintf("[%d > %s]", stack.stack_id, string(stack.boxes))
}

func (move Move) String() string {
	return fmt.Sprintf("(%d: %d => %d)", move.qty, move.source, move.destination)
}

// Put the box on top of the stack
func (s *Stack) push(box rune) {
	s.boxes = append(s.boxes, box)
}

// Returns true iff stack is empty
func (s Stack) empty() bool {
	return len(s.boxes) == 0
}

// Remove top box from the stack
func (s *Stack) pop() rune {
	size := len(s.boxes)
	top_box := s.boxes[size-1]
	s.boxes = s.boxes[:size-1]
	return top_box
}

// Look what box is on top of the stack
func (s Stack) peek() rune {
	return s.boxes[len(s.boxes)-1]
}

type RearrangementPlan struct {
	stacks []string
	moves  []Move
}

func (plan RearrangementPlan) String() string {
	formatted_stacks := Map(func(s string) string { return fmt.Sprintf("%s", s) }, plan.stacks)
	formatted_moves := Map(func(m Move) string { return fmt.Sprint(m) }, plan.moves)
	return strings.Join(append(formatted_stacks, formatted_moves...), "\n")
}

type CrateMover interface {
	ApplyMove(Move, map[int]*Stack)
}

type CrateMover9000 struct{}

func (_ CrateMover9000) ApplyMove(move Move, stacks map[int]*Stack) {
	for counter := 0; counter < move.qty; counter++ {
		if stacks[move.source].empty() {
			break
		}
		stacks[move.destination].push(stacks[move.source].pop())
	}
}

type CrateMover9001 struct{}

func (_ CrateMover9001) ApplyMove(move Move, stacks map[int]*Stack) {
	var boxes []rune
	for counter := 0; counter < move.qty; counter++ {
		if stacks[move.source].empty() {
			break
		}
		boxes = append(boxes, stacks[move.source].pop())
	}
	ForEach(func(box rune) { stacks[move.destination].push(box) }, Reverse(boxes))
}
