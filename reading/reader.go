package reading

// Task input reader abstraction
type ReaderAoC2022[T any] interface {
	Error() error                  // Returns an error (if any) encountered while processing input lines
	PerformFinalValidation() error // Post-reading input validation: returns error if validation fails
	Done() bool                    // Returns true iff reader is done (either because no more lines of input are needed or because an error happened)
	ProvideLine(string)            // Provide a line of input to the reader
	FinishAndGetInputData() T      // Get the input data that is ready for solver's consumption
}
