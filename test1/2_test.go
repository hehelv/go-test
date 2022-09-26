package test1

import (
	"reflect"
	"testing"
)

func TestPrint(t *testing.T) {
	tests := []struct {
		name string
		want []int
	}{
		// TODO: Add test cases.
		{"素数",[]int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Print(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Print() = %v, want %v", got, tt.want)
			}
		})
	}
}