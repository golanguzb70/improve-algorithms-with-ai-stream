# 004 · Group Anagrams

🔗 https://leetcode.com/problems/group-anagrams/ · 🟠 Medium · Week 1 · Arrays & Hashing

## Problem
You get a list of words `strs`.
Put words that are **anagrams of each other** (same letters, same counts, any order) into the same group.
Return the list of groups. The order of the groups, and the order of words inside a group, doesn't matter.

### Examples
```
Input:  strs = ["eat","tea","tan","ate","nat","bat"]
Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
Why:    eat/tea/ate use the letters a,e,t. tan/nat use a,n,t. bat has no partner.
```
```
Input:  strs = [""]
Output: [[""]]
Why:    One empty word is a group of its own.
```
```
Input:  strs = ["a"]
Output: [["a"]]
Why:    A single word is a group of one.
```

### Constraints
- 1 ≤ strs.length ≤ 10⁴
- 0 ≤ strs[i].length ≤ 100
- strs[i] has only lowercase English letters.

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
go test ./problems/004-group-anagrams/        # from the repo root
go test -run TestX -v .               # from inside this folder
```
