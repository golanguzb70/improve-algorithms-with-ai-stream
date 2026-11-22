# 005 · Top K Frequent Elements

🔗 https://leetcode.com/problems/top-k-frequent-elements/ · 🟠 Medium · Week 1 · Arrays & Hashing

## Problem
You get an array of integers `nums` and a number `k`.
Return the `k` values that appear **most often** in `nums`.
The answer is guaranteed to be unique, and you can return the values in any order.

### Examples
```
Input:  nums = [1,1,1,2,2,3], k = 2
Output: [1,2]
Why:    1 appears 3 times, 2 appears 2 times, 3 only once. The top 2 are 1 and 2.
```
```
Input:  nums = [1], k = 1
Output: [1]
Why:    Only one value, so it's the most frequent.
```
```
Input:  nums = [4,4,-1,-1,-1,7], k = 1
Output: [-1]
Why:    -1 appears 3 times, more than anything else.
```

### Constraints
- 1 ≤ nums.length ≤ 10⁵
- -10⁴ ≤ nums[i] ≤ 10⁴
- 1 ≤ k ≤ number of distinct values in nums
- The answer is unique (no ties at the cut-off).
- Follow-up: can you beat O(n log n) time?

## ✅ Tasks
- [x] 1. **Understand:** restate the problem in `card.md` (in your own words)
- [x] 2. **Match:** write down the *signals*, then name the pattern
- [x] 3. **Brute force:** describe it and give its Big-O
- [x] 4. **Plan:** the optimal idea in 3–5 bullets, plus its Big-O (no code yet)
- [x] 5. **Code:** implement it in `solution.go`
- [x] 6. **Test:** `go test .` passes, then add 2+ edge cases of your own to `solution_test.go`
- [x] 7. **Explain:** write the 60-second explanation in `card.md` and say it out loud
- [x] 8. **Review:** say "review"

⏱ Stuck for 25 minutes? Say "help".

## Run
```bash
go test ./problems/005-top-k-frequent-elements/        # from the repo root
go test -run TestX -v .               # from inside this folder
```
