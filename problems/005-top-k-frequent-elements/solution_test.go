package solution

import (
	"slices"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want []int // values in any order
	}{
		{name: "example 1: top two", nums: []int{1, 1, 1, 2, 2, 3}, k: 2, want: []int{1, 2}},
		{name: "example 2: single value", nums: []int{1}, k: 1, want: []int{1}},
		{name: "example 3: negative wins", nums: []int{4, 4, -1, -1, -1, 7}, k: 1, want: []int{-1}},

		// 🔍 Reviewer edge cases
		{name: "all the same value", nums: []int{9, 9, 9, 9}, k: 1, want: []int{9}},
		{name: "winner is spread out, not adjacent", nums: []int{5, 1, 5, 2, 5, 3}, k: 1, want: []int{5}},
		{name: "k = all distinct", nums: []int{3, 1, 2}, k: 3, want: []int{1, 2, 3}},
		{name: "most frequent appears last", nums: []int{1, 2, 3, 3, 3, 2}, k: 2, want: []int{2, 3}},
		{name: "extreme values", nums: []int{-10000, 10000, 10000}, k: 1, want: []int{10000}},

		// ✍️ Your edge cases (add at least 2!): think about k = number of distinct values, all the same value, unsorted input
		// {name: "...", nums: []int{...}, k: ..., want: []int{...}},
		{name: "example 4: all distinct ", nums: []int{1, 1, 1, 3, 3, 4, 4, 56, 6, 3}, k: 5, want: []int{1, 3, 4, 6, 56}},
		{name: "example 5: negative numbers", nums: []int{-1, -1, -1, 1, 1, 7}, k: 1, want: []int{-1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := slices.Clone(topKFrequent(slices.Clone(tt.nums), tt.k))
			slices.Sort(got)
			want := slices.Clone(tt.want)
			slices.Sort(want)
			if !slices.Equal(got, want) {
				t.Errorf("topKFrequent(%v, %d) = %v, want %v (any order)", tt.nums, tt.k, got, want)
			}
		})
	}
}
