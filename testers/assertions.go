package testers

import (
	"os"
	"testing"
)

func AssertEqual[T comparable](t *testing.T, actual, expected T) bool {
	result := actual == expected
	if !result {
		t.Errorf("Failed assertion: actual %v is not equal to expected %v", actual, expected)
	}
	return result
}

func AssertEqualWithEqFunc[T any](t *testing.T, actual, expected T, eq_func func(T, T) bool) bool {
	result := eq_func(actual, expected)
	if !result {
		t.Errorf("Failed assertion: actual %v is not equal to expected %v", actual, expected)
	}
	return result
}

func AssertNoError(t *testing.T, e error) bool {
	result := e == nil
	if !result {
		t.Errorf("Failed assertion: expected no error, but received error with the following message => %v", e)
	}
	return result
}

func AssertFileCanBeOpenedForReading(t *testing.T, filename string) bool {
	file, err := os.Open(filename)
	result := err == nil
	if !result {
		t.Errorf(`File "%s" cannot be opened for reading. Error: %s`, filename, err.Error())
	} else {
		file.Close()
	}
	return result
}

func TestThat[T any](inputs []T, test_func func(T)) {
	for _, input := range inputs {
		test_func(input)
	}
}
