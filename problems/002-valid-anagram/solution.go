package solution

// Valid Anagram
// https://leetcode.com/problems/valid-anagram/
//
// Time:  O(n)
// Space: O(1)
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	mp := make(map[rune]int, 26)
	for _, ru := range s {
		mp[ru]++
	}

	for _, ru := range t {
		mp[ru]--
		// for early false detection
		if mp[ru] < 0 {
			return false
		}
	}

	return true
}
