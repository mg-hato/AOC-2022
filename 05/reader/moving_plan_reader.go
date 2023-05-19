package reader

import (
	m "aoc/day05/models"
	e "aoc/envelope"
	f "aoc/functional"
	"aoc/reading"
	"regexp"
	"strconv"
)

type moving_plan_reader struct {
	err error

	line_number int

	current_reading_mode reading_mode
	line_processors      map[reading_mode]func(string)

	container_row_lines []string

	container_row_re *regexp.Regexp
	stack_ids_re     *regexp.Regexp

	move_re  *regexp.Regexp
	empty_re *regexp.Regexp

	plan m.MovingPlan
}

// Constructor function for section assignments list reader
func MovingPlanReader() reading.ReaderAoC2022[e.Envelope[m.MovingPlan]] {
	mpr := &moving_plan_reader{
		err: nil,

		line_number: 0,

		current_reading_mode: read_containers,
		line_processors:      make(map[reading_mode]func(string)),

		container_row_lines: make([]string, 0),

		container_row_re: regexp.MustCompile(`^ *(?:\[[A-Z]\] *)*$`),
		stack_ids_re:     regexp.MustCompile(`^ *(?:\d+ *)+$`),

		move_re:  regexp.MustCompile(`^ *move (\d+) from (\d+) to (\d+) *$`),
		empty_re: regexp.MustCompile(`^ *$`),

		plan: m.MovingPlan{
			StartingContainers: make([]string, 0),

			Moves: make([]m.Move, 0),
		},
	}
	mpr.line_processors[read_containers] = mpr.processContainerLine
	mpr.line_processors[read_move_instructions] = mpr.processMoveLine
	return mpr
}

func (mpr moving_plan_reader) Error() error {
	return mpr.err
}

func (mpr moving_plan_reader) PerformFinalValidation() error {
	return nil
}

func (mpr moving_plan_reader) Done() bool {
	return mpr.Error() != nil
}

func (mpr *moving_plan_reader) ProvideLine(line string) {
	mpr.line_number++
	mpr.line_processors[mpr.current_reading_mode](line)
}

// Processes container line or stack id line (which represents the end of container reading mode)
func (mpr *moving_plan_reader) processContainerLine(line string) {
	// If the line matches container row RE, save it for later
	if mpr.container_row_re.MatchString(line) {
		mpr.container_row_lines = append(mpr.container_row_lines, line)
		return
	}

	// Otherwise, it should be a line with stack IDs; If that is not the case, produce error
	if !mpr.stack_ids_re.MatchString(line) {
		mpr.err = bad_line_reading_container_rows_error(mpr.line_number, line)
		return
	}

	stacks, verification_err := verify_stack_ids_and_container_rows(line, f.Reverse(mpr.container_row_lines))
	if verification_err != nil {
		mpr.err = verification_err
		return
	}
	mpr.plan.StartingContainers = stacks
	mpr.current_reading_mode = read_move_instructions
}

func (mpr *moving_plan_reader) processMoveLine(line string) {
	// If line is empty, ignore it
	if mpr.empty_re.MatchString(line) {
		return
	}

	// Otherwise, the line must contain a move instruction. If not, produce an error
	matches := mpr.move_re.FindStringSubmatch(line)
	if len(matches) != 4 {
		mpr.err = bad_line_reading_move_instructions_error(mpr.line_number, line)
		return
	}

	numbers := f.Map(func(s string) int { i, _ := strconv.Atoi(s); return i }, matches[1:])
	move := m.MakeMove(numbers[0], numbers[1], numbers[2])

	verification_err := verify_move_instruction(move, len(mpr.plan.StartingContainers), mpr.line_number)
	if verification_err != nil {
		mpr.err = verification_err
		return
	}
	mpr.plan.Moves = append(mpr.plan.Moves, move)
}

func (mpr moving_plan_reader) FinishAndGetInputData() e.Envelope[m.MovingPlan] {
	return m.CreateMovingPlanEnvelope(mpr.plan)
}
