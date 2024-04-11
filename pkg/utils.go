package pkg

import "strings"

// SplitArg splits the argument by comma and returns the result
func SplitArg(arg string) (result []string) {
	value := strings.Split(arg, ",")
	if len(value) > 1 {
		return value
	}
	return []string{}
}
