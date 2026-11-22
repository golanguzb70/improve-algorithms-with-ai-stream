# 003 · Two Sum: Hash Map

**In my words:** Problem is asking to return indexes of two values that adds up to target. There will always be one correct answer.
**Signals:** Two numbers that adds up
**Brute force:** Just iterate through array and iterate through the array in nested loop and check every element if that adds up to the target  O(n^2)
**Optimal idea:** First map the nums array with key element value index and iterate through nums array and check if target-element is found in the map with different index. If so return.
**Complexity:** time O(n), space O(n)
**Edge cases:** Negative values.
**Mistake I made:** Didn't specify why optimal way beats brute force.

**60-second explanation:**
 I have solved the problem using index map and with complexity or of time complexity big O(n) and space complexity is big O(n), because I'm creating map that
  saves as a key the element itself and as a value the index of element, and it reads through the array again, nums array again, and if target subtracted element
  value is found as a map with different index, if so return those indexes, if not continue. 
The reason why hash map beats brute force is that Big O of brute force is O(n^2) while hash map is O(n) as it is not looping through the array again to find the second number, it is remembering that number.



⏰ Redo on 2026-11-25: solve it again from scratch, and finish the open items above.
