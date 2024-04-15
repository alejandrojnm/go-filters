package gofilters

import (
	"math/rand"
	"reflect"
	"strings"
	"testing"
	"time"
)

func TestLineNumbers(t *testing.T) {
	tests := []struct {
		name       string
		value      string
		autoescape bool
		want       string
	}{
		{
			name:       "Test with autoescape false",
			value:      "Hello\nWorld",
			autoescape: false,
			want:       "1. Hello\n2. World",
		},
		{
			name:       "Test with autoescape true",
			value:      "<Hello>\n<World>",
			autoescape: true,
			want:       "1. &lt;Hello&gt;\n2. &lt;World&gt;",
		},
		{
			name:       "Test with single line",
			value:      "Hello",
			autoescape: false,
			want:       "1. Hello",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LineNumbers(tt.value, tt.autoescape); got != tt.want {
				t.Errorf("LineNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDictsort(t *testing.T) {
	tests := []struct {
		name  string
		value []map[string]string
		arg   string
		want  []map[string]string
	}{
		{
			name: "Test with unsorted dictionaries",
			value: []map[string]string{
				{"name": "banana", "color": "yellow"},
				{"name": "apple", "color": "red"},
				{"name": "grape", "color": "purple"},
			},
			arg: "name",
			want: []map[string]string{
				{"name": "apple", "color": "red"},
				{"name": "banana", "color": "yellow"},
				{"name": "grape", "color": "purple"},
			},
		},
		{
			name: "Test with already sorted dictionaries",
			value: []map[string]string{
				{"name": "apple", "color": "red"},
				{"name": "banana", "color": "yellow"},
				{"name": "grape", "color": "purple"},
			},
			arg: "name",
			want: []map[string]string{
				{"name": "apple", "color": "red"},
				{"name": "banana", "color": "yellow"},
				{"name": "grape", "color": "purple"},
			},
		},
		{
			name: "Test with dictionaries sorted by color",
			value: []map[string]string{
				{"name": "banana", "color": "yellow"},
				{"name": "apple", "color": "red"},
				{"name": "grape", "color": "purple"},
			},
			arg: "color",
			want: []map[string]string{
				{"name": "grape", "color": "purple"},
				{"name": "apple", "color": "red"},
				{"name": "banana", "color": "yellow"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Dictsort(tt.value, tt.arg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Dictsort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRandomItem(t *testing.T) {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator

	tests := []struct {
		name  string
		value interface{}
	}{
		{
			name:  "Test with slice of integers",
			value: []int{1, 2, 3, 4, 5},
		},
		{
			name:  "Test with slice of strings",
			value: []string{"apple", "banana", "cherry"},
		},
		{
			name:  "Test with unsupported type",
			value: 42,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RandomItem(tt.value)

			switch v := tt.value.(type) {
			case []int:
				if got != nil {
					gotInt, ok := got.(int)
					if !ok {
						t.Errorf("RandomItem() returned a non-int value: %v", got)
					} else if !containsInt(v, gotInt) {
						t.Errorf("RandomItem() returned an integer not in the original slice: %v", gotInt)
					}
				}
			case []string:
				if got != nil {
					gotStr, ok := got.(string)
					if !ok {
						t.Errorf("RandomItem() returned a non-string value: %v", got)
					} else if !containsString(v, gotStr) {
						t.Errorf("RandomItem() returned a string not in the original slice: %v", gotStr)
					}
				}
			default:
				if got != nil {
					t.Errorf("RandomItem() should return nil for unsupported types, got: %v", got)
				}
			}
		})
	}
}

func TestDate(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		format  string
		want    string
		wantErr bool
	}{
		{
			name:    "Test with date format",
			value:   "2021-01-02",
			format:  "Y-m-d",
			want:    "2021-01-02",
		},
		{
			name:    "Test with date and time format",
			value:   "2021-01-02 15:04:05",
			format:  "Y-m-d H:i:s",
			want:    "2021-01-02 15:04:05",
		},
		{
			name:    "Test with invalid date format",
			value:   "2021-01-02",
			format:  "invalid",
			want:    "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Date(tt.value, tt.format)
			if got != tt.want {
				t.Errorf("Date() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTimeSince(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  string
	}{
		{
			name:  "Test with valid RFC3339 date",
			value: time.Now().Add(-5 * time.Minute).Format(time.RFC3339),
			want:  "5 minutes",
		},
		{
			name:  "Test with valid RFC3339 date",
			value: time.Now().Add(-1 * time.Hour).Format(time.RFC3339),
			want:  "1 hour",
		},
		{
			name:  "Test with valid RFC3339 date",
			value: time.Now().Add(-1 * time.Second).Format(time.RFC3339),
			want:  "2 second",
		},
		{
			name:  "Test with invalid date",
			value: "invalid date",
			want:  "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TimeSince(tt.value)
			if !strings.HasPrefix(got, tt.want) {
				t.Errorf("TimeSince() = %v, want prefix %v", got, tt.want)
			}
		})
	}
}

func TestTimeUntil(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  string
	}{
		{
			name:  "Test with valid RFC3339 date",
			value: time.Now().Add(5 * time.Minute).Format(time.RFC3339),
			want:  "5 minutes",
		},
		{
			name:  "Test with valid RFC3339 date",
			value: time.Now().Add(1 * time.Hour).Format(time.RFC3339),
			want:  "1 hour",
		},
		{
			name:  "Test with valid RFC3339 date",
			value: time.Now().Add(1 * time.Second).Format(time.RFC3339),
			want:  "1 second",
		},
		{
			name:  "Test with invalid date",
			value: "invalid date",
			want:  "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TimeUntil(tt.value)
			if !strings.HasPrefix(got, tt.want) {
				t.Errorf("TimeUntil() = %v, want prefix %v", got, tt.want)
			}
		})
	}
}

// Helper function to check if a slice of integers contains a specific integer
func containsInt(slice []int, item int) bool {
	for _, a := range slice {
		if a == item {
			return true
		}
	}
	return false
}

// Helper function to check if a slice of strings contains a specific string
func containsString(slice []string, item string) bool {
	for _, a := range slice {
		if a == item {
			return true
		}
	}
	return false
}
