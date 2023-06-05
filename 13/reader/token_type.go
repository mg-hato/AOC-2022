package reader

import "regexp"

type token_type = int

const (
	left_par token_type = iota
	right_par
	comma
	number
	unknown
)

func get_token_type(token string) token_type {
	switch token {
	case "[":
		return left_par
	case ",":
		return comma
	case "]":
		return right_par
	default:
		if regexp.MustCompile(`^\d+$`).MatchString(token) {
			return number
		}
		return unknown
	}
}

func get_token_type_name(tt token_type) string {
	switch tt {
	case comma:
		return "comma"
	case left_par:
		return `"["`
	case right_par:
		return `"]"`
	case number:
		return "number"
	default:
		return "unknown"
	}
}
