package reader

import (
	m "aoc/d13/models"
	f "aoc/functional"
	"fmt"
	"strings"
)

// prefix functions
func reading_error_prefix(line_number int) string {
	return fmt.Sprintf("reading error on line #%d", line_number)
}

func parsing_error_prefix(line_number int) string {
	return fmt.Sprintf("packet-parsing error on line #%d", line_number)
}

func final_validation_error_prefix() string {
	return "reader final validation error"
}

// error functions
func expected_empty_line_error(line_number int, line string) error {
	return fmt.Errorf(
		`%s: an empty line was expected. Actual line "%s"`,
		reading_error_prefix(line_number),
		line,
	)
}

func expected_first_packet_error(line_number int, line string) error {
	return fmt.Errorf(
		`%s: an empty line or first packet was expected. Actual line "%s"`,
		reading_error_prefix(line_number),
		line,
	)
}

func expected_second_packet_error(line_number int, line string) error {
	return fmt.Errorf(
		`%s: the second packet was expected. Actual line "%s"`,
		reading_error_prefix(line_number),
		line,
	)
}

func no_tokens_parsing_error(line_number int) error {
	return fmt.Errorf(
		`%s: no tokens were extracted`,
		parsing_error_prefix(line_number),
	)
}

func invalid_starting_token_parser_error(line_number int, token string) error {
	return fmt.Errorf(
		`%s: first token must be "[" but actual first token is "%s"`,
		parsing_error_prefix(line_number),
		token,
	)
}

func invalid_last_token_parser_error(line_number int, token string) error {
	return fmt.Errorf(
		`%s: last token must be "]" but actual last token is "%s"`,
		parsing_error_prefix(line_number),
		token,
	)
}

func invalid_successor_token_parser_error(
	line_number int,
	previous_tokens []string,
	token string,
	expected_token_types []token_type,
) error {
	return fmt.Errorf(
		`%s: after previous sequence of tokens "%s" the expected token was %s, but the actual token is "%s"`,
		parsing_error_prefix(line_number),
		strings.Join(previous_tokens, ""),
		strings.Join(f.Map(get_token_type_name, expected_token_types), " or "),
		token,
	)
}

func bad_number_token_parsing_error(line_number int, bad_token string) error {
	return fmt.Errorf(
		`%s: bad number token encountered "%s"`,
		parsing_error_prefix(line_number),
		bad_token,
	)
}

func too_many_closed_brackets_parsing_error(line_number int) error {
	return fmt.Errorf(
		"%s: too many closed brackets",
		parsing_error_prefix(line_number),
	)
}

func brackets_improperly_closed_parsing_error(line_number, open_bracket_count int) error {
	return fmt.Errorf(
		`%s: there %s %d outstanding unclosed bracket%s`,
		parsing_error_prefix(line_number),
		func() string {
			switch open_bracket_count {
			case 1:
				return "is"
			default:
				return "are"
			}
		}(),
		open_bracket_count,
		func() string {
			switch open_bracket_count {
			case 1:
				return ""
			default:
				return "s"
			}
		}(),
	)
}

func packet_count_parsing_error(line_number, packet_count int) error {
	return fmt.Errorf(
		"%s: exactly one top-level packet is expected per line, but %d packets were detected",
		parsing_error_prefix(line_number),
		packet_count,
	)
}

func second_packet_missing_final_validation_error() error {
	return fmt.Errorf(
		`%s: previous packet is missing its second packet to form a packet pair`,
		final_validation_error_prefix(),
	)
}

func divider_packet_detected_final_validation_error(divider_packet m.Packet, packet_number int) error {
	return fmt.Errorf(
		`%s: among the received packets a divider packet %s was detected as packet number #%d`,
		final_validation_error_prefix(),
		divider_packet,
		packet_number,
	)
}
