package argshandle

import (
	"fmt"
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
	solver1 func(IN) OUT,
	solver2 func(IN) OUT,
	outputChannel chan OUT,
) (bool, error) {

	// Starting configuration for the app execution
	conf := configuration{inputFilename: "", dispalyHelp: false, useSecondSolver: false}

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

	outputChannel <- selectedSolver(input)

	// return true to indicate that solution has been sent over a channel
	return true, nil
}

func check(e error, f *os.File) bool {
	if e != nil {
		f.Close()
	}
	return e == nil
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
	optionsToPrint := make([]string, 0)
	for _, argOpt := range argumentOptions {
		optionsToPrint = append(optionsToPrint, strings.Join(argOpt.parameters, ", "))
	}

	// find maximum length of the options
	maxLength := 0
	for _, otp := range optionsToPrint {
		if maxLength < len(otp) {
			maxLength = len(otp)
		}
	}

	// padding = distance in ' ' between options & descriptions
	padding := 5

	// margin = number of spaces before the first option print
	margin := 3
	margin_str := strings.Repeat(" ", margin)

	// print out help
	for i, otp := range optionsToPrint {
		gap := maxLength - len(otp)
		fmt.Printf("%s%s%s%s\n", margin_str, otp, strings.Repeat(" ", padding+gap), argumentOptions[i].description)
	}

}
