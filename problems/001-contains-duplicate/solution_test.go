package solution

import "testing"

func TestContainsDuplicate(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{name: "example 1: repeat at the ends", nums: []int{1, 2, 3, 1}, want: true},
		{name: "example 2: all different", nums: []int{1, 2, 3, 4}, want: false},
		{name: "example 3: many repeats", nums: []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, want: true},

		// 🔍 Reviewer edge cases
		{name: "single element", nums: []int{7}, want: false},
		{name: "two equal", nums: []int{5, 5}, want: true},
		{name: "negatives repeat", nums: []int{-1, 2, -1}, want: true},
		{name: "negative vs positive are different", nums: []int{-3, 3}, want: false},
		{name: "extreme values", nums: []int{-1_000_000_000, 1_000_000_000, -1_000_000_000}, want: true},
		{name: "zeros", nums: []int{0, 1, 0}, want: true},

		// ✍️ Your edge cases: think about size, signs, and where the repeat is
		// {name: "...", nums: []int{...}, want: ...},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsDuplicate(tt.nums); got != tt.want {
				t.Errorf("containsDuplicate(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}
