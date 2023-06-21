package argshandle

type configuration struct {
	input_filename    string
	display_help_flag bool
	use_second_solver bool
}

type argument_option struct {
	argument_keywords []string
	description       string
	handler           func(*int, *configuration, []string) error
}

// Options registered
func getSupportedArgumentOptions() []argument_option {
	return []argument_option{
		{
			argument_keywords: []string{"-h", "--h", "-help", "--help"},
			description:       "Print help in console",
			handler: func() func(*int, *configuration, []string) error {
				used_before := false
				return func(_ *int, c *configuration, _ []string) error {
					if used_before {
						return argument_option_used_multiple_times_error("printing help")
					}
					used_before = true
					c.display_help_flag = true
					return nil
				}
			}(),
		},
		{
			argument_keywords: []string{"--i", "-i", "--input", "--f", "-f", "--file"},
			description:       "Set the input file",
			handler: func() func(*int, *configuration, []string) error {
				used_before := false
				return func(i *int, c *configuration, args []string) error {
					if used_before {
						return argument_option_used_multiple_times_error("setting the input file")
					}
					used_before = true
					*i = *i + 1
					if *i >= len(args) {
						return argument_option_filename_not_provided_error()
					}
					c.input_filename = args[*i]
					return nil
				}
			}(),
		},
		{
			argument_keywords: []string{"--a", "-a", "--solver2"},
			description:       "Use second solver",
			handler: func() func(*int, *configuration, []string) error {
				used_before := false

				return func(i *int, c *configuration, s []string) error {
					if used_before {
						return argument_option_used_multiple_times_error("selecting the second solver")
					}
					used_before = true
					c.use_second_solver = true
					return nil
				}
			}(),
		},
	}
}
