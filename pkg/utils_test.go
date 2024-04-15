package pkg

import (
	"testing"
	// "time"
)

func TestDateFormats(t *testing.T) {
	tests := []struct {
		name   string
		format string
		want   string
	}{
		{
			name:   "Test 1",
			format: "Y-m-d",
			want:   "2006-01-02",
		},
		{
			name:   "Test 2",
			format: "d/m/Y",
			want:   "02/01/2006",
		},
		{
			name:   "Test 3",
			format: "H:i:s",
			want:   "15:04:05",
		},
		{
			name:   "Test 4",
			format: "Y-m-d H:i:s",
			want:   "2006-01-02 15:04:05",
		},
		{
			name:   "Test 5",
			format: "D d M Y",
			want:   "Mon 02 Jan 2006",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DateFormats(tt.format)
			if err != nil {
				t.Errorf("DateFormats() error = %v", err)
				return
			}
			if got != tt.want {
				t.Errorf("DateFormats() = %v, want %v", got, tt.want)
			}
		})
	}
}
