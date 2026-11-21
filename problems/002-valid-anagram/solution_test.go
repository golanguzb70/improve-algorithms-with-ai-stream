package solution

import "testing"

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		name string
		s, t string
		want bool
	}{
		{name: "example 1: rearranged", s: "anagram", t: "nagaram", want: true},
		{name: "example 2: different letters", s: "rat", t: "car", want: false},
		{name: "example 3: same letters, different counts", s: "aab", t: "abb", want: false},

		// 🔍 Reviewer edge cases
		{name: "identical strings", s: "listen", t: "listen", want: true},
		{name: "single char equal", s: "a", t: "a", want: true},
		{name: "single char differ", s: "a", t: "b", want: false},
		{name: "different lengths", s: "a", t: "ab", want: false},
		{name: "same letter repeated vs one off", s: "aaaa", t: "aaab", want: false},
		{name: "last chars differ", s: "ddddga", t: "ddddgb", want: false},
		{name: "all same but order is different", s: "qwerty", t: "ytrewq", want: true},
		{name: "all 26 letters reversed", s: "abcdefghijklmnopqrstuvwxyz", t: "zyxwvutsrqponmlkjihgfedcba", want: true},

		// ✍️ Your edge cases (add at least 2!): think about length, repeats, identical strings
		// {name: "...", s: "...", t: "...", want: ...},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAnagram(tt.s, tt.t); got != tt.want {
				t.Errorf("isAnagram(%q, %q) = %v, want %v", tt.s, tt.t, got, tt.want)
			}
		})
	}
}
