# Lesson 02 · Two Pointers

⏱ ~40 min · Week 1 · Problems 010–014

## 🎯 After this lesson you can
- Replace a nested loop with two indexes that move toward each other or in the same direction
- Know which pointer to move, and **why** that move is safe

---

## 1. The idea in one line
Use **two indexes** into the same array. Each step, move one of them based on a rule. Together they visit each element once → **O(n)** instead of O(n²).

## 2. Two shapes

### Shape A: Opposite ends (converging)
```
[ 1  3  4  6  8  11 ]
  L→              ←R
```
```go
l, r := 0, len(a)-1
for l < r {
    if done(a[l], a[r]) { ... }
    if needBigger { l++ } else { r-- }
}
```
Use it for: **sorted** arrays (pair sums), palindromes, "pick two walls".

### Shape B: Same direction (slow & fast / reader & writer)
```
[ 0  1  0  3  12 ]
  w               ← writer: where the next "kept" element goes
  r→              ← reader: scans everything
```
```go
w := 0
for r := 0; r < len(a); r++ {
    if keep(a[r]) {
        a[w] = a[r]
        w++
    }
}
// a[:w] is the result, done in place
```
Use it for: in-place filtering, removing duplicates, partitioning.

## 3. The key question: why is moving that pointer safe? 🧠
This is what interviewers want to hear. In a **sorted** array, looking for `a[l] + a[r] == target`:
- the sum is too small → `l++`. Every pair using the old `a[l]` with anything ≤ `a[r]` is too small too, so they can all be discarded.
- the sum is too big → `r--`, by the same logic in reverse.

> Each move **throws away a whole row of the O(n²) table** of possible pairs. That's why it's O(n).

Always be ready to say *"I move this pointer because nothing it could pair with can be the answer."*

## 4. Signals: when to think "two pointers"

| The problem says… | Think |
|---|---|
| "sorted array" + pair/triplet with a sum | Opposite ends |
| "palindrome" | Opposite ends |
| "in place", "O(1) extra space", "remove", "move to the end" | Reader/writer |
| "pick two lines/walls to maximize area" | Opposite ends, move the **shorter** one |
| Triplets (a+b+c) | Sort, fix one element, then two pointers on the rest → O(n²) |

## 5. Complexity
- Each pointer moves at most n times → **O(n) time, O(1) space**
- If you sort first: **O(n log n)**

## 6. Common mistakes ⚠️
- `l <= r` vs `l < r`: for pairs, use `l < r` (an element can't pair with itself)
- Forgetting to skip duplicates when the problem wants *unique* triplets
- Using two pointers on an **unsorted** array for a pair sum (use hashing instead, as in lesson 01)
- Infinite loops: every iteration must move at least one pointer

## 7. Hashing vs two pointers

| | Hashing | Two pointers |
|---|---|---|
| Input sorted? | Not needed | Usually required |
| Time | O(n) | O(n), or O(n log n) with a sort |
| Space | O(n) | O(1) |

---

## 🧪 Worked example: Move Zeroes (LeetCode 283, not in our list)
> Move all zeros to the end **in place**, keeping the order of the non-zeros.
> `[0,1,0,3,12]` → `[1,3,12,0,0]`

1. **Signals:** "in place" + "keep the order" → **reader/writer**.
2. **Brute force:** for every zero, shift everything left → O(n²).
3. **Optimal:** the writer `w` marks where the next non-zero goes. The reader scans; on a non-zero, swap it into `w` and move `w` forward.
```go
func moveZeroes(nums []int) {
    w := 0
    for r := range nums {
        if nums[r] != 0 {
            nums[w], nums[r] = nums[r], nums[w]
            w++
        }
    }
}
```
4. **Trace** `[0,1,0,3,12]`:
```
r=0 (0)  skip                 [0,1,0,3,12] w=0
r=1 (1)  swap 0↔1             [1,0,0,3,12] w=1
r=2 (0)  skip                 [1,0,0,3,12] w=1
r=3 (3)  swap 1↔3             [1,3,0,0,12] w=2
r=4 (12) swap 2↔4             [1,3,12,0,0] w=3
```
5. **Complexity:** O(n) time, O(1) space.
6. **60-second explanation:** *"I use a writer pointer for where the next non-zero belongs. The reader scans left to right, and each non-zero it finds gets swapped into the writer's slot, then the writer advances. The non-zeros keep their order and the zeros drift to the end. That's one pass, so O(n) time and O(1) space."*

---

## 🏋️ Warm-up exercise
Open `exercise.go`, implement the 2 functions, then run:
```bash
go test ./lessons/02-two-pointers/
```

## ✍️ Quiz

**1.** Sorted array, find a pair with sum = target. The sum is too big. Which pointer moves, and why?
<details><summary>Answer</summary><code>r--</code>. Pairing <code>a[r]</code> with anything ≥ <code>a[l]</code> is also too big, so <code>a[r]</code> can't be in the answer.</details>

**2.** The array is **unsorted** and you need a pair sum. Two pointers or hashing?
<details><summary>Answer</summary>Hashing: O(n) time. Two pointers would need a sort first, giving O(n log n), though with O(1) space. Mention both trade-offs.</details>

**3.** How do you get the triplets a+b+c = 0 in O(n²)?
<details><summary>Answer</summary>Sort. For each i, run opposite-end two pointers on <code>i+1..n-1</code> looking for <code>-a[i]</code>. Skip duplicate values.</details>

**4.** Which shape fits "remove all occurrences of val in place"?
<details><summary>Answer</summary>Reader/writer (same direction).</details>

---

## ✅ Done when
- [ ] All tests in `exercise_test.go` pass
- [ ] Quiz: 3 of 4 correct
- [ ] You can explain **why** moving a pointer is safe, without notes

➡️ Next: Lesson 03 · Sliding Window (week 2) · Practice: problems **010–014**
