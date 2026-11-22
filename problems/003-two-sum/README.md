# 003 · Two Sum

🔗 https://leetcode.com/problems/two-sum/ · 🟢 Easy · Week 1 · Arrays & Hashing

## Problem
You get an array of integers `nums` and an integer `target`.
Find **two different positions** whose values add up to `target`, and return those two **indexes**.
There is always exactly one answer. You can't use the same position twice. The order of the two indexes doesn't matter.

### Examples
```
Input:  nums = [2,7,11,15], target = 9
Output: [0,1]
Why:    nums[0] + nums[1] = 2 + 7 = 9
```
```
Input:  nums = [3,2,4], target = 6
Output: [1,2]
Why:    2 + 4 = 6. Not [0,0]: you can't use the 3 twice.
```
```
Input:  nums = [3,3], target = 6
Output: [0,1]
Why:    Two different positions can hold the same value.
```

### Constraints
- 2 ≤ nums.length ≤ 10⁴
- -10⁹ ≤ nums[i] ≤ 10⁹
- -10⁹ ≤ target ≤ 10⁹
- Exactly one valid answer exists.
- Follow-up: can you beat O(n²) time?

## ✅ Tasks
- [x] 1. **Understand:** restate the problem in `card.md` (in your own words)
- [x] 2. **Match:** write down the *signals*, then name the pattern
- [x] 3. **Brute force:** describe it and give its Big-O
- [x] 4. **Plan:** the optimal idea in 3–5 bullets, plus its Big-O (no code yet)
- [x] 5. **Code:** implement it in `solution.go`
- [ ] 6. **Test:** `go test .` passes, then add 2+ edge cases of your own to `solution_test.go`
- [x] 7. **Explain:** write the 60-second explanation in `card.md` and say it out loud
- [x] 8. **Review:** say "review"

⏱ Stuck for 25 minutes? Say "help".

## Run
```bash
go test ./problems/003-two-sum/        # from the repo root
go test -run TestX -v .               # from inside this folder
```
