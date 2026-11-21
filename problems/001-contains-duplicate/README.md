# 001 · Contains Duplicate

🔗 https://leetcode.com/problems/contains-duplicate/ · 🟢 Easy · Week 1 · Arrays & Hashing

## Problem
You get a slice of integers `nums`.
Return `true` if **any value appears at least twice**. Return `false` if every value is different.

### Examples
```
Input:  nums = [1, 2, 3, 1]
Output: true
Why:    1 appears at index 0 and at index 3.
```
```
Input:  nums = [1, 2, 3, 4]
Output: false
Why:    all four values are different.
```
```
Input:  nums = [1, 1, 1, 3, 3, 4, 3, 2, 4, 2]
Output: true
Why:    several values repeat (1, 3, 4, 2). One repeat is enough.
```

### Constraints
- `1 <= len(nums) <= 10⁵`
- `-10⁹ <= nums[i] <= 10⁹`

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
go test ./problems/001-contains-duplicate/        # from the repo root
go test -run TestContainsDuplicate -v .           # from inside this folder
```
