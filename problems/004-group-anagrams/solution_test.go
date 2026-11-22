package solution

import (
	"slices"
	"strings"
	"testing"
)

// normalize sorts words inside each group, then sorts the groups,
// so any valid order of the answer compares equal.
func normalize(groups [][]string) [][]string {
	out := make([][]string, 0, len(groups))
	for _, g := range groups {
		c := slices.Clone(g)
		slices.Sort(c)
		out = append(out, c)
	}
	slices.SortFunc(out, func(a, b []string) int {
		return strings.Compare(strings.Join(a, ","), strings.Join(b, ","))
	})
	return out
}

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		name string
		strs []string
		want [][]string // any order of groups and of words inside a group
	}{
		{name: "example 1: three groups", strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			want: [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}}},
		{name: "example 2: one empty word", strs: []string{""}, want: [][]string{{""}}},
		{name: "example 3: one word", strs: []string{"a"}, want: [][]string{{"a"}}},

		// 🔍 Reviewer edge cases
		{name: "same letters, different counts", strs: []string{"aab", "abb", "bab"}, want: [][]string{{"aab"}, {"abb", "bab"}}},
		{name: "duplicate words", strs: []string{"ab", "ab", "ba"}, want: [][]string{{"ab", "ab", "ba"}}},
		{name: "several empty words", strs: []string{"", "a", ""}, want: [][]string{{"", ""}, {"a"}}},
		{name: "different lengths, same letters", strs: []string{"a", "aa", "aaa"}, want: [][]string{{"a"}, {"aa"}, {"aaa"}}},

		// ✍️ Your edge cases (add at least 2!): think about duplicate words, same letters with different counts, all in one group
		// {name: "...", strs: []string{...}, want: [][]string{{...}}},
		{name: "example 4: Lots of different groups ", strs: []string{"a", "b", "c", "d"}, want: [][]string{{"a"}, {"b"}, {"c"}, {"d"}}},
		{name: "example 5: All in one group ", strs: []string{"adb", "dba", "bda", "dab"}, want: [][]string{{"adb", "dba", "bda", "dab"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := normalize(groupAnagrams(slices.Clone(tt.strs)))
			want := normalize(tt.want)
			if !slices.EqualFunc(got, want, slices.Equal[[]string]) {
				t.Errorf("groupAnagrams(%q) = %q, want %q (any order)", tt.strs, got, want)
			}
		})
	}
}
