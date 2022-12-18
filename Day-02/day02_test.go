package main

import (
	"aoc/functional"
	"fmt"
	"testing"
)

func TestDay02_TestReader(t *testing.T) {
	expected := map[int]EncryptedStrategyGuide{
		1: {[]Round{
			{A, Y},
			{B, X},
			{C, Z},
		}},
		2: {[]Round{{C, X}}},
		3: {[]Round{
			{A, Z},
			{A, X},
			{B, X},
			{C, Y},
			{C, X},
			{C, X},
		}},
	}

	for k, exp := range expected {
		filename := fmt.Sprintf("./test/input.%d", k)
		if readStrategy, _ := ReadStrategy(filename); !readStrategy.equals(exp) {
			t.Errorf("Test #%d failed: ReadStrategy(\"%s\")", k, filename)
			t.Errorf("Returned: %v", *readStrategy)
			t.Errorf("Expected: %v", exp)
		}
	}
}

func (lhs EncryptedStrategyGuide) equals(rhs EncryptedStrategyGuide) bool {
	return functional.ArrayEqual(lhs.rounds, rhs.rounds)
}

func TestDay02_Solver(t *testing.T) {

	// A reminder of scoring: Shapes[R -> 1, P -> 2, S -> 3] & Outcome[W -> 6, Draw -> 3, L -> 0]

	tests := []struct {
		round      Round
		xyzDecoder XYZDecoderStrategy
		expected   int
	}{
		{Round{A, X}, DirectlyAsShape, 1 + 3}, // RR: Shape(1), Draw(3)
		{Round{A, Z}, DirectlyAsShape, 3 + 0}, // RS: Shape(3), Lose(0)
		{Round{C, X}, DirectlyAsShape, 1 + 6}, // SR: Shape(1), Win(6)
		{Round{B, Y}, DirectlyAsShape, 2 + 3}, // PP: Shape(2), Draw(3)

		{Round{A, X}, AsDesiredOutcome, 3 + 0}, // RS: Shape(3), Lose(0)
		{Round{C, Z}, AsDesiredOutcome, 1 + 6}, // SR: Shape(1), Win(6)
		{Round{B, Y}, AsDesiredOutcome, 2 + 3}, // PP: Shape(2), Draw(3)
		{Round{B, Z}, AsDesiredOutcome, 3 + 6}, // PS: Shape(3), Win(6)
	}

	for test_number, test := range tests {
		esg := EncryptedStrategyGuide{[]Round{test.round}}
		if result := CalculateScore(&esg, test.xyzDecoder); result != test.expected {
			t.Errorf("Test #%d failed: CalculateScore(%s, %s)", test_number+1, esg, test.xyzDecoder)
			t.Errorf("Returned: %d", result)
			t.Errorf("Expected: %d", test.expected)
		}
	}
}
