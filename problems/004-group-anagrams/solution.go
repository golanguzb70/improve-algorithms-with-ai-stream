package solution

import (
	"slices"
)

// Group Anagrams
// https://leetcode.com/problems/group-anagrams/
//
// Time:  O(n*mlogm)
// Space: O(nm)
func groupAnagrams(strs []string) [][]string {
	groupedAnagramWords := make([][]string, 0)
	indexMap := make(map[string]int)
	for _, word := range strs { // -> time complexity  O(n)
		
		wordInBytes := []byte(word) // -> Space complexity O(m) m = len(word)
		slices.Sort(wordInBytes)

		sortedStr := string(wordInBytes) // Space complexity O(m)
		index, ok := indexMap[sortedStr]
		if ok {
			groupedAnagramWords[index] = append(groupedAnagramWords[index], word)
		} else {
			indexMap[sortedStr] = len(groupedAnagramWords)
			groupedAnagramWords = append(groupedAnagramWords, []string{word})
		}
	}

	return groupedAnagramWords
}
