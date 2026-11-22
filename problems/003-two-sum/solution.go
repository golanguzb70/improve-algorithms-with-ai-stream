package solution

// Two Sum
// https://leetcode.com/problems/two-sum/
//
// Time:  O(n)
// Space: O(n)
func twoSum(nums []int, target int) []int {
	indexOf := make(map[int]int, len(nums))
	for i, e := range nums {
		index, ok := indexOf[target-e]
		if ok {
			return []int{i, index}
		}
		indexOf[e] = i
	}
	return nil
}
