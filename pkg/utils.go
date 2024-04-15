package pkg

import (
	"errors"
	"strings"
	"time"
)

// SplitArg splits the argument by comma and returns the result
func SplitArg(arg string) (result []string) {
	value := strings.Split(arg, ",")
	if len(value) > 1 {
		return value
	}
	return []string{}
}

// DateFormats convert the php date formats to go date formats
func DateFormats(format string) (string, error) {
	// check if the format is empty or have a valid value
	if format == "" {
		return "", errors.New("empty format")
	}

	formats := []string{"Y", "y", "m", "n", "d", "j", "H", "G", "h", "g", "i", "s", "A", "a", "M", "F", "l", "D", "w", "N", "S", "z", "W", "t", "L", "o", "B", "g", "U", "e", "I", "O", "P", "T", "Z", "c"}
	replacements := []string{"2006", "06", "01", "1", "02", "2", "15", "15", "03", "3", "04", "05", "PM", "pm", "Jan", "January", "Monday", "Mon", "1", "1", "th", "1", "1", "28", "false", "2006", "1", "3", "1", "MST", "false", "-0700", "-07:00", "MST", "1", time.RFC3339}
	// valueList := map[string]string{}

	values := strings.Split(format, "")

	for k, value := range values {
		for i, f := range formats {
			if value == f {
				values[k] = replacements[i]
			}
		}
	}

	return strings.Join(values, ""), nil
}
