package solution

import (
	"slices"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int // indexes in any order
	}{
		{name: "example 1: first two", nums: []int{2, 7, 11, 15}, target: 9, want: []int{0, 1}},
		{name: "example 2: can't reuse one index", nums: []int{3, 2, 4}, target: 6, want: []int{1, 2}},
		{name: "example 3: same value twice", nums: []int{3, 3}, target: 6, want: []int{0, 1}},

		// 🔍 Reviewer edge cases
		{name: "negatives", nums: []int{-3, 4, 3, 90}, target: 0, want: []int{0, 2}},
		{name: "zeros", nums: []int{0, 4, 3, 0}, target: 0, want: []int{0, 3}},
		{name: "pair at far ends", nums: []int{1, 5, 6, 7, 9}, target: 10, want: []int{0, 4}},
		{name: "half of target appears once", nums: []int{5, 1, 9}, target: 10, want: []int{1, 2}},
		{name: "negative values", nums: []int{5, -11, 21}, target: 10, want: []int{1, 2}},

		{name: "extreme values", nums: []int{1_000_000_000, -1_000_000_000, 7}, target: 0, want: []int{0, 1}},

		// ✍️ Your edge cases (add at least 2!): think about negatives, zero, the pair at the far ends
		// {name: "...", nums: []int{...}, target: ..., want: []int{..., ...}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := slices.Clone(twoSum(slices.Clone(tt.nums), tt.target))
			slices.Sort(got)
			if !slices.Equal(got, tt.want) {
				t.Errorf("twoSum(%v, %d) = %v, want %v (any order)", tt.nums, tt.target, got, tt.want)
			}
		})
	}
}
