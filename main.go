package gofilters

import (
	"fmt"
	"html"
	"math"
	"math/rand/v2"
	"sort"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/alejandrojnm/go-filters/pkg"
	"github.com/hako/durafmt"
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

// LineNumbers displays text with line numbers.
func LineNumbers(value string, autoescape bool) string {
	lines := strings.Split(value, "\n")
	width := len(fmt.Sprintf("%d", len(lines)))
	format := fmt.Sprintf("%%0%dd. %%s", width)
	for i, line := range lines {
		if autoescape {
			line = html.EscapeString(line)
		}
		lines[i] = fmt.Sprintf(format, i+1, line)
	}
	return strings.Join(lines, "\n")
}

// Lower converts a string into all lowercase.
func Lower(value string) string {
	return strings.ToLower(value)
}

// Upper converts a string into all uppercase.
func Upper(value string) string {
	return strings.ToUpper(value)
}

// IreEncode encodes a string as a quoted-printable string.
func IreEncode(value string) string {
	// Escape an IRI value for use in a URL.
	return strings.ReplaceAll(value, " ", "%20")
}

// URLEncode encodes a string as a quoted-printable string.
func URLEncode(value string) string {
	// Escape a URL string.
	return strings.ReplaceAll(value, " ", "%20")
}

// Slugify converts a string into a slug. A slug is a string where spaces and special characters are replaced by a hyphen, suitable for use in a URL.
func Slugify(value string) string {
	// Convert a string into a slug.
	return strings.ReplaceAll(value, " ", "-")
}

// Title converts a string into titlecase.
func Title(value string) string {
	// Convert a string into titlecase.
	return strings.Title(value)
}

// Truncatechars truncates a string after a certain number of characters.
func Truncatechars(value string, arg int) string {
	// Truncate a string after a certain number of characters.
	if len(value) > arg {
		return value[:arg]
	}
	return value
}

// Truncatewords truncates a string after a certain number of words.
func Truncatewords(value string, arg int) string {
	// Truncate a string after a certain number of words.
	words := strings.Fields(value)
	if len(words) > arg {
		return strings.Join(words[:arg], " ...")
	}
	return value
}

// Wordcount returns the number of words in a string.
func Wordcount(value string) int {
	// Return the number of words in a string.
	return len(strings.Fields(value))
}

// Wordwrap wraps a string after a certain number of characters.
func Wordwrap(value string, arg int) string {
	// Wrap a string after a certain number of characters.
	words := strings.Fields(value)
	var lines []string
	var line string
	for _, word := range words {
		if len(line)+len(word) > arg {
			lines = append(lines, line)
			line = word
		} else {
			line += " " + word
		}
	}
	lines = append(lines, line)
	return strings.Join(lines, "\n")
}

// Cut removes all values of a certain string from the input.
func Cut(value string, arg string) string {
	// Remove all values of a certain string from the input.
	return strings.ReplaceAll(value, arg, "")
}

// Dictsort takes a list of dictionaries and returns that list sorted by the key given in the argument.
func Dictsort(value []map[string]string, arg string) []map[string]string {
	// Sort a list of dictionaries by a key.
	sort.Slice(value, func(i, j int) bool {
		return value[i][arg] < value[j][arg]
	})
	return value
}

// Dictsortreversed takes a list of dictionaries and returns that list sorted by the key given in the argument in reverse order.
func Dictsortreversed(value []map[string]string, arg string) []map[string]string {
	// Sort a list of dictionaries by a key in reverse order.
	sort.Slice(value, func(i, j int) bool {
		return value[i][arg] > value[j][arg]
	})
	return value
}

// FirstItem returns the first item in a list, can by a list of int or string
func FirstItem(value interface{}) interface{} {
	// Return the first item in a list.
	switch v := value.(type) {
	case []int:
		return v[0]
	case []string:
		return v[0]
	}
	return nil
}

// LastItem returns the last item in a list, can by a list of int or string
func LastItem(value interface{}) interface{} {
	// Return the last item in a list.
	switch v := value.(type) {
	case []int:
		return v[len(v)-1]
	case []string:
		return v[len(v)-1]
	}
	return nil
}

// RandomItem returns a random item from the list, can by a list of int or string
func RandomItem(value interface{}) interface{} {
	// Return a random item from a list.
	switch v := value.(type) {
	case []int:
		size := len(v)
		if size == 0 {
			return nil
		}

		// Get a random index
		index := rand.IntN(size)
		return v[index]
	case []string:
		size := len(v)
		if size == 0 {
			return nil
		}

		// Get a random index
		index := rand.IntN(size)
		return v[index]
	}
	return nil
}

// Date formats a date according to the given format.
// The format is like PHP's date function format.
func Date(value string, format string) string {
	// Convert PHP date format to Go's date format
	goFormat, err := pkg.DateFormats(format)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	// Parse the date
	t, err := time.Parse(goFormat, value)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	// Return the formatted date
	return t.Format(goFormat)
}

// TimeSince returns the time since the given date.
// Format a date as the time since that date (i.e. "4 days, 6 hours").
func TimeSince(value string) string {
	// Parse the date
	t, err := time.Parse(time.RFC3339, value)
	if err != nil {
		return ""
	}

	duration, err := durafmt.ParseString(time.Since(t).Round(time.Second).String())
	if err != nil {
		fmt.Println(err)
	}

	// Return the time since the date in integer format
	return duration.String()
}