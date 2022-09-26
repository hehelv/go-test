package test1

import (
	"reflect"
	"testing"
)

func TestCountChar(t *testing.T) {
	tests := []struct {
		name string
		str string
		want map[byte]int
	}{
		// TODO: Add test cases.
		{"统计字母出现个数","abcabcccd",map[byte]int{'a': 2, 'b': 2, 'c': 4, 'd': 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountChar(tt.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CountChar() = %v, want %v", got, tt.want)
			}
		})
	}
}
