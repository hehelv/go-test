package test1

import "testing"

// table driver tests
func TestIsLeapYear(t *testing.T) {

	tests := []struct {
		name string
		year int
		want bool
	}{
		// TODO: Add test cases.
		{"普通闰年 2004", 2004, true},
		{"不是闰年 2002", 2002, false},
		{"特殊闰年2000", 2000, true},
		{"不是闰年1900", 1900, false},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsLeapYear(tt.year); got != tt.want {
				t.Errorf("IsLeapYear() = %v, want %v", got, tt.want)
			}
		})
	}
}