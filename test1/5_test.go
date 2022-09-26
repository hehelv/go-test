package test1

import (
	"reflect"
	"testing"
)

func TestPrintSlice(t *testing.T) {
	tests := []struct {
		name string
		want []int
	}{
		// TODO: Add test cases.
		{"求切片", []int{1, 2, 4, 5, 7, 8, 10, 11, 13, 14, 16, 17, 19, 20, 22, 23, 25, 26, 28, 29, 31, 32, 34, 35, 37, 38, 40, 41, 43, 44, 46, 47, 49, 50, 666}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PrintSlice(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PrintSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
