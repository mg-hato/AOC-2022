package main

import (
	c "aoc/common"
	m "aoc/d22/models"
	"aoc/d22/reader"
	"aoc/d22/solver"
	"aoc/reading"
	ts "aoc/testers"
	"testing"
)

func TestD22_IntegrationTest(t *testing.T) {
	ts.IntegrationTesterForComparableResults[m.SolverInput, int](t).
		ProvideReader(reading.ReadWith(reader.MonkeyMapReader)).
		ProvideSolver(solver.GetFinalCoordinates(m.CreateSimpleFieldMap)).
		ProvideSolver(solver.GetFinalCoordinates(m.CubeMapCreator)).
		AddTestCase("./tests/example.txt", ts.ExpectResult(6_032), ts.ExpectResult(5_031)).
		RunIntegrationTests()
}

func TestD22_StepByStepTest(t *testing.T) {

	// Function to process all instructions and return an ordered sequence of pointers
	type FieldMapCreator = func([]m.Field) (m.FieldMap, error)
	solver_f := func(field_map_creator_func FieldMapCreator) func(m.SolverInput) ([]m.Pointer, error) {
		return func(si m.SolverInput) ([]m.Pointer, error) {
			fields, instrs := si.Get().Get()
			fieldmap, err := field_map_creator_func(fields)
			if err != nil {
				return nil, err
			}
			ptr := fieldmap.GetInitialPointer()
			ptr_arr := []m.Pointer{ptr}

			for _, instr := range instrs {
				ptr = fieldmap.UpdatePointer(ptr, instr)
				ptr_arr = append(ptr_arr, ptr)
			}
			return ptr_arr, nil
		}
	}

	ts.IntegrationTester[m.SolverInput, []m.Pointer](t).
		ProvideEqualityFunctionForResults(c.ArrayEqual[m.Pointer]).
		ProvideReader(reading.ReadWith(reader.MonkeyMapReader)).
		ProvideSolver(solver_f(m.CreateSimpleFieldMap)).
		ProvideSolver(solver_f(m.CubeMapCreator)).
		AddTestCase(
			"./tests/example.txt",
			ts.ExpectResult([]m.Pointer{
				m.MakePointer(1, 9, m.East),
				m.MakePointer(1, 11, m.East),
				m.MakePointer(1, 11, m.South),
				m.MakePointer(6, 11, m.South),
				m.MakePointer(6, 11, m.East),
				m.MakePointer(6, 4, m.East),
				m.MakePointer(6, 4, m.South),
				m.MakePointer(8, 4, m.South),
				m.MakePointer(8, 4, m.East),
				m.MakePointer(8, 8, m.East),
				m.MakePointer(8, 8, m.South),
				m.MakePointer(6, 8, m.South),
				m.MakePointer(6, 8, m.East),
				m.MakePointer(6, 8, m.East),
			}),
			ts.ExpectResult([]m.Pointer{
				m.MakePointer(1, 9, m.East),
				m.MakePointer(1, 11, m.East),
				m.MakePointer(1, 11, m.South),
				m.MakePointer(6, 11, m.South),
				m.MakePointer(6, 11, m.East),
				m.MakePointer(11, 15, m.South),
				m.MakePointer(11, 15, m.West),
				m.MakePointer(11, 11, m.West),
				m.MakePointer(11, 11, m.South),
				m.MakePointer(6, 2, m.North),
				m.MakePointer(6, 2, m.East),
				m.MakePointer(6, 7, m.East),
				m.MakePointer(6, 7, m.North),
				m.MakePointer(5, 7, m.North),
			}),
		).RunIntegrationTests()
}
