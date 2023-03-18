package argshandle

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestArgsHandle_WhenReaderSuccessfullyReads_CheckThatSolverIsCalled(t *testing.T) {

	// Mocked solver
	var expectedOutcome string = "Solver is called"
	solver := func(_ string) (string, error) {
		return expectedOutcome, nil
	}

	// Buffered channel of size 1
	channel := make(chan string, 1)

	// Mocked command-line arguments
	arguments := []string{"program_name", "--i", "some-input"}

	// Mocked reader. Returns successfully.
	reader := func(_ string) (string, error) {
		return "Reader reads", nil
	}

	isSolutionPassed, e := HandleArgumentsAndExecute(arguments, reader, solver, solver, channel)

	if e != nil {
		t.Errorf("HandleArgumentsAndExecute has encountered an error: %v", e)
	}

	if !isSolutionPassed {
		t.Error("Solution was not passed, but is expected to be")
	}

	if outcome := <-channel; expectedOutcome != outcome {
		t.Errorf("Returned solution does not match.\n\tExpected: %s\n\tActual: %s\n", expectedOutcome, outcome)
	}
}

func TestArgsHandle_WhenReaderUnsuccessful_ExpectNoSolution(t *testing.T) {

	// Mocked solver
	var solverSolution string = "Solver is called"
	solver := func(_ string) (string, error) {
		return solverSolution, nil
	}

	// Buffered channel of size 1
	channel := make(chan string, 1)

	// Mocked command-line arguments
	arguments := []string{"program_name", "--i", "some-input"}

	// Mocked reader. Returns an error
	var readerErrorMessage string = "I do not want to read"
	reader := func(_ string) (string, error) {
		return "", errors.New(readerErrorMessage)
	}

	isSolutionPassed, e := HandleArgumentsAndExecute(arguments, reader, solver, solver, channel)

	if isSolutionPassed {
		t.Errorf("Not matching boolean value returned.\n\tExpected: %v\n\tActual: %v\n", false, isSolutionPassed)
	}

	var expectedError error = errors.New(readerErrorMessage)
	if e.Error() != expectedError.Error() {
		t.Errorf("Not matching error (messages) returned.\n\tExpected: %v\n\tActual: %v\n", expectedError, e)
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

	isSolutionPassed, e := HandleArgumentsAndExecute(arguments1, reader, solver, solver, channel)

	if isSolutionPassed {
		t.Error("The boolean flag \"isSolutionPassed\" was expected to be false")
	}

	if e.Error() != some_error_message {
		t.Error("Error message was not the one expected")
		t.Errorf("Actual: %s", e.Error())
		t.Errorf("Expected: %s", some_error_message)
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
	isSolutionPassed, e := HandleArgumentsAndExecute(arguments1, reader, solver1, solver2, channel)

	// Expect that solution (number 1) is passed with no errors
	if outcome := <-channel; !isSolutionPassed || e != nil || outcome != 1 {
		t.Errorf(
			"Actual values do not match expected ones.\n\tActual boolean flag: %v\n\tActual error: %v\n\tActual outcome: %v\n",
			isSolutionPassed, e, outcome,
		)
	}

	// Call #2
	isSolutionPassed, e = HandleArgumentsAndExecute(arguments2, reader, solver1, solver2, channel)

	// Expect that solution (number 2) is passed with no errors
	if outcome := <-channel; !isSolutionPassed || e != nil || outcome != 2 {
		t.Errorf(
			"Actual values do not match expected ones.\n\tActual boolean flag: %v\n\tActual error: %v\n\tActual outcome: %v\n",
			isSolutionPassed, e, outcome,
		)
	}
}

// Simple numbers reader. It reads the first line of the file and returns all integers on it,
// where the integers are separated by one space
func numbersReader(filename string) ([]int, error) {
	file, e := os.Open(filename)
	if e != nil {
		return nil, e
	}

	// Scan the first line & check for errors
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	if !scanner.Scan() || scanner.Err() != nil {
		file.Close()
		return nil, errors.New("Issues with reading the test file")
	}

	// Read the first line
	numbers_strings := strings.Split(scanner.Text(), " ")
	file.Close()

	// Parse it & get the numbers
	numbers := make([]int, len(numbers_strings))
	for i, v := range numbers_strings {
		numbers[i], _ = strconv.Atoi(v)
	}

	return numbers, nil
}

func TestArgsHandle_TestWholePipeline(t *testing.T) {
	// Buffered channel of size 1
	channel := make(chan int, 1)

	// Run #1
	HandleArgumentsAndExecute([]string{"program_name", "--i", "file"}, dummyReader, sumSolver, maxSolver, channel)
	if outcome := <-channel; outcome != 15 {
		t.Errorf("Run #1. Incorrect value returned. Expected %v, actual %v", 15, outcome)
	}

	// Run #2
	HandleArgumentsAndExecute([]string{"program_name", "--i", "file", "--a"}, dummyReader, sumSolver, maxSolver, channel)
	if outcome := <-channel; outcome != 7 {
		t.Errorf("Run #2. Incorrect value returned. Expected %v, actual %v", 7, outcome)
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
