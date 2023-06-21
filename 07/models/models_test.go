package models

import (
	c "aoc/common"
	ts "aoc/testers"
	"testing"
)

func TestD07_EnvelopeTest(t *testing.T) {

	data_generator := func() []Command {
		return []Command{
			MakeCommandCd("/"),
			MakeCommandLs(
				MakeFile("a.txt", 1_001),
				MakePartialDirectory("b"),
				MakeFile("x.log", 100_100),
			),
			MakeCommandCd("b"),
			MakeCommandLs(),
		}
	}

	envelope := CreateCommandsEnvelope(data_generator()...)

	commands := envelope.Get()
	commands[3] = MakeCommandCd("x")

	ts.AssertEqualWithEqFunc(t, envelope.Get(), data_generator(), c.ArrayEqualWith(CommandEqualityFunc))
}

func TestD07_ShallowCopyAndEquality(t *testing.T) {
	root := &Directory{name: "home"}
	items := []Item{
		&File{name: "x.txt", size: 100, parent: root},
		&Directory{name: "test", parent: root, items: map[string]Item{
			"ee":   MakeFile("ee", 150),
			"tyy":  MakeFile("tyy", 300),
			"docs": MakePartialDirectory("docs"),
		}},
	}

	new_items := c.Map(Item.shallow_copy, items)

	file, ok := new_items[0].(*File)
	ts.Assert(t, ok)
	ts.AssertEqual(t, file.name, "x.txt")
	ts.AssertEqual(t, file.size, 100)
	ts.Assert(t, file.parent == nil)

	dir, ok := new_items[1].(*Directory)
	ts.Assert(t, ok)
	ts.AssertEqual(t, dir.name, "test")
	ts.Assert(t, dir.parent == nil)
	ts.Assert(t, dir.items == nil)

	ts.AssertEqualWithEqFunc(t, items, new_items, c.ArrayEqualWith(ItemShallowEqualityFunction))
}

func TestD07_CreateFilesystem(t *testing.T) {
	type Test struct {
		input          []Command
		expected_error error
	}

	ts.TestThat([]Test{
		{
			[]Command{
				MakeCommandCd("/"),
				MakeCommandLs(
					MakeFile("a", 100),
					MakeFile("b", 101),
				),
				MakeCommandCd("a"),
			},
			create_filesystem_error(3, MakeCommandCd("a"), not_a_directory_error("/", "a")),
		},
		{
			[]Command{
				MakeCommandCd("/"),
				MakeCommandCd("x"),
			},
			create_filesystem_error(2, MakeCommandCd("x"), directory_content_is_unknown_error("/", "x")),
		},
		{
			[]Command{
				MakeCommandCd("/"),
				MakeCommandLs(
					MakePartialDirectory("docs"),
					MakePartialDirectory("pics"),
				),
				MakeCommandCd("music"),
			},
			create_filesystem_error(3, MakeCommandCd("music"), directory_does_not_exist_error("/", "music")),
		},
		{
			[]Command{
				MakeCommandCd("/"),
				MakeCommandLs(
					MakePartialDirectory("x"),
					MakeFile("a.log", 101),
				),
				MakeCommandCd("x"),
				MakeCommandLs(
					MakeFile("info.txt", 100),
					MakeFile("alarm.avi", 5_550),
				),
				MakeCommandCd(".."),
				MakeCommandLs(
					MakePartialDirectory("x"),
					MakeFile("a.log", 105),
				),
			},
			create_filesystem_error(
				6,
				MakeCommandLs(
					MakePartialDirectory("x"),
					MakeFile("a.log", 105),
				),
				ls_items_do_not_match_error("/"),
			),
		},
		{
			[]Command{
				MakeCommandCd("/"),
				MakeCommandLs(
					MakePartialDirectory("a"),
					MakePartialDirectory("b"),
				),
				MakeCommandCd("a"),
				MakeCommandLs(),
				MakeCommandCd("/"),
				MakeCommandLs(
					MakePartialDirectory("a"),
					MakePartialDirectory("b"),
				),
			},
			create_filesystem_verification_error(directory_is_unexplored_error("b")),
		},
	}, func(test Test) {
		_, e := CreateFilesystem(test.input)
		ts.AssertEqual(t, e.Error(), test.expected_error.Error())
	})
}

func TestD07_CalculateItemSizes(t *testing.T) {
	root := &Directory{name: "/"}
	root.parent = root

	// a file in root directory
	checks := &File{name: "checks.zip", size: 99_801, parent: root}

	// documents directory
	docs := &Directory{name: "docs", parent: root}
	contract := &File{name: "contract.pdf", size: 2_560, parent: docs}
	cv := &File{name: "cv.pdf", size: 1_500, parent: docs}
	docs.items = map[string]Item{
		"contract.pdf": contract,
		"cv.pdf":       cv,
	}

	// pics directory
	pics := &Directory{name: "pics", parent: root}
	dog := &File{name: "dog.jpg", size: 15_121, parent: pics}
	cat := &File{name: "cat.jpg", size: 13_256, parent: pics}
	pics.items = map[string]Item{
		"cat.jpg": cat,
		"dog.jpg": dog,
	}

	root.items = map[string]Item{
		"pics":       pics,
		"checks.zip": checks,
		"docs":       docs,
	}

	sizes := CalculateItemSizes(root)
	expected_sizes := map[Item]int64{
		checks: int64(checks.size),

		cat:  int64(cat.size),
		dog:  int64(dog.size),
		pics: int64(cat.size + dog.size),

		cv:       int64(cv.size),
		contract: int64(contract.size),
		docs:     int64(contract.size + cv.size),

		root: int64(cv.size + contract.size + dog.size + cat.size + checks.size),
	}

	ts.AssertEqual(t, len(sizes), len(expected_sizes))
	c.ForEach(func(item Item) {
		ts.AssertEqual(t, sizes[item], expected_sizes[item])
	}, []Item{root, checks, pics, cat, dog, docs, cv, contract})
}
