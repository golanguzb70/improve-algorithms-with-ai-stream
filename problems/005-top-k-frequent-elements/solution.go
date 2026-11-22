package solution

// Top K Frequent Elements
// https://leetcode.com/problems/top-k-frequent-elements/
//
// Time:  O(n)
// Space: O(n)
func topKFrequent(nums []int, k int) []int {
	numsFrequencyMap := make(map[int]int)
	for _, number := range nums {
		numsFrequencyMap[number]++
	}

	buckets := make([][]int, len(nums)+1)
	for number, frequency := range numsFrequencyMap {
		buckets[frequency] = append(buckets[frequency], number)
	}
	response := make([]int, 0, k)
	for i := len(buckets) - 1; i > 0; i-- {
		if k <= 0 {
			break
		}
		for _, number := range buckets[i] {
			response = append(response, number)
			k--
		}
	}
	return response
}
