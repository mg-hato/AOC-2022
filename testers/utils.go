package testers

import (
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
	formatted_keywords := make([]string, len(keywords))
	for i := 0; i < len(keywords); i++ {
		formatted_keywords[i] = fmt.Sprintf(`"%s"`, keywords[i])
	}
	return fmt.Sprintf("[%s]", strings.Join(formatted_keywords, ", "))
}
