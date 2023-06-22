module aoc/d16

replace aoc/reading => ../reading

replace aoc/common => ../common

replace aoc/argshandle => ../argshandle

replace aoc/testers => ../testers

go 1.19

require (
	aoc/argshandle v0.0.0-00010101000000-000000000000
	aoc/common v0.0.0-00010101000000-000000000000
	aoc/reading v0.0.0-00010101000000-000000000000
	aoc/testers v0.0.0-00010101000000-000000000000
)

require golang.org/x/exp v0.0.0-20230522175609-2e198f4a06a1 // indirect
