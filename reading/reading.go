package reading

import (
	"bufio"
	"os"
)

func ReadWith[T any](
	reader_supplier func() ReaderAoC2022[T],
) func(string) (T, error) {
	return func(filename string) (T, error) {
		reader := reader_supplier()
		// Try to open the file
		file, err := os.Open(filename)
		if err != nil {
			return reader.FinishAndGetInputData(), err
		}

		// "schedule" file closing after the function is finished
		defer file.Close()

		// prepare line-scanner
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)

		for scanner.Scan() && !reader.Done() {
			reader.ProvideLine(scanner.Text())
		}

		// If the loop is left because scanner had an error
		if scanner.Err() != nil {
			return reader.FinishAndGetInputData(), scanner.Err()
		}

		// If reader encountered any error
		if reader.Error() != nil {
			return reader.FinishAndGetInputData(), reader.Error()
		}

		// Perform any last validation of the input data before finishing
		if err = reader.PerformFinalValidation(); err != nil {
			return reader.FinishAndGetInputData(), err
		}

		// Finish successfully
		return reader.FinishAndGetInputData(), nil
	}
}
