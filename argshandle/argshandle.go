package argshandle

import (
	c "aoc/common"
	"fmt"
	"log"
	"os"
	"strings"
)

// Handle provided arguments.
// If provided arguments define executable configuration, input will be read from a specified file.
// If input is read successfully, selected solver will be applied and solution sent over the channel provided.
// Returns (b, e):
//   - b is true iff solution was sent over the chanel provided
//   - e is the error that occurred during reading input (if any)
func HandleArgumentsAndExecute[IN any, OUT any](
	arguments []string,
	readInput func(string) (IN, error),
	solver1 func(IN) (OUT, error),
	solver2 func(IN) (OUT, error),
	outputChannel chan OUT,
) (bool, error) {

	// Starting configuration for the app execution
	conf := configuration{}

	supported_argument_options := getSupportedArgumentOptions()

	argument_to_argument_option := c.CreateKeyValueMap(
		c.FlatMap(
			func(option argument_option) []c.Pair[string, argument_option] {
				return c.Map(
					func(keyword string) c.Pair[string, argument_option] { return c.MakePair(keyword, option) },
					option.argument_keywords,
				)
			},
			supported_argument_options,
		),
		c.GetFirst[string, argument_option],
		c.GetSecond[string, argument_option],
	)

	// process all arguments
	var i int = 1
	for i < len(arguments) {
		if !c.MapContains[string, argument_option](arguments[i])(argument_to_argument_option) {
			return false, invalid_argument_error(arguments[i])
		} else if err := argument_to_argument_option[arguments[i]].handler(&i, &conf, arguments); err != nil {
			return false, err
		}
		i++
	}

	// If display-help flag is set, display help and return
	if conf.display_help_flag {
		display_help(supported_argument_options)
		return false, nil
	}

	if conf.input_filename == "" {
		return false, input_file_not_provided_error()
	}

	// try to read the input
	input, e := readInput(conf.input_filename)
	if e != nil {
		return false, e
	}

	// At this point, input is successfully retrieved.
	// Select solver based on configuration inferred from arguments
	selected_solver := solver1
	if conf.use_second_solver {
		selected_solver = solver2
	}

	solution, e := selected_solver(input)
	if e != nil {
		return false, e
	}

	outputChannel <- solution

	// return true to indicate that solution has been sent over a channel
	return true, nil
}

// `Program` is usually how a program will be executed
// for any of the Advent of Code 2022 problems
func Program[IN any, OUT any](
	reader func(string) (IN, error),
	solver1 func(IN) (OUT, error),
	solver2 func(IN) (OUT, error),
) {
	var channel chan OUT = make(chan OUT, 1)
	ok, err := HandleArgumentsAndExecute(os.Args, reader, solver1, solver2, channel)

	if err != nil {
		log.Fatal(err)
	}

	if ok {
		fmt.Println("Solution received")
		fmt.Printf("\t%v\n", <-channel)
	}
}

func display_help(argument_options []argument_option) {
	params_to_print := c.Map(
		func(ao argument_option) string {
			return strings.Join(ao.argument_keywords, ", ")
		},
		argument_options,
	)

	// find maximum length among `paramsToPrint`
	max_length := len(c.MaximumBy(
		params_to_print,
		func(lhs, rhs string) bool { return len(lhs) < len(rhs) },
	))

	// padding = distance in ' ' between options & descriptions
	var padding int = 5

	// margin = number of spaces before the first option print
	var margin int = 3
	var margin_str string = strings.Repeat(" ", margin)

	c.ForEach(
		func(pair c.Pair[string, argument_option]) {
			gap := max_length - len(pair.First)
			fmt.Printf(
				"%s%s%s%s\n",
				margin_str,
				pair.First,
				strings.Repeat(" ", padding+gap),
				pair.Second.description,
			)
		},
		c.Zip(params_to_print, argument_options),
	)
}
