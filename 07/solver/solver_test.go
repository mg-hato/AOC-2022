package solver

import (
	m "aoc/d07/models"
	ts "aoc/testers"
	"testing"
)

func TestD07_SolverTest(t *testing.T) {
	spec := SimpleFilesystemSpec(200, 100)
	ts.SolverTesterForComparableResults[m.SolverInput, int64](t).
		ProvideSolver(AnalyseFilesystem(spec, SumDirectoriesOfSizeAtMost(50))).
		ProvideSolver(AnalyseFilesystem(spec, FindSmallestDirectoryEnablingUpdate())).
		// Directories of size at most 50:
		// - a with size 1
		// - b is too big
		// - c with size exactly 50
		//   - env within c of size 37
		// In total: 88
		// ----------
		// Total memory: 200
		// Required free memory for update: 100
		// Used memory: 170
		// - so we need a file that frees at least 70 units of memory
		// - two candidates: root ("/") and directory "b"
		// - former would free up all used memory: 170, latter would free up 119 units of memory
		AddTestCase(m.CreateCommandsEnvelope(
			m.MakeCommandCd("/"),
			m.MakeCommandLs(
				m.MakePartialDirectory("a"),
				m.MakePartialDirectory("b"),
				m.MakePartialDirectory("c"),
			),

			m.MakeCommandCd("a"),
			m.MakeCommandLs(
				m.MakeFile("a", 1),
			),

			m.MakeCommandCd("/"),

			m.MakeCommandCd("b"),
			m.MakeCommandLs(
				m.MakeFile("app.log", 99),
				m.MakeFile("load.sh", 20),
			),

			m.MakeCommandCd("/"),

			m.MakeCommandCd("c"),
			m.MakeCommandLs(
				m.MakeFile("e.gzip.zip", 13),
				m.MakePartialDirectory("env"),
			),
			m.MakeCommandCd("env"),
			m.MakeCommandLs(
				m.MakeFile("application.properties", 37),
			),
		), ts.ExpectResult[int64](88), ts.ExpectResult[int64](119)).
		RunSolverTests()
}
