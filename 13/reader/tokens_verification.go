package reader

import (
	f "aoc/functional"
	"regexp"
)

func verify_tokens(line_number int, tokens []string) error {
	for _, verification_function := range []func(int, []string) error{
		verify_tokens_are_not_empty,
		verify_starting_token,
		verify_successor_tokens,
		verify_last_token,
		verify_number_tokens,
		verify_bracket_closures,
		verify_packet_count,
	} {
		if verification_err := verification_function(line_number, tokens); verification_err != nil {
			return verification_err
		}
	}
	return nil
}

func verify_tokens_are_not_empty(line_number int, tokens []string) error {
	if len(tokens) == 0 {
		return no_tokens_parsing_error(line_number)
	}
	return nil
}

func verify_starting_token(line_number int, tokens []string) error {
	if tokens[0] != "[" {
		return invalid_starting_token_parser_error(line_number, tokens[0])
	}
	return nil
}

func verify_successor_tokens(line_number int, tokens []string) error {
	next_expected_token_types := func(previous_token string) []token_type {
		switch get_token_type(previous_token) {
		case left_par:
			return []token_type{left_par, right_par, number}
		case comma:
			return []token_type{number, left_par}
		case right_par:
			return []token_type{right_par, comma}
		case number:
			return []token_type{comma, right_par, left_par}
		default:
			return []token_type{}
		}
	}
	for i := 1; i < len(tokens); i++ {
		previous_token, token := tokens[i-1], tokens[i]
		if !f.ArrayContains(next_expected_token_types(previous_token), get_token_type(token)) {
			return invalid_successor_token_parser_error(
				line_number,
				tokens[:i],
				token,
				next_expected_token_types(previous_token),
			)
		}
	}
	return nil
}

func verify_last_token(line_number int, tokens []string) error {
	if tokens[len(tokens)-1] != "]" {
		return invalid_last_token_parser_error(line_number, tokens[len(tokens)-1])
	}
	return nil
}

func verify_number_tokens(line_number int, tokens []string) error {
	number_tokens := f.Filter(func(token string) bool { return get_token_type(token) == number }, tokens)
	valid_number_token_re := regexp.MustCompile(`^(?:0|[1-9]\d*)$`)
	for _, token := range number_tokens {
		if !valid_number_token_re.MatchString(token) {
			return bad_number_token_parsing_error(line_number, token)
		}
	}
	return nil
}

func verify_bracket_closures(line_number int, tokens []string) error {
	var bracket_level int = 0
	for _, token := range f.Filter(func(s string) bool { return s == "[" || s == "]" }, tokens) {
		switch token {
		case "[":
			bracket_level++
		case "]":
			bracket_level--
		}
		if bracket_level < 0 {
			return too_many_closed_brackets_parsing_error(line_number)
		}
	}

	if bracket_level != 0 {
		return brackets_improperly_closed_parsing_error(line_number, bracket_level)
	}
	return nil
}

func verify_packet_count(line_number int, tokens []string) error {
	var packet_count, bracket_level int = 0, 0
	for _, token := range f.Filter(func(s string) bool { return s == "[" || s == "]" }, tokens) {

		switch token {
		case "[":
			bracket_level++
		case "]":
			bracket_level--
			if bracket_level == 0 {
				packet_count++
			}
		}
	}

	if packet_count != 1 {
		return packet_count_parsing_error(line_number, packet_count)
	}
	return nil
}
