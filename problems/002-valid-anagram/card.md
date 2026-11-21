# 002 · Valid Anagram: Counter

**In my words:** Problem is asking if 2 strings are anagrams which means that do they contain same letter in same count.
**Signals:** Anagram, counter of letters, Did i see  before?
**Brute force:** Loop through s and count each character in both s and t, and if they are not same count return false. otherwise continue. And do same with t.  O(n^2)
**Optimal idea:** I have 2 solutions. 
Solution1: Return false if length is not same. Hash letters of s with values count and loop through t and decrease the counts in the map with letter of t as a key. And if any count goes negative, return false. Otherwise return true. Complexity: time O(n), space O(1). 
Solution2: Return false if length is not same. Sort both strings and check if every index's elements are same. If yes, it will anagram. Complexity. time O(n log n), space O(n).
**Complexity:** time O(n), space O(1) Solution 1 is chosen
**Edge cases:** length not same, only last letter differs, 
**Mistake I made:** This time it's the space and brute-force Big-O. Write that down.

**60-second explanation:**
So I solved the problem by using hash map and iterating through the S and keeping the count of each letter. And next loop, I iterated through the next string,
  which is T, and decrease the counts in the map. if the count values goes negative return false. And the time complexity is big O(n), and space complexity is  big O(1), because I'm creating map that saves only letter from a to z always. And also there is the one
  earlier false return case, which is if the length of S and T is not equals, it returns false.


## psudo solution 1:
s=anagram

t=nagaram
mp= {a:3, n:1, g:1, r:1, m:1}

for letter in t {
    mp[letter]--
    if  mp[letter] < 0 {
        return false
    }
}

t=nagaram
loop1: n
mp= {a:3, n:0, g:1, r:1, m:1}

loop2: a
mp= {a:2, n:0, g:1, r:1, m:1}

loop3: g
mp= {a:2, n:0, g:0, r:1, m:1}

loop4: a
mp= {a:1, n:0, g:0, r:1, m:1}

loop5: r
mp= {a:1, n:0, g:0, r:0, m:1}

loop6: a
mp= {a:0, n:0, g:0, r:0, m:1}

loop7: m
mp= {a:0, n:0, g:0, r:0, m:0}

result: 
for value in mp {
    if value != 0 {
        return false
    }
}

return true
## Review
2026-11-21 · Code 8/10 · Explanation 7/10 · Tests 9/9 ✅
- 🔴 Space is O(1), not O(n): only 26 lowercase letters are possible, so the map never holds more than 26 keys. (O(k) for an alphabet of size k.)
- 🔴 Brute force is O(n²), not O(n³): n characters × O(n) to count each one.
- 🔴 Sorting space in Go isn't O(1): strings are immutable, so you must copy into []byte or []rune first, which is O(n).
- 🔴 Add 2+ edge cases of your own to solution_test.go (skipped twice now).
- 🟡 The last loop over the map is never needed. Figure out why.

2026-11-21 (re-review) · Code 9/10 · Explanation 7/10 · Tests 11/11 ✅ · ✅ Solved
- Fixed: brute force O(n²), sorting O(n) space, O(1) space in the code, your own edge cases added, the extra loop removed.
- 🟡 The Complexity line still says "space O(n)". Make it O(1).
- 🟡 The idea and the 60-second explanation still describe the removed final loop. Update them, and add *why* it's safe to drop.

2026-11-21 (polish) · Explanation 8/10. The Complexity line and the 60-second explanation are fixed. Still open: the Solution 1 idea mentions the removed loop; "why it is safe" is missing; "Mistake I made" is not yet in your own words.
