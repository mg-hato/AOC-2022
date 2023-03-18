package argshandle

import (
	. "aoc/functional"
	"fmt"
	"log"
	"os"
	"strings"
)

type configuration struct {
	inputFilename   string
	dispalyHelp     bool
	useSecondSolver bool
}

type argumentOption struct {
	parameters  []string
	description string
	handler     func(*int, *configuration, []string)
}

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

	// Options registered
	supportedArgumentOptions := []argumentOption{
		{
			parameters:  []string{"-h", "--h", "-help", "--help"},
			description: "Print help in console",
			handler:     func(_ *int, c *configuration, _ []string) { c.dispalyHelp = true },
		},
		{
			parameters:  []string{"--i", "-i", "--input", "--f", "-f", "--file"},
			description: "Set the input file",
			handler: func(i *int, c *configuration, args []string) {
				*i = *i + 1
				if *i < len(args) {
					c.inputFilename = args[*i]
				}
			},
		},
		{
			parameters:  []string{"--a", "-a", "--solver2"},
			description: "Use second solver",
			handler:     func(i *int, c *configuration, s []string) { c.useSecondSolver = true },
		},
	}
	mapping := getParameterToArgOptionMapping(supportedArgumentOptions)

	// process all parameters
	var i int = 1
	for i < len(arguments) {
		if o, ok := mapping[arguments[i]]; ok {
			o.handler(&i, &conf, arguments)
		}
		i++
	}

	// If input filename is not provided or display-help flag is set, display help and return
	if conf.inputFilename == "" || conf.dispalyHelp {
		displayHelp(supportedArgumentOptions)
		return false, nil
	}

	// try to read the input
	input, e := readInput(conf.inputFilename)
	if e != nil {
		return false, e
	}

	// At this point, input is successfully retrieved. Close the file.
	// Select solver based on configuration inferred
	selectedSolver := solver1
	if conf.useSecondSolver {
		selectedSolver = solver2
	}

	solution, e := selectedSolver(input)
	if e != nil {
		return false, e
	}

	outputChannel <- solution

	// return true to indicate that solution has been sent over a channel
	return true, nil
}

// `AoC2022DefaultProgram` is usually how a program will be executed
// for any of the Advent of Code 2022 problems
func AoC2022DefaultProgram[IN any, OUT any](
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

// Convert array of `argumentOption` into mapping: (parameter) => argumentOption
func getParameterToArgOptionMapping(argOpts []argumentOption) map[string]argumentOption {
	mapping := make(map[string]argumentOption)
	for _, argOpt := range argOpts {
		for _, parameter := range argOpt.parameters {
			mapping[parameter] = argOpt
		}
	}
	return mapping
}

func displayHelp(argumentOptions []argumentOption) {
	paramsToPrint := Map(func(ao argumentOption) string {
		return strings.Join(ao.parameters, ", ")
	}, argumentOptions)

	// find maximum length among `paramsToPrint`
	maxLength := len(Maximum(paramsToPrint, func(lhs, rhs string) bool { return len(lhs) < len(rhs) }))

	// padding = distance in ' ' between options & descriptions
	var padding int = 5

	// margin = number of spaces before the first option print
	var margin int = 3
	var margin_str string = strings.Repeat(" ", margin)

	ForEach(
		func(pair Pair[string, argumentOption]) {
			gap := maxLength - len(pair.First)
			fmt.Printf("%s%s%s%s\n", margin_str, pair.First, strings.Repeat(" ", padding+gap), pair.Second.description)
		},
		Zip(paramsToPrint, argumentOptions),
	)
}
