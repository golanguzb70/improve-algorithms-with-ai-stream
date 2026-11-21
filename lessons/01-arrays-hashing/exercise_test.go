package lesson

import (
	"reflect"
	"testing"
)

func TestWordFrequency(t *testing.T) {
	tests := []struct {
		in   []string
		want map[string]int
	}{
		{[]string{"go", "is", "go"}, map[string]int{"go": 2, "is": 1}},
		{[]string{"a"}, map[string]int{"a": 1}},
		{[]string{}, map[string]int{}},
	}
	for _, tt := range tests {
		got := WordFrequency(tt.in)
		if len(got) != len(tt.want) || (len(tt.want) > 0 && !reflect.DeepEqual(got, tt.want)) {
			t.Errorf("WordFrequency(%v) = %v, want %v", tt.in, got, tt.want)
		}
	}
}

func TestFirstRepeated(t *testing.T) {
	tests := []struct {
		in   []int
		want int
	}{
		{[]int{3, 1, 4, 1, 3}, 1},
		{[]int{1, 2, 3}, -1},
		{[]int{}, -1},
		{[]int{0, 0}, 0},
		{[]int{-5, 2, -5}, -5},
	}
	for _, tt := range tests {
		if got := FirstRepeated(tt.in); got != tt.want {
			t.Errorf("FirstRepeated(%v) = %d, want %d", tt.in, got, tt.want)
		}
	}
}

func TestRangeSums(t *testing.T) {
	tests := []struct {
		nums    []int
		queries [][2]int
		want    []int
	}{
		{[]int{2, 4, 1, 3}, [][2]int{{0, 1}, {1, 3}}, []int{6, 8}},
		{[]int{5}, [][2]int{{0, 0}}, []int{5}},
		{[]int{-1, 2, -3, 4}, [][2]int{{0, 3}, {2, 2}}, []int{2, -3}},
	}
	for _, tt := range tests {
		if got := RangeSums(tt.nums, tt.queries); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("RangeSums(%v, %v) = %v, want %v", tt.nums, tt.queries, got, tt.want)
		}
	}
}
