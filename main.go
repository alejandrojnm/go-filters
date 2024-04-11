package gofilters

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"unicode"

	"github.com/alejandrojnm/go-filters/pkg"
)

// Pluralize returns a plural suffix if the value is not 1, '1', or an object of
func Pluralize(size int, word string, arg ...string) string {

	/*
	   Return a plural suffix if the value is not 1, '1', or an object of
	   length 1. By default, use 's' as the suffix:
	   * If value is 0, vote{{ value|pluralize }} display "votes".
	   * If value is 1, vote{{ value|pluralize }} display "vote".
	   * If value is 2, vote{{ value|pluralize }} display "votes".

	   If an argument is provided, use that string instead:
	   * If value is 0, class{{ value|pluralize:"es" }} display "classes".
	   * If value is 1, class{{ value|pluralize:"es" }} display "class".
	   * If value is 2, class{{ value|pluralize:"es" }} display "classes".

	   If the provided argument contains a comma, use the text before the comma
	   for the singular case and the text after the comma for the plural case:
	   * If value is 0, cand{{ value|pluralize:"y,ies" }} display "candies".
	   * If value is 1, cand{{ value|pluralize:"y,ies" }} display "candy".
	   * If value is 2, cand{{ value|pluralize:"y,ies" }} display "candies".
	*/

	switch {
	case size == 0:
		if len(arg) == 0 {
			return fmt.Sprintf("%ss", word)
		} else {
			bits := pkg.SplitArg(arg[0])
			if len(bits) > 1 {
				return fmt.Sprintf("%s%s", word, bits[1])
			} else {
				return fmt.Sprintf("%s%s", word, arg[0])
			}
		}
	case size == 1:
		if len(arg) == 0 {
			return word
		} else {
			bits := pkg.SplitArg(arg[0])
			if len(bits) > 1 {
				return fmt.Sprintf("%s%s", word, bits[0])
			} else {
				return word
			}
		}
	case size > 1:
		if len(arg) == 0 {
			return fmt.Sprintf("%ss", word)
		} else {
			bits := pkg.SplitArg(arg[0])
			if len(bits) > 1 {
				return fmt.Sprintf("%s%s", word, bits[1])
			} else {
				return fmt.Sprintf("%s%s", word, arg[0])
			}
		}
	}

	return ""
}

// AddSlashes adds slashes before quotes. Useful for escaping strings in CSV, for example.
// Less useful for escaping JavaScript; use the “escapejs“ filter instead.
func AddSlashes(value string) string {
	value = strings.ReplaceAll(value, "\\", "\\\\")
	value = strings.ReplaceAll(value, "\"", "\\\"")
	value = strings.ReplaceAll(value, "'", "\\'")
	return value
}

// CapFirst capitalizes the first character of the value.
func CapFirst(value string) string {
	if value == "" {
		return ""
	}
	r := []rune(value)
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}

// FloatFormat displays a float to a specified number of decimal places.
func FloatFormat(text float64, arg int) string {
	if arg < 0 {
		if text == math.Trunc(text) {
			return fmt.Sprintf("%.0f", text)
		}
		arg = -arg
	}
	format := "%." + strconv.Itoa(arg) + "f"
	return fmt.Sprintf(format, text)
}
