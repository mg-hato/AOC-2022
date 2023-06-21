module aoc/d06

go 1.19

replace aoc/argshandle => ../argshandle

replace aoc/common => ../common

replace aoc/reading => ../reading

replace aoc/testers => ../testers

require aoc/argshandle v0.0.0-00010101000000-000000000000

require (
	aoc/common v0.0.0-00010101000000-000000000000
	aoc/reading v0.0.0-00010101000000-000000000000
	aoc/testers v0.0.0-00010101000000-000000000000
)

require golang.org/x/exp v0.0.0-20230522175609-2e198f4a06a1 // indirect
