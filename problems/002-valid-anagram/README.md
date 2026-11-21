# 002 · Valid Anagram

🔗 https://leetcode.com/problems/valid-anagram/ · 🟢 Easy · Week 1 · Arrays & Hashing

## Problem
You get two strings, `s` and `t`.
Return `true` if `t` is an **anagram** of `s`. An anagram uses exactly the same letters, each the same number of times, in any order. Otherwise return `false`.

### Examples
```
Input:  s = "anagram", t = "nagaram"
Output: true
Why:    both have a×3, n×1, g×1, r×1, m×1, just rearranged.
```
```
Input:  s = "rat", t = "car"
Output: false
Why:    s has a 't' and no 'c'; t has a 'c' and no 't'.
```
```
Input:  s = "aab", t = "abb"
Output: false
Why:    same letters, but different counts (a×2 vs a×1).
```

### Constraints
- `1 <= len(s), len(t) <= 5 · 10⁴`
- `s` and `t` contain only lowercase English letters (`a`–`z`)

**Follow-up (interviewers often ask this):** what changes if the strings can contain any Unicode characters?

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
go test ./problems/002-valid-anagram/        # from the repo root
go test -run TestIsAnagram -v .              # from inside this folder
```
