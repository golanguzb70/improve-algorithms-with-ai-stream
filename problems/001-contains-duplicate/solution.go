package solution

// Contains Duplicate
// https://leetcode.com/problems/contains-duplicate/
//
// Time:  O(n)
// Space: O(n)
func containsDuplicate(nums []int) bool {
	mp := make(map[int]struct{}, len(nums))
	for _, n := range nums {
		if _, ok := mp[n]; ok {
			return true
		}
		mp[n] = struct{}{}
	}
	return false
}
