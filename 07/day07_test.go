package main

import (
	. "aoc/functional"
	"aoc/testers"
	"testing"
)

func f(size int, name string) ListedItem {
	return ListedItem{size, FileType, name}
}

func dir(name string) ListedItem {
	return ListedItem{0, DirectoryType, name}
}

func cd(arg string) Command {
	return Command{ChangeDirectory, arg, nil}
}

func ls(items ...ListedItem) Command {
	return Command{List, "", items}
}

func TestDay07_Reader(t *testing.T) {
	tester := testers.DefaultReaderTester(ReadTerminalOutput, "ReadTerminalOutput")
	tester.ProvideEqualityFunctionForTypeT(ArrayEqualWith(areCommandsEqual))

	tester.AddGoodInputTests([]Command{
		cd("/"),
		ls(dir("a"), f(14848514, "b.txt"), f(8504156, "c.dat"), dir("d")),
		cd("a"),
		ls(dir("e"), f(29116, "f"), f(2557, "g"), f(62596, "h.lst")),
		cd("e"),
		ls(f(584, "i")),
		cd(".."),
		cd(".."),
		cd("d"),
		ls(f(4060174, "j"), f(8033020, "d.log"), f(5626152, "d.ext"), f(7214296, "k")),
	})

	tester.AddGoodInputTests([]Command{
		cd("/"),
		ls(dir("x"), dir("y"), dir("z"), f(169, "funny.jpg"), f(4096, "cat.png")),
		cd("z"),
		ls(f(1050912, "prog.exe"), f(133, "prog.c")),
		cd(".."),
		cd("x"),
		ls(),
		cd(".."),
		cd("y"),
		ls(f(11, "lel")),
	})

	tester.AddErrorInputTests(
		"Line 14: cd command is not followed by listed items",
		"Line 16: command 'cls' is not supported",
		"Lines 3-4: listed items start with a number (size) first",
	)

	tester.RunBothGoodAndErrorInputTests(t)
}

// Returns true iff two provided commands are identical
func areCommandsEqual(lhs, rhs Command) bool {
	return lhs.command_type == rhs.command_type && lhs.argument == rhs.argument && ArrayEqual(lhs.listed_items, rhs.listed_items)
}

func TestDay07_Solver(t *testing.T) {
	tester := testers.DefaultSolverTesterForComparableTypeR(
		ForFilesystemGet(SumOfDirectoriesLte100k),
		ForFilesystemGet(SmallestUpdateEnablingDirectorySize),
		"find-sum-of-directories(<=100k)",
		"find-size-of-smallest-update-enabling-directory",
	)
	tester.AddTest([]Command{
		cd("/"),
		ls(dir("a"), f(14848514, "b.txt"), f(8504156, "c.dat"), dir("d")),
		cd("a"),
		ls(dir("e"), f(29116, "f"), f(2557, "g"), f(62596, "h.lst")),
		cd("e"),
		ls(f(584, "i")),
		cd(".."),
		cd(".."),
		cd("d"),
		ls(f(4060174, "j"), f(8033020, "d.log"), f(5626152, "d.ext"), f(7214296, "k")),
	}, 95_437, 24_933_642)

	/* Custom input #1
	/
	-- i
	    -- logs
	    	-- 1050: app.log
			-- 98_951: oldapp.log
	-- j
	    -- conf
	    	-- 201: application.properties
			-- 89: app.env
		-- 99_710: app.jar
	-- k
		-- 9: app.version
	*/
	// Two files <= 100K: k (9), conf (290) & j (100,000)
	tester.AddTest([]Command{
		cd("/"),
		ls(dir("i"), dir("j"), dir("k")),

		cd("k"),
		ls(f(9, "app.version")),

		cd(".."),
		cd("i"),
		ls(dir("logs")),
		cd("logs"),
		ls(f(1050, "app.log"), f(98_951, "oldapp.log")),
		cd("/"),

		cd("j"),
		ls(dir("conf"), f(99_710, "app.jar")),
		cd("conf"),
		ls(f(201, "application.properties"), f(89, "app.env")),
	}, 100_299, 0)

	tester.RunBothSolversTests(t)
}
