package argshandle

import (
	ts "aoc/testers"
	"errors"
	"testing"
)

func TestArgsHandle_WhenReaderSuccessfullyReads_CheckThatSolverIsCalled(t *testing.T) {

	// Mocked solver
	var expected_outcome string = "Solver is called"
	solver := func(_ string) (string, error) {
		return expected_outcome, nil
	}

	// Buffered channel of size 1
	channel := make(chan string, 1)

	// Mocked command-line arguments
	arguments := []string{"program_name", "--i", "some-input"}

	// Mocked reader. Returns successfully.
	reader := func(_ string) (string, error) {
		return "Reader reads", nil
	}

	ts.Assert(t, true)

	solution_received, e := HandleArgumentsAndExecute(arguments, reader, solver, solver, channel)
	ts.AssertNoError(t, e)
	ts.AssertEqual(t, solution_received, true)
	if solution_received {
		ts.AssertEqual(t, <-channel, expected_outcome)
	}
}

func TestArgsHandle_WhenReaderUnsuccessful_ExpectNoSolution(t *testing.T) {

	// Mocked solver
	var solver_solution string = "Solver is called"
	solver := func(_ string) (string, error) {
		return solver_solution, nil
	}

	// Buffered channel of size 1
	channel := make(chan string, 1)

	// Mocked command-line arguments
	arguments := []string{"program_name", "--i", "some-input"}

	// Mocked reader. Returns an error
	var reader_error_msg string = "I do not want to read"
	reader := func(_ string) (string, error) {
		return "", errors.New(reader_error_msg)
	}

	solution_received, e := HandleArgumentsAndExecute(arguments, reader, solver, solver, channel)

	ts.AssertEqual(t, solution_received, false)
	ts.AssertError(t, e)
	if e != nil {
		ts.AssertEqual(t, e.Error(), reader_error_msg)
	}

	select {
	case x := <-channel:
		{
			t.Errorf("The passed channel is not expected to have anything queued. But a value was queued: %v", x)
		}
	default:
		{
			// All ok
		}
	}
}

func TestArgsHandle_WhenSolverReturnsAnError_SolutionIsNotPassedOverChannel(t *testing.T) {

	some_error_message := "Some error occurred"
	// Mocked solvers that return an error
	solver := func(_ string) (int, error) {
		return 0, errors.New(some_error_message)
	}

	// Mocked command-line arguments
	arguments1 := []string{"program_name", "--i", "some-input"}

	// Mocked reader. Returns successfully
	reader := func(_ string) (string, error) {
		return "", nil
	}

	// Buffered channel of size 1
	channel := make(chan int, 1)

	solution_received, e := HandleArgumentsAndExecute(arguments1, reader, solver, solver, channel)
	ts.AssertEqual(t, solution_received, false)
	ts.AssertError(t, e)
	if e != nil {
		ts.AssertEqual(t, e.Error(), some_error_message)
	}

	if len(channel) != 0 {
		t.Errorf("The channel was expected to receive no messages, but %d messages were received", len(channel))
	}
}

func TestArgsHandle_CheckThatCorrectSolverIsCalled(t *testing.T) {

	// Mocked solvers
	solver1 := func(_ string) (int, error) {
		return 1, nil
	}
	solver2 := func(_ string) (int, error) {
		return 2, nil
	}

	// Buffered channel of size 1
	channel := make(chan int, 1)

	// Mocked command-line arguments
	arguments1 := []string{"program_name", "--i", "some-input"}
	arguments2 := []string{"program_name", "--i", "some-input", "--solver2"}

	// Mocked reader. Returns successfully
	reader := func(_ string) (string, error) {
		return "", nil
	}

	// Call #1
	solution_received, e := HandleArgumentsAndExecute(arguments1, reader, solver1, solver2, channel)
	ts.Assert(t, solution_received)
	ts.AssertNoError(t, e)
	if solution_received {
		ts.AssertEqual(t, <-channel, 1)
	}

	// Call #2
	solution_received, e = HandleArgumentsAndExecute(arguments2, reader, solver1, solver2, channel)
	ts.Assert(t, solution_received)
	ts.AssertNoError(t, e)
	if solution_received {
		ts.AssertEqual(t, <-channel, 2)
	}
}

func TestArgsHandle_TestWholePipeline(t *testing.T) {
	// Buffered channel of size 1
	channel := make(chan int, 1)

	// Run #1
	solution_received, err := HandleArgumentsAndExecute(
		[]string{"program_name", "--i", "file"},
		dummyReader,
		sumSolver,
		maxSolver,
		channel,
	)
	ts.Assert(t, solution_received)
	ts.AssertNoError(t, err)
	if solution_received {
		ts.AssertEqual(t, <-channel, 15)
	}

	// Run #2
	solution_received, err = HandleArgumentsAndExecute(
		[]string{"program_name", "--i", "file", "--a"},
		dummyReader,
		sumSolver,
		maxSolver,
		channel,
	)
	ts.Assert(t, solution_received)
	ts.AssertNoError(t, err)
	if solution_received {
		ts.AssertEqual(t, <-channel, 7)
	}
}

// Returns the sum of all numbers
func sumSolver(numbers []int) (int, error) {
	var sum int = 0
	for _, x := range numbers {
		sum += x
	}
	return sum, nil
}

// Returns the largest number, if no numbers, returns -1
func maxSolver(numbers []int) (int, error) {
	var max int = -1
	var max_set bool = false
	for _, x := range numbers {
		if !max_set || max < x {
			max_set = true
			max = x
		}
	}
	return max, nil
}

// Dummy reader. Does not read anything, it only returns the fixed array of numbers
func dummyReader(_ string) ([]int, error) {
	return []int{3, 7, 5}, nil
}
