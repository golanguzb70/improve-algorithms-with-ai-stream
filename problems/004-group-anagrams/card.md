# 004 · Group Anagrams: Hash Map

**In my words:** Problem is asking for group of strings that are anagrams of each others. In parameter they are give as a array of strings. n is length of array and m lenth of strings.  
**Signals:** Anagram, Count of letter in each string.
**Brute force:** Loop through the array and check any element except iteslf if it is anagram of that word or not. Append those anagram to response array and remove them from strs array.  loop: for index in strs: 
    nested loop: for e in strs[index:]: 
        nested: check if anagram.  O(m)
So main time complexiy is O(n^2 * m)
**Optimal idea:** Create a map that saves one sorted word as key and value is array of string. Loop through the strings array and check if sorted string is anagargram of any key and if so add to the group, if not create new group. To make this more optimal to return instant response we save values in [][]strings and in map we just save the index of each group.  Space complexity becomes O(n*m) because we are creating response array and sorting each string. Space that is used for sorting is not mentioned because it is in worst case constant.
**Complexity:** time O(n * m log m), space O(n * m)
**Edge cases:** Lots of different groups, All in one group.
**Mistake I made:** I did mistace on Big O calculation again dropped m. 

**60-second explanation:**
Okay, I solved the problem with the time complexity of n times m log m and space complexity nm, big of m nm. And n is the length of request
  strings array, and m is the length of each word. So the solution that I did is first creating index map, which is saves key. The key is a sorted string, and the
  value is index of that word's anagram group, and it loops through the strings array and sort each element and append it to its corresponding group. As a result,
  return the group. 
So in brute force, we are skipping nested loop to check anagrams and check all the groups first, and removing the already calculated or concluded strings from
  strings array. And also time complexity was n squared times m in brute force, but now it's n times m log m, which is much faster.

## Review
2026-11-22 · Code 8/10 · Explanation 6/10 · Tests 9/9 ✅
- 🔴 Big-O: word length m isn't a constant. Time is O(n · m log m), space is O(n · m) (the map keys are sorted copies). Only things fixed by the problem itself, like 26 letters, become O(1).
- 🔴 Brute force: comparing every pair is n², and one anagram check costs O(m), so O(n² · m), not O(n³). You mixed up n (words) and m (letters).
- 🟡 Go: `slices.Sort(b)` is enough, no SortFunc + cmp needed. Drop the magic `len(strs)/4` capacity. `e` → `word`, `responseStrGroups` → `groups`.
- 🟡 60-sec: ~150 words, says the Big-O twice, and "why it beats brute force" is only the numbers.
- 🟡 "Mistake I made" should be an actual mistake (e.g. dropping m).
- 💡 Optimal? Good, but there's a faster key. Hint: in 002 you didn't sort.

2026-11-22 (re-review) · Code 8/10 · Explanation 7/10 · Tests 9/9 ✅
- Fixed: Big-O in the card, brute force O(n²·m), "Mistake I made", `slices.Sort`.
- 🔴 The solution.go header still says Time O(n) / Space O(n). Make it match the card.
- 🟡 The 60-sec says O(n) first and corrects itself later. State the right Big-O once, at the end.
- 🟡 Still open: `len(strs)/4` capacity, names `e`/`responseStrGroups`, run `gofmt`.

2026-11-22 (re-review 2) · Code 9/10 · Explanation 7/10 · Tests 9/9 ✅ · ✅ Solved
- Fixed: solution.go header, names, capacity, gofmt. The 60-sec is shorter and gives the Big-O once.
- 🟡 The 60-sec lost the "why it beats brute force" part. Order: idea → what the map replaces → Big-O.
- 🟡 Stale comments: `// sort e`, `m = len(e)`.
- 💡 Bonus: a key built without sorting gets O(n · m).

2026-11-22 (re-review 3) · Code 9/10 · Explanation 7/10 · ✅ Solved
- Fixed: stale comments. The 60-sec now has a "why".
- 🟡 The "why" is muddled, and the 60-sec is ~150 words and opens with Big-O. Say it in one line: "The sorted word is a shared key, so one map lookup replaces comparing against every other word."
- 🟡 gofmt: an empty line with spaces is left in the loop.
