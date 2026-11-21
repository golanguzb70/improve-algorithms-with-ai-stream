package lesson

import (
	"reflect"
	"slices"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		in   []int
		want []int
	}{
		{[]int{1, 1, 2, 3, 3}, []int{1, 2, 3}},
		{[]int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{7, 7, 7}, []int{7}},
		{[]int{}, []int{}},
		{[]int{-2, -2, 0, 5, 5, 5}, []int{-2, 0, 5}},
	}
	for _, tt := range tests {
		nums := slices.Clone(tt.in)
		k := RemoveDuplicates(nums)
		if k != len(tt.want) || !reflect.DeepEqual(nums[:k], tt.want) {
			t.Errorf("RemoveDuplicates(%v): k=%d nums[:k]=%v, want %v", tt.in, k, nums[:k], tt.want)
		}
	}
}

func TestSortedSquares(t *testing.T) {
	tests := []struct {
		in   []int
		want []int
	}{
		{[]int{-4, -1, 0, 3, 10}, []int{0, 1, 9, 16, 100}},
		{[]int{-7, -3, 2, 3, 11}, []int{4, 9, 9, 49, 121}},
		{[]int{1, 2, 3}, []int{1, 4, 9}},
		{[]int{-3, -2, -1}, []int{1, 4, 9}},
		{[]int{0}, []int{0}},
	}
	for _, tt := range tests {
		if got := SortedSquares(tt.in); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("SortedSquares(%v) = %v, want %v", tt.in, got, tt.want)
		}
	}
}
