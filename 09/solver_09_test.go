package main

import (
	"aoc/testers"
	"testing"
)

func TestSolver(t *testing.T) {
	tester := testers.DefaultSolverTesterForComparableTypeR(
		SimulateRopeWithTailCount(1),
		SimulateRopeWithTailCount(9),
		"CountVisitedOfTail",
		"CountVisitedOf9thTail",
	)

	// This is the problem statement example
	tester.AddTest([]Motion{
		motion("R", 4),
		motion("U", 4),
		motion("L", 3),
		motion("D", 1),
		motion("R", 4),
		motion("D", 1),
		motion("L", 5),
		motion("R", 2),
	}, 13, 1)

	// Custome input #1
	// Go left 13 times
	// On start: tail is at (0,0)
	// On 1 move left: tail stays
	// 12 more moves left: new tail position
	// i.e. 13
	tester.AddTest([]Motion{
		motion("L", 13),
	}, 13, 5)

	// Custom input #2
	// Same as #1 but we circle around "L 13"-position
	// i.e.
	// go up (still diagonal)
	// go right 2 (still directly/diagonal)
	// go down 2 (still directly/diagonal)
	// go left 2 (still directly/diagonal)
	// go up 2 (still directly/diagonal)
	tester.AddTest([]Motion{
		motion("L", 13),
		motion("U", 1),
		motion("R", 2),
		motion("D", 2),
		motion("L", 2),
		motion("U", 2),
	}, 13, 5)

	// Another sample input given for task 2
	// For task 1, the visited tiles are given below
	tester.AddTest([]Motion{
		motion("R", 5),
		motion("U", 8),
		motion("L", 8),
		motion("D", 3),
		motion("R", 17),
		motion("D", 10),
		motion("L", 25),
		motion("U", 20),
	}, 88, 36)

	tester.RunBothSolversTests(t)
}

/*
- H is the position of head at the end
- # is a tile visited by the 1st tail (only tail for task 1)

H.........................
#.........................
#.........................
#.........................
#.........................
#.........................
#.........................
#........#######..........
#.......#.......#.........
#.......#.......#.........
#........################.
#...............#........#
#...............#........#
#...............#........#
#...............#........#
#..........#####.........#
#........................#
#........................#
#........................#
#........................#
.########################.
*/
