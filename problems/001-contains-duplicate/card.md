# 001 · Contains Duplicate: Hash set

**In my words:** I understand that if one values contains in the nums array more than once I should return true otherwise false.
**Signals:** Seen at least twice, have I seen before?, 
**Brute force:** Brute force is to do nested to loop to search the element from the array again skipp the index of the element. O(n^2)
**Optimal idea:** Optimal idea is to store elements as keys in hash map while looping through the nums array and return true when that elements occurs again. And sorting can also be option because its time complexity is O(n log n) and space is O(1). I chose hash set instead. 
**Complexity:** time O(n), space O(n)
**Edge cases:** a single element, negative numbers, extreme values, and a repeat that sits next to its twin or at the very end.
**Mistake I made:** I didn't make any mistakes.

**60-second explanation:**
Okay, the problem is to find duplicate values from array, and brute force is to do a nested loop search, and its time complexity is n squared, but it
  is not option because the constraints does not allow us to do so. But and I have two optimal options, which is there are trade-offs to choose between
  them. One is hash set, looping through the nums array and remembering each element in specific map. And if the element occurs more than once, or at
  least twice, then I will return true. If not, just go until the end. This is the according to time complexity, it's O(n), and the space complexity is
  also O(n) on worst case. But my next solution is sorting the array first, and while looping through the array, when looking at, let's say, when two
  neighboring elements are equals, and it means it's your the duplicate occurred, and its space complexity is O(1) and time complexity is O(n log n). And
  according to trade-offs, we lose it loses according to time complexity, but wins on space complexity. But I choose to go with a hash map solution because I want to optimize for CPU. 





## Review
2026-11-21 · Code 9/10 · Explanation 6/10 · Tests 9/9 ✅
- 🔴 Edge cases: "No edge cases" is wrong. List them: a single element, negative numbers, extreme values, and a repeat that sits next to its twin or at the very end.
- 🔴 60-second explanation: the map stores **values**, not indexes. Add the Big-O and *why* a set beats the alternatives.
- 🟡 Sorting is the *alternative* (O(n log n) time, O(1) extra space), not the brute force. The brute force is comparing every pair: O(n²).
- 🟡 Pattern name: "Hash set", not "Duplicate finder".

2026-11-21 (re-review) · Code 10/10 · Explanation 8/10 · ✅ Solved
- Both must-fix items are done. The brute force is labelled correctly, and the trade-off is explained well.
- 🟡 Open with the answer and keep it under 60 seconds (about 120 words).
- 🟡 Sorting also mutates the input, and Go's sort uses O(log n) stack space. Mention both.
