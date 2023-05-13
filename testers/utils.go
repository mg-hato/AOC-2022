package testers

import (
	f "aoc/functional"
	"fmt"
	"strings"
)

// Formats message with a given prefix. If prefix is non-empty, prefix will be added.
func format_with_prefix(prefix, message string) string {
	prefix = strings.TrimSpace(prefix)
	if len(prefix) == 0 {
		return message
	} else {
		return fmt.Sprintf("%s: %s", prefix, message)
	}
}

// Formats array of strings
func format_string_array(keywords []string) string {
	comma_separated_keywords := strings.Join(
		f.Map(
			func(keyword string) string { return fmt.Sprintf(`"%s"`, keyword) },
			keywords,
		), ", ",
	)
	return fmt.Sprintf("[%s]", comma_separated_keywords)
}
