# 005 · Top K Frequent Elements: bucket sort, grouping by frequency.

**In my words:** I will be given array of integers and I need to return top k numbers according to their frequency.
**Signals:** Frequency, Count of each numbers, counts are bounded by n
**Brute force:** Sort the array by comparing count of each values and it should count each number in the array for each comparison. Return first k elements.
nums.sort(func(a,b) {
    // count a in nums 
    // cont b in nums 
    // and return comparison of those nums. 
})
 O(n^2 log n)
**Optimal idea:** First, I build a map from each number to its count. That's O(n). Then make an array of n+1 slices. Put each number into the slot that matches its count. Walk the array from the highest slot down, gathering numbers until you have k.
**Complexity:** time O(n), space O(n)
**Edge cases:** k is too big, nums[i] negative
**Mistake I made:** I dropped the O(n) counting loop.

**60-second explanation:**
Idea: Count frequencies and remember. Group numbers by frequency because frequency is always between 1 and n. Gather results from end to start unti k. 
Why: By this way, we avoid sorting whole array and as I used indexes as frequency buckets are always sorted.
Big-O: Time O(n), Space O(n)

## Review
2026-11-22 · Code 8/10 · Explanation 6/10 · Tests 10/10 ✅
- 🔴 Big-O drops the counting loop. You count all n numbers first, then sort m keys: time O(n + m log m), worst case (all distinct) O(n log n). Space O(m) ✅.
- 🟡 "Mistake I made" is a placeholder again. Fill it in.
- 🟡 60-sec: opens with Big-O, ~130 words. Order: idea → why → Big-O.
- 🟡 Go: `cmp.Compare(freq[b], freq[a])` sorts descending, no `-1 *` needed. Typo `uniquNumbers`. Run gofmt on the test file.
- 💡 Follow-up (beat O(n log n)): a frequency can never be bigger than n. What if frequencies were used as indexes?

2026-11-22 (re-review, bucket version) · Code 9/10 · Explanation 8/10 · Tests 10/10 ✅ · 🟡 (level-3 hint) · Redo 2026-11-23
- Fixed: Big-O O(n)/O(n) correct everywhere, a real "Mistake I made", and a 60-sec in idea → why → Big-O order.
- 🟡 "Optimal idea" is the tutor's hint copied word for word. Rewrite it in your words.
- 🟡 Code: `if len(...) > 0` isn't needed (ranging over an empty slice does nothing), and k is checked in two places. One loop condition is enough. `topFrequents` → `buckets`.
- 🟡 Still open: gofmt the test file; rename "k is too big" → "k = all distinct".

2026-11-22 (re-review 2) · Code 9/10 · Explanation 8/10 · Tests 10/10 ✅ · gofmt ✅ · 🟡 · Redo 2026-11-23
- Fixed: redundant checks, `buckets` name, gofmt, the test name, "Space O(n)".
- 🟡 "Optimal idea" is half rewritten. The last two sentences are still the hint's words.
- 💡 Interview follow-up: the inner loop no longer stops at k. That's safe only because the answer is guaranteed unique. Be ready to say why.
