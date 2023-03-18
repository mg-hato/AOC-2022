package testers

import (
	"fmt"
	"os"
	"testing"
)

// Tests reader for "good" and "error" inputs
//
//   - "Good" inputs are the ones that result in successful read (no errors) of the file and return a value of type T.
//     The good tests will aim to check that the returned value matches the expected.
//
//   - "Error" inputs are the ones that are expected to result in error when reading the file.
//     The error tests will aim to check that reading of the file will indeed result in an error.
type ReaderTester[T any] struct {
	reader      func(string) (T, error) // Reader function
	reader_name string                  // Reader name (to be displayed when running tests)

	equals func(T, T) bool // Equality function used for good-input tests

	good_tests        map[int]T // Good tests: mapping from test-number to expected returned value
	good_file_pattern string    // A file pattern used for good-tests

	error_tests        map[int]string // Error tests: mapping from test-number to a reason why the file reading should result in an error
	error_file_pattern string         // A file pattern used for error-tests
}

// Returns a reader tester with some defaults pre-defined.
// User is expected to provide an equality-function for comparing T-types.
func DefaultReaderTester[T any](
	reader func(string) (T, error),
	reader_name string,
) ReaderTester[T] {
	tester := ReaderTester[T]{
		reader:             reader,
		reader_name:        reader_name,
		good_tests:         map[int]T{},
		good_file_pattern:  "./test/input.%d",
		error_tests:        map[int]string{},
		error_file_pattern: "./test/bad-input.%d",
	}
	return tester
}

// Returns a reader tester with some defaults pre-defined.
// The equality-function for T-types is provided out-of-the-box (type T must be comparable)
func DefaultReaderTesterForComparableTypes[T comparable](
	reader func(string) (T, error),
	reader_name string,
) ReaderTester[T] {
	tester := DefaultReaderTester(reader, reader_name)
	tester.equals = func(lhs, rhs T) bool { return lhs == rhs }
	return tester
}

// Override the pattern for filenames that contain input for "good" tests.
func (tester *ReaderTester[T]) OverrideGoodFilePattern(new_pattern string) *ReaderTester[T] {
	tester.good_file_pattern = new_pattern
	return tester
}

// Override the pattern for filenames that contain input for "error" tests.
func (tester *ReaderTester[T]) OverrideErrorFilePattern(new_pattern string) *ReaderTester[T] {
	tester.error_file_pattern = new_pattern
	return tester
}

// !!DEPRECATED!! Add a tests for good input: Add the expected values that the reader should return
func (tester *ReaderTester[T]) AddGoodInputTests(expected_values ...T) {
	for _, expected_value := range expected_values {
		test_number := len(tester.good_tests) + 1
		tester.good_tests[test_number] = expected_value
	}
}

// Add a test for good input: Add the expected value that the reader should return
func (tester *ReaderTester[T]) AddGoodInputTest(expected_value T) *ReaderTester[T] {
	test_number := len(tester.good_tests) + 1
	tester.good_tests[test_number] = expected_value
	return tester
}

// !!DEPRECATED!! Add a tests for error input: Add the reasons (to be displayed in the case of test-failure) why reading the input is expected to result in an error
func (tester *ReaderTester[T]) AddErrorInputTests(reasons ...string) {
	for _, reason := range reasons {
		test_number := len(tester.error_tests) + 1
		tester.error_tests[test_number] = reason
	}
}

// Add a test for error input: Add the reason (to be displayed in the case of test-failure) why reading the input is expected to result in an error
func (tester *ReaderTester[T]) AddErrorInputTest(reason string) *ReaderTester[T] {
	test_number := len(tester.error_tests) + 1
	tester.error_tests[test_number] = reason
	return tester
}

// Provide equality-function for T-type (used only in good-input tests)
func (tester *ReaderTester[T]) ProvideEqualityFunctionForTypeT(equality_func func(T, T) bool) {
	tester.equals = equality_func
}

// Run good-input tests
func (tester *ReaderTester[T]) RunGoodInputTests(t *testing.T) {
	for test_number, expected_value := range tester.good_tests {

		if tester.equals == nil {
			t.Errorf("Error: Equality function for T-type not defined")
			return
		}

		filename := fmt.Sprintf(tester.good_file_pattern, test_number)
		result, err := tester.reader(filename)

		if err != nil || !tester.equals(result, expected_value) {
			t.Errorf("Good-input test #%d failed: %s(\"%s\")", test_number, tester.reader_name, filename)

			if err != nil {
				t.Error("An unexpected error occurred during reading the input. Error is given below")
				t.Error(err)
			} else {
				t.Error("Returned value does not match the expected value")
				t.Errorf("Returned: %v", result)
				t.Errorf("Expected: %v", expected_value)
			}
		}
	}
}

// Run error-input tests
func (tester *ReaderTester[T]) RunErrorInputTests(t *testing.T) {
	for test_number, reason := range tester.error_tests {

		filename := fmt.Sprintf(tester.error_file_pattern, test_number)

		// Test that the file exists:
		if file, err := os.Open(filename); err != nil {
			t.Errorf("Error-input test #%d failed: %s(\"%s\")", test_number, tester.reader_name, filename)
			t.Error("File could not be opened for reading. Does the file even exists?")
			continue
		} else {
			file.Close()
		}

		// If the file exists, check whether reading the input yields errors (as expected)
		if _, err := tester.reader(filename); err == nil {
			t.Errorf("Error-input test #%d failed: %s(\"%s\")", test_number, tester.reader_name, filename)
			t.Error("No error has been returned from reading bad input")
			t.Errorf("Input is expected to return an error because: %s", reason)
		}
	}
}

// Run both group of tests
func (tester *ReaderTester[T]) RunBothGoodAndErrorInputTests(t *testing.T) {
	tester.RunGoodInputTests(t)
	tester.RunErrorInputTests(t)
}
